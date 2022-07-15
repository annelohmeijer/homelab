#!/bin/sh
# script to run mlflow server
mlflow server \
  --backend-store-uri sqlite:///mlflow.db \
  --artifacts-destination artifacts --serve-artifacts \
  --host 0.0.0.0 --port 80
