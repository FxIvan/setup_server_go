#!/bin/bash

# Cambia el directorio a /home/almendra/set_up_server/microservice
cd /home/almendra/set_up_server/microservice

# Toma el primer parámetro como la versión
version=$1

# Concatenar el nombre del contenedor con la versión
name_container="NAME_CONTAINER_ENV=microservice-v$version"

# Escribe el nombre del contenedor en el archivo nameContainer.txt
cd /home/almendra/script
echo "$name_container" >> .env
cd /home/almendra/set_up_server/microservice
docker compose down
docker compose rm
docker rmi $(docker images -q)

# Construye la imagen del microservicio
docker compose build --no-cache > log-docker-compose-build.log
#Sacar los log dentro del log-dockerCompose.log
docker compose up --build -d > log-docker-compose-up.log
