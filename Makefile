.PHONY:
.SILENT:

build:
	go build -o ./.bin/app cmd/app/main.go

run: build
	./.bin/app

swag:
	swag init -g cmd/app/main.go