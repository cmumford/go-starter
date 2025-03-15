
.PHONY: dev
dev:
	docker-compose -f docker-compose.dev.yml up --build --watch

.PHONY: prod
prod:
	docker-compose -f docker-compose.prod.yml up --build
