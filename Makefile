LOGFILE=$(LOGPATH) `date +'%A-%b-%d-%Y-TIME-%H-%M-%S'`
branch := $(shell git branch --show-current)

.PHONY: help
help: ## Shows help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

.PHONY: up-pgsql
up-pgsql: ## docker-compose -f deploy/postgres.yml up
	docker-compose -f deploy/postgres.yml up

.PHONY: dn-pgsql
dn-pgsql: ## docker-compose -f deploy/postgres.yml down
	docker-compose -f deploy/postgres.yml down

.PHONY: up-vault
up-vault: ## docker-compose -f deploy/vault.yml up
	docker-compose -f deploy/vault.yml up

.PHONY: dn-vault
dn-vault: ## docker-compose -f deploy/vault.yml down
	docker-compose -f deploy/vault.yml down

.PHONY: cm
cm: ## 🌱 git commit
	git add .
	git commit -m "$(branch)-${LOGFILE}"
	git push origin $(branch)