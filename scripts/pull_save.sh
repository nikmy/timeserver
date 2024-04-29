#!/bin/bash

while(true); do
  curl "localhost:8080/statistics" > cache.txt
  sleep 5
done
