cd /home/almendra/set_up_server/microservice/cmd/app

go build -o micro-app

cd /home/almendra/set_up_server/microservice

docker compose down

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o cmd/app/micro-app ./cmd/app

docker compose build

docker tag microservice-microservice-main almendraivan/microservice-coupon:$1
docker push almendraivan/microservice-coupon:$1

# SSH and execute commands on the remote server
ssh root@64.23.206.1 << EOF
cd /var/www/microservice-coupon
echo "VERSION_IMAGE=$1" > .env
docker compose down
docker pull almendraivan/microservice-coupon:$1
export VERSION_IMAGE=$1
docker compose up -d
EOF