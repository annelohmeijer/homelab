#!/bin/sh
# script to run mlflow model
mlflow models serve \
  --no-conda \
  --host 0.0.0.0 \
  --port 8000 \
  --model-uri $MODEL_URI