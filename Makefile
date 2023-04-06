linter:
	golangci-lint run


# Migration Commands
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)


# Docker Compose Commands
start:
	docker-compose up -d

stop:
	docker-compose down

hard-stop:
	docker-compose down --volumes

restart:
	docker-compose restart
