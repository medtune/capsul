from flask import (
    make_response,
    jsonify,
)

def sendJSON(obj):
    return make_response(jsonify(obj))