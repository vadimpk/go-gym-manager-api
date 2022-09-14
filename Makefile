.PHONY:
.SILENT:

build:
	go build -o ./.bin/app cmd/app/main.go

run: build
	./.bin/app

swag:
	swag init -g cmd/app/main.go
	
test:
	migrate -database postgres://postgres:lz921skm0001p@localhost:5432/testing?sslmode=disable -path db/migrations up
	go test -v ./...
	migrate -database postgres://postgres:lz921skm0001p@localhost:5432/testing?sslmode=disable -path db/migrations down

endtest:
	migrate -database postgres://postgres:lz921skm0001p@localhost:5432/testing?sslmode=disable -path db/migrations down