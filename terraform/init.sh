#!/bin/bash
yum update -y
yum install -y git
amazon-linux-extras install docker -y
service docker start
usermod -a -G docker ec2-user

mkdir -p /home/ec2-user/Development && cd /home/ec2-user/Development
git clone https://github.com/andrewjamesmoore/andrew-projects-api
cd andrew-projects-api

docker run -d --name mongo -p 27017:27017 mongo

docker build -t graphql-api .
docker run -d --name graphql-api --link mongo -p 8080:8080 graphql-api
