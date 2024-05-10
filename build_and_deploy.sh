#!/bin/bash

docker build -t myink/time_server:latest -f Dockerfile .
docker build -t myink/statistics_script:latest -f Dockerfile.script .

docker push myink/time_server:latest
docker push myink/statistics_script:latest

kubectl apply -f k8s/app.yaml
kubectl apply -f k8s/job.yaml
