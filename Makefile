GITHUB_USER := mrofi
IMAGE_NAME := $(GITHUB_USER)/bashscript-server
TAG := latest
PORT ?= 8080
SCRIPT_DIR ?= ./scripts

.PHONY: all build run push clean

all: build

build:
	go build -o server main.go

docker-build:
	docker build --build-arg SCRIPT_DIR=$(SCRIPT_DIR) -t ghcr.io/$(IMAGE_NAME):$(TAG) .

run:
	docker run -it --rm -p $(PORT):8080 \
		-e SCRIPT_DIR=$(SCRIPT_DIR) \
		-v $(PWD)/scripts:$(SCRIPT_DIR) \
		ghcr.io/$(IMAGE_NAME):$(TAG)

push:
	docker push ghcr.io/$(IMAGE_NAME):$(TAG)

clean:
	rm -f server
