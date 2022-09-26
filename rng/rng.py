import logging
from redis import Redis
from flask import Flask, Response
import os
import socket
import time

DEBUG = os.environ.get("DEBUG", "").lower().startswith("y")

log = logging.getLogger(__name__)
if DEBUG:
    logging.basicConfig(level=logging.DEBUG)
else:
    logging.basicConfig(level=logging.INFO)
    logging.getLogger("requests").setLevel(logging.WARNING)

app = Flask(__name__)

# Enable debugging if the DEBUG environment variable is set and starts with Y
app.debug = os.environ.get("DEBUG", "").lower().startswith('y')

hostname = socket.gethostname()

urandom = os.open("/dev/urandom", os.O_RDONLY)

redis = Redis("redis")

@app.route("/")
def index():
    return "RNG running on {}\n".format(hostname)


@app.route("/<int:how_many_bytes>")
def rng(how_many_bytes):
    # Simulate a little bit of delay
    time.sleep(0.1)
    random_bytes = os.read(urandom, how_many_bytes)
    created = redis.hset("cache", "current", random_bytes)
    log.info("hset bytes: {}".format(created))
    return Response(
        random_bytes,
        content_type="application/octet-stream")


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=80)

