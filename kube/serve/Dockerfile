FROM python:3.8

ARG MODEL_URI
ENV MODEL_URI=$MODEL_URI

RUN mkdir -p /app/scripts

COPY kube/serve/run.sh run.sh
COPY mlruns mlruns
COPY requirements.txt requirements.txt

COPY *pip.conf /etc/pip.conf

RUN pip install \
    -r requirements.txt

ENTRYPOINT ["/usr/bin/env"]

CMD ["bash", "run.sh"]
