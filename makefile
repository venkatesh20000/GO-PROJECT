# Makefile for building and testing

build:
	go build -o go-project main.go

test:
	curl -X POST -d '{"username":"admin", "password":"password123"}' -H "Content-Type: application/json" http://loocalhost/login
