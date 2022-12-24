run:
	go run cmd/server/main.go

setup: ## Perform setup script, install necessary plugins/tools
	@scripts/setup.sh

gen/%: ## Generate models using sqlboiler following pre-defined templates
	# @sqlboiler psql --wipe --add-soft-deletes --templates ./db/templates --config ./db/$*.toml
	@sqlboiler psql --wipe --add-soft-deletes --config ./db/$*.toml

db/migrate: ## Migrate database structure
	@scripts/migrate.sh up

db/up: ## Apply all the migration to the latest version to the local database
	@make db/migrate

db/down: ## Remove every in the database! (only for DEV)
	@scripts/migrate.sh down 

db/drop: ## Remove every in the database! (only for DEV)
	@scripts/migrate.sh drop -f

db/connect:
	pgcli -h localhost -p 54322 -u postgres

docker/build: ## Build docker compose
	docker-compose build

docker/up: ## Run docker compose
	docker-compose -f ./docker/docker-compose.yaml up -d

docker/down: ## Stop docker compose
	docker-compose -f ./docker/docker-compose.yaml down

docker/reset:
	@make docker/down
	rm -rf docker/tmp
	@make docker/up
