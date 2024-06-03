#!bin/bash

if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi
#Explicacion codigo bash de export
#if [ -f .env ] -> verifica si el archivo .env existe
#cat .env | sed 's/#.*//g' | xargs
#cat .env -> lee el archivo .env
#sed 's/#.*//g' -> elimina los comentarios
#xargs -> convierte la salida en argumentos
#export -> exporta las variables de entorno

#.ENV
source .env

#Environment
echo "ENV | Variables de entorno"
#Mongo
#NAME_CONTAINER_DB_MONGO = $NAME_CONTAINER_DB_MONGO
echo  $NAME_CONTAINER_DB_MONGO
#USERNAME_DB_MONGO = $USERNAME_DB_MONGO
#PASSWORD_DB_MONGO = $PASSWORD_DB_MONGO

#Mysql
#NAME_CONTAINER_DB_MYSQL = $NAME_CONTAINER_DB_MYSQL
#USERNAME_DB_MYSQL = $USERNAME_DB_MYSQL
#PASSWORD_DB_MYSQL = $PASSWORD_DB_MYSQL

#Eliminarmos contenedores activos
echo "STOP | contenedores activos"
docker stop $NAME_CONTAINER_DB_MONGO
docker stop $NAME_CONTAINER_DB_MYSQL

#Eliminamos contenedores
echo "REMOVE | contenedores"
docker rm $NAME_CONTAINER_DB_MONGO
docker rm $NAME_CONTAINER_DB_MYSQL

#Corremos los contenedores
echo "RUN | contenedores"
docker run -d -p 27017:27017 --name $NAME_CONTAINER_DB_MONGO -e MONGO_INITDB_ROOT_USERNAME=$USERNAME_DB_MONGO -e MONGO_INITDB_ROOT_PASSWORD=$PASSWORD_DB_MONGO -v mongodb-data:/data/db mongo
docker run -d -p 33060:3306 --name $NAME_CONTAINER_DB_MYSQL  -e MYSQL_ROOT_PASSWORD=$PASSWORD_DB_MYSQL --mount src=$NAME_CONTAINER_DB_MYSQL-data,dst=/var/lib/mysql mysql

#Notas
# Podemos eliminar contenedor por que tenemos un volumen que persiste la información
