
.PHONY: dev
dev:
	docker-compose -f docker-compose.yml up --build --watch

.PHONY: test
test:
	go test ./... -v
