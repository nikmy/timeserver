#!/bin/bash

git clone https://github.com/prometheus-operator/kube-prometheus.git

docker build -t myink/time_server:latest -f Dockerfile .
docker build -t myink/statistics_script:latest -f Dockerfile.script .
docker build -t myink/timeserver_stats_exporter:latest -f Dockerfile.metrics .

docker push myink/time_server:latest
docker push myink/statistics_script:latest
docker push myink/timeserver_stats_exporter:latest

kubectl apply -f k8s/app.yaml
kubectl apply -f k8s/job.yaml

kubectl apply -f k8s/metrics.yaml
kubectl apply --server-side -f kube-prometheus/manifests/setup