# Makefile

VERSION=$(shell git rev-parse --verify HEAD)

glide-update:
	rm -rf ./vendor
	glide update --cache

build:
	O15VENDOREXPERIMENT=1 go build -ldflags "-X main.version=$(VERSION)" -o dist/main ./main.go