all: run

.PHONY: build
build:
	 docker build -t app .

.PHONY: run
run: build
	docker compose -f docker-compose.yml up -d --build app

.PHONY: db
db:
	docker compose -f docker-compose.yml up -d db