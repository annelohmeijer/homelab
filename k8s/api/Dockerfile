FROM python:3.8

COPY k8s/api/main.py main.py

COPY *pip.conf /etc/pip.conf

RUN pip install \
    fastapi uvicorn

ENTRYPOINT ["/usr/bin/env"]

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "80"]
