# see docs: https://kompose.io/
# make sure .env is set
set -a
source .env
set +a
kompose convert -f docker-compose.yml
sed -i '' 's/resources: {}/&\n          imagePullPolicy: Never/g' backend-deployment.yaml
sed -i '' 's/spec:/&\n  type: NodePort/g' backend-service.yaml