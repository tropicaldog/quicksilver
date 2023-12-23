.PHONY: build

version ?= latest

build:
	go build -o ../build/icq-relayer

docker:
	docker build -f Dockerfile . -t quicksilverzone/interchain-queries:${version}

dockerx:
	docker buildx build --platform linux/amd64 -f Dockerfile . -t quicksilverzone/interchain-queries:${version}
	
docker-push:
	docker push quicksilverzone/interchain-queries:${version}
