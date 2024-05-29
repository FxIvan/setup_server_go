## Ejecutar Go cmd/app
> go run . -enginedb mongodb -portdb 27017 -userdb almendra -passwordb 1ASWWWaeq -dbname user

## Ejecutar contenedor docker
> sudo docker run -d -p 27017:27017 --name mongodb-container -e MONGO_INITDB_ROOT_USERNAME=almendra -e MONGO_INITDB_ROOT_PASSWORD="1ASWWWaeq" -v mongodb-data:/data/db mongo

### En caso de este error
>docker: Error response from daemon: Conflict. The container name "/mongodb-container" is already in use by container "c2d38b2113b05c511f85526835a2b02de0604623f66cb567976db28f4a6d5496". You have to remove (or rename) that container to be able to reuse that name.

### Hacer los siguientes pasos`
`docker stop mongodb-container`
`docker rm mongodb-container`
`Ejecutar otra vez el comando docker run ...`