BINARY=api
build:
	CGO_ENABLED=0 go build -o bin/${BINARY} main.go

clean:
	if [ -f bin/${BINARY} ] ; then rm bin/${BINARY} ; fi

image:
	docker build -f ./deployment/Dockerfile -t go-api-skeleton .

lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

run:
	docker-compose -f deployment/development/docker-compose.yml --project-directory="$(shell pwd)" up -d

stop:
	docker-compose -f deployment/development/docker-compose.yml --project-directory="$(shell pwd)" down

serve:
	gin -i --appPort 8080 run main.go

test: 
	go test -v -cover -covermode=atomic ./...

unittest:
	go test -short ./...

watch:
	docker-compose -f deployment/development/docker-compose.yml --project-directory="$(shell pwd)" up

.PHONY: build clean docker lint lint-prepare run stop serve test unittest watch