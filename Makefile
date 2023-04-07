# Migration Commands
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

# Docker Commands
build:
	docker build -t toggl-cards .

start:
	docker-compose up -d

stop:
	docker-compose down

hard-stop:
	docker-compose down --volumes

restart:
	docker-compose restart
