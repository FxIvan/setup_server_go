# Imagen base
FROM golang:1.22-alpine AS builder

# Instala git
RUN apk add --no-cache git

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del proyecto
COPY . .

# Configura Go Modules (si es necesario)
RUN go mod tidy

# Descarga las dependencias
RUN go get -d -v ./...

# Compila la aplicación (ajusta la ruta a tu archivo principal)
RUN go build -o /go/bin/app -v ./cmd/app

# Imagen final
FROM alpine:latest

# Establece el directorio de trabajo en la nueva imagen
WORKDIR /root/

# Copia el binario de la imagen de compilación
COPY --from=builder /go/bin/app .

# Configura el usuario
RUN adduser -D -u 1000 appuser

# Establece el usuario no root
USER appuser

# Define el punto de entrada
#MongoDB
#ENTRYPOINT ["./app", "-enginedb", "mongodb","-hostdb","mongodb-container","-portdb", "27017", "-userdb", "almendra", "-passwordb", "1ASWWWaeq", "-dbname", "microservicio"]
#MySQL
ENTRYPOINT ["./app", "-enginedb", "mysql" ,"-hostdb","micoservice-mysql-microservice-mysql.i.aivencloud.com","-portdb", "11808", "-userdb", "avnadmin", "-passwordb", "AVNS_dezgF0XPVi58xCFBl6b"] 
