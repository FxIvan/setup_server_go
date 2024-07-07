echo "Starting Ferrutx Server $1"
docker compose down

docker compose build
#TEST PARA AVERIGUAR EL WEBHOOK
echo "Starting Ferrutx Server $1"

echo "almendraivan/microservice-coupon:$1"
docker tag set_up_server-microservice almendraivan/microservice-coupon:$1
docker push almendraivan/microservice-coupon:$1

echo "almendraivan/microservice-ualabis:$1"
docker tag set_up_server-microservice-payment-ualabis almendraivan/microservice-uala:$1
docker push almendraivan/microservice-uala:$1

# SSH and execute commands on the remote server
ssh root@64.23.206.1 << EOF
cd /var/www
echo "VERSION_IMAGE=$1" > .env
docker compose down
docker pull almendraivan/microservice-coupon:$1
docker pull almendraivan/microservice-uala:$1
export VERSION_IMAGE=$1
docker compose up -d
EOF