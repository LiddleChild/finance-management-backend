ROOT_PATH=$$(pwd)

dev:
	@nodemon --exec go run main.go --signal SIGTERM

build:
	go build -o $(ROOT_PATH)/bin/server $(ROOT_PATH)/main.go

start:
	@./bin/server

# up:
# 	docker-compose up 

# down:
# 	docker-compose down

# bup:
# 	docker-compose up --build