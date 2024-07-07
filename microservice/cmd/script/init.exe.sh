#!/bin/bash

cd /home/almendra/set_up_server/microservice/cmd/app
#Eliminar app.exe
rm app
go build

# Desarrollo Local
./app -enginedb mongodb -hostdb 127.0.0.1 -portdb 27017 -userdb almendra -passwordb 1ASWWWaeq -dbname microservicio

#Desarrollo Contenedor Docker

