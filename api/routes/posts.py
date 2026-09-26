from datetime import datetime

from flask import Blueprint, jsonify, request
from sqlalchemy.orm import Session

from extensions import db
from models import Post

posts_bp = Blueprint("posts", __name__, url_prefix="/posts")


@posts_bp.route("", methods=["POST"])
def create_post():
    data = request.get_json()

    if not data or "title" not in data or "content" not in data:
        return jsonify({"error": "no title or content"}), 400

    new_post = Post(title=data["title"], content=data["content"])
    db.session.add(new_post)
    db.session.commit()

    return jsonify(new_post.to_dict()), 201


@posts_bp.route("<int:post_id>", methods=["DELETE"])
def delete_post_by_id(post_id):
    with Session(db.engine) as session:
        post = session.get(Post, post_id)
        if post:
            session.delete(post)
            session.commit()
        else:
            return jsonify({"error": "no post with this id"}), 404

    return jsonify({"success": "post was deleted"}), 204


@posts_bp.route("", methods=["GET"])
def get_posts():
    posts = Post.query.all()
    if posts is None:
        return jsonify({"error": "no posts"}), 404

    return jsonify([p.to_dict() for p in posts]), 200


@posts_bp.route("<int:post_id>", methods=["GET"])
def get_post_by_id(post_id):
    post = Post.query.get_or_404(post_id)

    return jsonify(post.to_dict())


@posts_bp.route("<int:post_id>", methods=["POST"])
def update_post_by_id(post_id):
    data = request.get_json()

    if not data or "title" not in data or "content" not in data:
        return jsonify({"error": "no title or content"}), 400

    with Session(db.engine) as session:
        post = session.get(Post, post_id)
        if post:
            post.title = data["title"]
            post.content = data["content"]
            post.updated_at = datetime.now(None)
            session.commit()
        else:
            return jsonify({"error": "could not update the post"}), 500

    return jsonify({"success": f"post with the id {post_id} was updated"}), 200
