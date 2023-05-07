all: run

.PHONY: build
build:
	docker compose -f docker-compose.yml up -d app

.PHONY: runServer
runServer:
	 docker run -p 8080:8080 myapp

.PHONY: buildServer
buildServer:
	docker build -t myapp .

.PHONY: run
run: build
	docker run --rm -p 80:8080 myapp

.PHONY: db
db:
	docker compose -f docker-compose.yml up -d db

