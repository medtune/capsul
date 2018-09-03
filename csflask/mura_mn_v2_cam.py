import tensorflow as tf
slim = tf.contrib.slim

import research.slim.nets.mobilenet.mobilenet_v2 as mobilenet_v2
import DenseNet.preprocessing.densenet_pre as dp

import cv2
import os
import numpy as np

from flask import (
    Flask, 
    redirect, 
    url_for, 
    render_template, 
    request, 
    session,
)

from flask import (
    make_response,
    jsonify,
)

DATASET = "MURA v1.1"
MODEL = "mobilenet v2"

flags = tf.app.flags

# Data flags
flags.DEFINE_string('weights', './weights','weights path')
flags.DEFINE_string('checkpoint', 'model-115001', 'checkpoint file')
flags.DEFINE_string('data_path', 'data', 'data directory path')

# Server flags
flags.DEFINE_integer('port', 11020, 'flask listen port')
flags.DEFINE_boolean('debug', False, 'debug flask server')
FLAGS = flags.FLAGS

# Define consts
base_path = FLAGS.data_path

if not os.path.exists(base_path):
    print('creating base path:', base_path)
    os.mkdir(base_path)

checkpoint_dir = FLAGS.weights
checkpoint_file = os.path.join(checkpoint_dir, FLAGS.checkpoint)
image_size = 224

print("[INFO] initializing tf graph...")

# Create main graph
MainGraph = tf.Graph()
with MainGraph.as_default():
    tf.logging.set_verbosity(tf.logging.INFO)
    file_input = tf.placeholder(tf.string, ())
    image = tf.image.decode_image(tf.read_file(file_input), channels=3)
    copy = tf.identity(image)
    image = tf.image.convert_image_dtype(image, tf.float32)
    image.set_shape([None,None,3])
    image_a = dp.preprocess_image(image, 224,224, is_training=False)
    images_bis = tf.expand_dims(image_a,0)
    with slim.arg_scope(mobilenet_v2.training_scope(is_training=False)):
        #TODO: Check mobilenet_v1 module, var "excluding
        logits, end_points = mobilenet_v2.mobilenet(images_bis,depth_multiplier=1.4, num_classes = 2)
    variables = slim.get_variables_to_restore()
    init_fn=slim.assign_from_checkpoint_fn(checkpoint_file, variables)
    embedding = end_points["layer_18/output"][0]
    weights = tf.reduce_mean(tf.get_default_graph().get_tensor_by_name("MobilenetV2/expanded_conv_16/project/weights:0"), axis=[0,1,2])

print("[INFO] initialized")

# Run cam on image
def run_cam_on_image(graph, src, out):
    sess = tf.Session(graph=graph)
    init_fn(sess)
    emb, ng, raw_img = sess.run([embedding, weights,image], feed_dict={file_input:src})
    cam = np.zeros(emb.shape[0 : 2], dtype = np.float32)
    for j, w in enumerate(ng):
        cam +=  emb[:, :, j]*w
    cam /= np.max(cam)
    cam = cv2.resize(cam, (raw_img.shape[1], raw_img.shape[0]))
    cam = cv2.applyColorMap(np.uint8(255*cam), cv2.COLORMAP_JET)
    cam = cv2.cvtColor(cam, cv2.COLOR_BGR2RGB)
    alpha = 0.0015
    new_img = raw_img+alpha*cam
    cv2.imwrite(out, np.uint8(255*new_img))
    sess.close()

# Flask app
app = Flask(__name__)

# send json response
def sendJSON(obj):
    return make_response(jsonify(obj))

@app.route('/cam', methods=["POST"])
def run_mura_cam():
    try:
        target = request.json.get('target')
        dest = request.json.get('destination')
        force = request.json.get('force')
        destTemp = dest
        targetTemp = target

        # if destination is not given
        # output target + model signature . ext
        # image_0.png -> image_0_mn_v2_cam.png
        if (dest is None) or (dest == ''):
            elm = target.split('.')
            dest = elm[0] + '_mn_v2_cam.' + elm[1]
            destTemp = dest

        target = os.path.join(base_path, target)
        dest = os.path.join(base_path, dest)

        if os.path.exists(dest):
            if force:
                os.remove(dest)
            else:
                return sendJSON({
                    'success' : False,
                    'errors' : [
                        'destination already exist: ' + str(destTemp),
                    ],
                }), 200

        run_cam_on_image(MainGraph, target, dest)

        return sendJSON({
            'success' : True,
            'target' : targetTemp,
            'destination' : destTemp,
        }), 200

    except Exception as e:
        return sendJSON({
            'success' : False,
            'errors' : [repr(e)],
            'target' : targetTemp,
            'destination' : destTemp,
        }), 200

@app.route('/status', methods=["GET"])
def mura_status():
    return sendJSON({
        'success': True,
        'status' : 'AVAILABLE',
        'version' : 1,
        'pwd' : os.getcwd(),
        'base_path' : base_path,
        'checkpoint' : checkpoint_file,
        'flask_debug' : FLAGS.debug,
        'flask_port' : FLAGS.port,
        'dataset': DATASET,
        'model' : MODEL,
    }), 200

@app.route('/list', methods=["GET"])
def get_available_data():
    try:
        files = []
        for fname in os.listdir(base_path):
            path = os.path.join(base_path, fname)
            if os.path.isdir(path):
                continue
            files += [path]

        return sendJSON({
            'success' : True,
            'files' : files,
        })
    except Exception as e:
        return sendJSON({
            'success' : False,
            'errors' : [repr(e)],
        }), 200

if __name__ == "__main__":
    # Run server
    app.run(
        host='0.0.0.0',
        port=FLAGS.port,
        debug=FLAGS.debug,
    )

