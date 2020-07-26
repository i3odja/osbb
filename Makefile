default: help


up: ##build and run project in docker container
	@docker-compose up --build -d
.PHONY: up

down: ##stop and remove all container
	@docker-compose down --remove-orphans
.PHONY: down


help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make [target]\033[36m\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help