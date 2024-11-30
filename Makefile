build:
	docker-compose -f deploy/docker/docker-compose.yml build

up:
	docker-compose -f deploy/docker/docker-compose.yml up -d

down:
	docker-compose -f deploy/docker/docker-compose.yml down --remove-orphans
