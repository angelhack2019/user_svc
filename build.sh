#!/usr/bin/env bash

docker build -t doughyou/user_svc:latest .

docker push doughyou/user_svc:latest