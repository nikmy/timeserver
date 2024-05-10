#!/bin/bash

while(true); do
  curl "http://timeserver/statistics" > stats.txt
  sleep 5
done
