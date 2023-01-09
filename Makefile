fmt:
	go fmt ./...

up:
	docker-compose up -d --build

down:
	docker-compose down

.PHONY: fmt
