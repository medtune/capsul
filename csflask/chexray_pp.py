import tensorflow as tf

import DenseNet.preprocessing.densenet_pre as dp

import pandas as pd
import json, codecs
import os

ROLE="chexray preprocessing for densenet 121"

flags = tf.app.flags

flags.DEFINE_string('data_path', 'data', 'data directory path')

# Server flags
flags.DEFINE_integer('port', 12030, 'flask listen port')
flags.DEFINE_boolean('debug', False, 'debug flask server')
FLAGS = flags.FLAGS

# Define consts
base_path = FLAGS.data_path

if not os.path.exists(base_path):
    print('creating base path:', base_path)
    os.mkdir(base_path)

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
    image = dp.preprocess_image(image, 224,224, is_training=False)
    image = tf.expand_dims(image,0)

print("[INFO] initialized")

# Run cam on image
def pp_image(graph, src):
    sess = tf.Session(graph=graph)
    img = sess.run([image], {file_input : src})
    sess.close()
    return img

if not os.path.exists(base_path):
    print('creating base path:', base_path)
    os.mkdir(base_path)

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

# Flask app
app = Flask(__name__)

# send json response
def sendJSON(obj):
    return make_response(jsonify(obj))

@app.route('/process', methods=["POST"])
def process_image():
    try:
        target = request.json.get('target')
        targetTemp = target
        target = os.path.join(base_path, target)

        array = pp_image(MainGraph, target)

        return sendJSON({
            'success' : True,
            'target' : targetTemp,
            'out' : pd.Series(array).to_json(orient='values'),
        }), 200

    except Exception as e:
        return sendJSON({
            'success' : False,
            'target' : targetTemp,
            'errors' : [repr(e)],
        }), 200

@app.route('/status', methods=["GET"])
def status():
    return sendJSON({
        'success': True,
        'status' : 'AVAILABLE',
        'version' : 1,
        'pwd' : os.getcwd(),
        'base_path' : base_path,
        'flask_debug' : FLAGS.debug,
        'flask_port' : FLAGS.port,
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