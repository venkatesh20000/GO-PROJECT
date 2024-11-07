# Makefile for building and testing

build:
	go build -o go-project main.go

test:
	curl -X POST -d '{"username":"admin", "password":"password123"}' -H "Content-Type: application/json" http://15.207.19.199:8080/login
