[![pipeline status](https://gitlab.com/AlphaWong/Simple-Bank/badges/master/pipeline.svg)](https://gitlab.com/AlphaWong/Simple-Bank/commits/master)
[![coverage report](https://codecov.io/gl/AlphaWong/Simple-Bank/branch/master/graph/badge.svg)](https://codecov.io/gl/AlphaWong/Simple-Bank)

# Test
```
go test -covermode=count ./...
```
## Start minikube
```sh
minikube start
```

## Set docker env
```sh
eval $(minikube docker-env)
```

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
mysql> CREATE DATABASE `simple-bank`;
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