# Test
```
go test -covermode=count ./...
```
## Start minikube
minikube start

## Set docker env
eval $(minikube docker-env)

## Build image
```sh
docker build . -t bank:0.0.1
```
## Init db
```sh
kubectl create -f ./mysql-deployment.yaml
```

## Create database
```
kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -phello123
```
```mysql
CREATE DATABASE `simple-bank`;
```

# Start service
```sh
kubectl run simple-bank --env="GOPATH=/go" --image=bank:0.0.1 --replicas=2 --port=3000 --image-pull-policy=Never -- --host=mysql
kubectl port-forward deployment.apps/simple-bank 3000:3000
```

## Postman collection
https://www.getpostman.com/collections/23256aad01e9937c001d

# Limition
1. DB time type stores as mysql `DATETIME`

# Report issue
alpha.wong@tuta.io