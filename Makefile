swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

infra-up:
	docker network create ppio || true && docker-compose -f infrastructure/docker-compose.yaml up -d

infra-down:
	docker-compose -f infrastructure/docker-compose.yaml down -v

lab-up:
	docker-compose up -d

lab-down:
	docker-compose down -v