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

ENV APP_ENV=production
ENV HTTP_URL=localhost
ENV HTTP_PORT=4000
ENV HTTP_ALLOWED_ORIGINS="*"
ENV PRIVATE_KEY=0af7915144c2f2f95e9f5ea995bfd1871618e711fbf58d19b3c37a744991e55c4d9f60d821f61ccf75dbff8b91916fbab3943fc13947f5ac6cb242582e4f7e7d
ENV PUBLIC_KEY=4d9f60d821f61ccf75dbff8b91916fbab3943fc13947f5ac6cb242582e4f7e7d
ENV TOKEN_DURATION=24h

# Define el punto de entrada
ENTRYPOINT ["./app", "-enginedb", "mongodb","-hostdb","mongodb-container","-portdb", "27017", "-userdb", "almendra", "-passwordb", "1ASWWWaeq", "-dbname", "microservicio"]
