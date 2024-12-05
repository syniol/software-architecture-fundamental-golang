build:
	docker-compose -f deploy/docker/docker-compose.yml build --no-cache

up:
	docker-compose -f deploy/docker/docker-compose.yml up -d

down:
	docker-compose -f deploy/docker/docker-compose.yml down

down-dev:
	docker-compose -f deploy/docker/docker-compose.yml down --remove-orphans --volumes
