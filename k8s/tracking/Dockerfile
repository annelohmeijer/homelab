FROM python:3.8

# Make RUN commands use `bash --login`
SHELL ["/bin/bash", "--login", "-c"]

COPY *pip.conf /etc/pip.conf

RUN pip install mlflow==1.26.1

COPY kube/tracking_server/run.sh run.sh
RUN mkdir -p artifacts

ENTRYPOINT ["/usr/bin/env"]

CMD ["bash", "run.sh"]
