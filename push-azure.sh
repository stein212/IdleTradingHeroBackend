#!/bin/bash

docker build -t idle-trading-hero-server .
docker tag idle-trading-hero-server idletradinghero.azurecr.io/idle-trading-hero-server
docker push idletradinghero.azurecr.io/idle-trading-hero-server