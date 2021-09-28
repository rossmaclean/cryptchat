build:
	go build -o bin/cryptchat cmd/cryptchat/Main.go

build-image:
	docker build -f build/Dockerfile -t rossmaclean/cryptchat:test .

barf:
	docker build -f build/Dockerfile -t rossmaclean/cryptchat:test .
	docker stop cc || true # Pipe to true so it doesn't fail if container doesn't exists
	docker rm cc || true # ^
	docker run -d -p 8000:8000 --name cc rossmaclean/cryptchat:test

push-private:
	docker image tag rossmaclean/cryptchat:test registry.rossmac.co.uk/rossmaclean/cryptchat:test
	docker image push registry.rossmac.co.uk/rossmaclean/cryptchat:test

exec:
	docker exec -it cc /bin/sh

run:
	go run cmd/cryptchat/Main.go