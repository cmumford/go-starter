
.PHONY: dev
dev:
	docker-compose -f docker-compose.yml up --build --watch

.PHONY: build
build:
	docker build -t go-starter:latest --build-arg GIT_SHA1=foobar .

.PHONY: run
run:
	docker run -p 80:8080 go-starter:latest

.PHONY: test
test:
	go test ./... -v
