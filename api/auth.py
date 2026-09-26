import datetime
from functools import wraps

import jwt
from flask import current_app, jsonify, request


def gen_token(id):
    payload = {
        "sub": id,
        "exp": datetime.datetime.now() + datetime.timedelta(hours=8),
        "iat": datetime.datetime.now(),
    }

    return jwt.encode(payload, current_app.config["SECRET_KEY"], algorithm="HS256")


def token_required(f):
    @wraps(f)
    def decorated(*args, **kwargs):
        auth_header = request.headers.get("Authorization", "")

        if not auth_header.startswith("Bearer "):
            return jsonify({"error": "no token"}), 401

        token = auth_header.split(" ")[1]

        try:
            payload = jwt.decode(
                token, current_app.config["SECRET_KEY"], algorithms=["HS256"]
            )
            request.id = payload["sub"]
        except jwt.ExpiredSignatureError:
            return jsonify({"error": "expired token"}), 401
        except jwt.InvalidTokenError:
            return jsonify({"error": "invalid token"}), 401
        except:
            return jsonify({"error": "internal error"}), 500

        return f(*args, **kwargs)

    return decorated
