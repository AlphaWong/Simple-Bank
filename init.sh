#!/bin/bash
docker run -v $PWD/db:/var/lib/mysql -p 3306:3306 --name bank-db -e MYSQL_ROOT_PASSWORD=hello123 -d mysql:5.6
