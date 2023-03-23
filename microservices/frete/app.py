import random
from flask import Flask, jsonify
from flask_cors import CORS

app = Flask(__name__)
cors = CORS(app)

@app.route('/frete')
def gerar_numero_aleatorio():
    numero_aleatorio = random.uniform(10.0, 100.0)
    return jsonify({'frete': numero_aleatorio})

if __name__ == '__main__':
    app.run(port=3003)
