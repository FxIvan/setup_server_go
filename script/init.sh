#!/bin/bash

echo microservice:$1 >> nameContainer.txt
cd /home/almendra/set_up_server/microservice
docker build -t almendraivan/microservice:v$1 .
docker run -p 4000:4000 --name microservice-v$1 --network service-mongodb almendraivan/microservice:v$1
