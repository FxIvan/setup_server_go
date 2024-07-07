# Usa una imagen base ligera de Alpine Linux
FROM alpine:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo ejecutable en el contenedor
COPY cmd/app/micro-app .
COPY .env .

# Asegura los permisos de ejecución para el binario
RUN chmod +x micro-app

# Exponer el puerto que utiliza tu aplicación (si tu aplicación expone algún puerto)
EXPOSE 4000

# Comando de inicio para ejecutar la aplicación
CMD ["./micro-app", "-enginedb", "mongodb", "-hostdb", "mongodb-container", "-portdb", "27017", "-userdb", "almendra", "-passwordb", "1ASWWWaeq", "-dbname", "microservicio"]
