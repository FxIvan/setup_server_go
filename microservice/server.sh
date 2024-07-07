echo "Starting Ferrutx Server $1"
docker compose down

docker compose build

echo "Starting Microservice Server $1"

echo "almendraivan/microservice-coupon:$1"
docker tag microservice-app almendraivan/microservice-coupon:$1
docker push almendraivan/microservice-coupon:$1

# SSH and execute commands on the remote server
ssh root@64.23.206.1 << EOF
cd /var/www
echo "VERSION_IMAGE=$1" > .env
docker compose down
docker pull almendraivan/microservice-coupon:$1
export VERSION_IMAGE=$1
docker compose up -d
EOF