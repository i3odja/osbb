up: ##build and run project in docker container
	@docker-compose up --build -d
down: ##stop and remove all container
	@docker-compose down --remove-orphans