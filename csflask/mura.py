import os

from flask import (
    Flask, 
    redirect, 
    url_for, 
    render_template, 
    request, 
    session,
)

from common_utils import (
    sendJSON,
)


def find_weights(weights_dir):
    return

def load_model():
    return

def run_inference_on_bytes():
    return

def run_cam_on_bytes():
    return

def get_status():
    return


# Flask app
app = Flask(__name__)

@app.route('/predict', methods=["GET", "POST"])
def mura_predict():
    return

@app.route('/cam', methods=["GET", "POST"])
def mura_cam():
    return

@app.route('/status', methods=["GET"])
def mura_status():
    return


if __name__ == "__main__":
    # Run server
    app.run(port=8002, debug=True)

    """
    python3 script.py --port=9020 \
                      --weights_dir=./w \
                      --debug=False
    """