all: run

.PHONY: build
build:
	docker-compose build

.PHONY: run
run: build
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose down

.PHONY: runServer
runServer:
	 docker run -p 8080:8080 myapp

.PHONY: buildServer
buildServer:
	docker build -t myapp .

.PHONY: db
db:
	docker compose -f docker-compose.yml up -d db


