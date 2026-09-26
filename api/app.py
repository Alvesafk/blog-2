from flask import Flask

from extensions import db
from routes.posts import posts_bp


def create_app():
    app = Flask(__name__)
    app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///app.db"

    db.init_app(app)
    app.register_blueprint(posts_bp)

    with app.app_context():
        db.create_all()

    return app


if __name__ == "__main__":
    app = create_app()
    app.run(debug=True)
