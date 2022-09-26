FROM python:alpine
RUN pip install Flask
RUN pip install redis
RUN pip install requests

COPY rng.py /
CMD ["python", "rng.py"]
EXPOSE 80
