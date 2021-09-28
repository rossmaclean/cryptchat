build:
	go build -o bin/main main.go

build-image:
	docker build -f build/Dockerfile -t rossmaclean/cryptchat:test .

run:
	go run cmd/cryptchat/Main.go