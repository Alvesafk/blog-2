from flask import Flask, jsonify, request
from flask.helpers import datetime
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///app.db"
db = SQLAlchemy(app)


class Post(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(100), nullable=False)
    content = db.Column(db.Text, nullable=False)
    created_at = db.Column(db.DateTime, default=datetime.now)
    updated_at = db.Column(db.DateTime, default=datetime.now)

    def to_dict(self):
        return {
            "id": self.id,
            "title": self.title,
            "content": self.content,
            "created_at": self.created_at,
            "updated_at": self.updated_at,
        }


@app.route("/posts", methods=["POST"])
def create_post():
    data = request.get_json()

    if not data or "title" not in data or "content" not in data:
        return jsonify({"error": "no title or content"}), 400

    new_post = Post(title=data["title"], content=data["content"])
    db.session.add(new_post)
    db.session.commit()

    return jsonify(new_post.to_dict()), 201


@app.route("/posts", methods=["GET"])
def get_posts():
    posts = Post.query.all()
    return jsonify([p.to_dict() for p in posts]), 200


with app.app_context():
    db.create_all()

if __name__ == "__main__":
    app.run(debug=True)
