[^1]: go build -o micro-app

> esto copiarlo en el archivo donde se encuentra la ubicacion de dockerfile
> [^2]:

    Compilación del binario para Alpine Linux
    Go para que sea compatible con Alpine Linux, que usa musl libc en lugar de glibc. Ejecuta este comando desde la raíz de tu proyecto:

`GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cmd/app/micro-app ./cmd/app`

Este comando compilará tu aplicación Go para la arquitectura amd64 y el sistema operativo linux, sin dependencias de C (CGO_DISABLED=0).

[^3]: Como quedara el Dockerfile

```
# Usa una imagen base ligera de Alpine Linux desde un espejo local
FROM mirror.gcr.io/library/alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo ejecutable en el contenedor
COPY cmd/app/micro-app .

# Asegura los permisos de ejecución para el binario
RUN chmod +x micro-app

# Exponer el puerto que utiliza tu aplicación (si tu aplicación expone algún puerto)
EXPOSE 4000

# Comando de inicio para ejecutar la aplicación
CMD ["./micro-app", "-enginedb", "mongodb", "-hostdb", "mongo", "-portdb", "27017", "-userdb", "almendra", "-passwordb", "1ASWWWaeq", "-dbname", "microservicio"]
```

[^4]: Comandos a Ejecutar

> docker build -t my_microservice_image .
> docker run -d --name my_microservice_container -p 4000:4000 my_microservice_image
