#!/bin/bash


docker-compose up -d

docker exec -it infra_roach1_1 ./cockroach init --insecure
