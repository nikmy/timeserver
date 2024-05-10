#!/bin/bash

while(true); do
  curl "http://timeservice/statistics" > stats.txt
  sleep 5
done
