from flask import Flask, request, jsonify
from table import get_item,put_item, delete_item

app = Flask(__name__)

@app.route("/<key>", methods=['GET'])
def get(key):
    item = get_item(key)
    if item:
        return jsonify(item)
    else:
        return jsonify({"error": "Item not found"}), 404

@app.route("/<key>", methods=['PUT'])
def put(key):
    body = request.json
    body['key'] = key
    res = put_item(body)
    if res:
        return jsonify(res), 200
    else:
        return jsonify({"error": "Failed to put item"}), 500

@app.route("/<key>", methods=['DELETE'])
def delete(key):
    res = delete_item(key)
    if res:
        return jsonify({"message": "Item deleted"}), 200
    else:
        return jsonify({"error": "Failed to delete item"}), 500





