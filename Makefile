build:
	docker build -t app ./

launch: build
	docker compose -f ./cmd/docker-compose.yml -p localtesting up