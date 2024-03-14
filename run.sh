#CLEAR 
echo "Cleaning up..."
rm -R localstack
mkdir localstack


## Building Lambdas
echo "Building Lambdas..."
cd lambdas
./script.sh
cd ..


## Create Newtwork
echo "Creating Network..."
network_name="stori-external"
network_exists=$(docker network ls | grep $network_name)

# Crear la red si no existe
if [ -z "$network_exists" ]; then
  echo "La red $network_name no existe. Creándola..."
  docker network create $network_name
else
  echo "La red $network_name ya existe."
fi

# Docker run 
echo "Running Docker..."
docker-compose stop 
docker-compose rm -f

docker-compose up -d

sleep 30
# Run Terraform 
echo "Running Terraform..."
cd aws-pipelines

terraform init
terraform plan
terraform apply -auto-approve

cd ..


#aws s3 cp ./example.csv s3://stori-bucket --region us-east-1 --endpoint-url=http://s3.localhost.localstack.cloud:4566





