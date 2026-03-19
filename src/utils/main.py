import os
import json
from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from flask_marshmallow import Marshmallow

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///database.db'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
db = SQLAlchemy(app)
ma = Marshmallow(app)

class User(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(80), nullable=False)
    email = db.Column(db.String(120), nullable=False)

class UserSchema(ma.ModelSchema):
    class Meta:
        model = User

@app.route('/users', methods=['POST'])
def create_user():
    data = request.json
    if 'username' not in data or 'email' not in data:
        return jsonify({'error': 'Missing required fields'}), 400
    user = User(username=data['username'], email=data['email'])
    db.session.add(user)
    db.session.commit()
    return jsonify(UserSchema(exclude=['id']).dump(user))

if __name__ == '__main__':
    with app.app_context():
        db.create_all()
    app.run(debug=True)