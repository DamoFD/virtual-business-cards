migration:
	@goose -dir ./cmd/migrate/migrations/ create $(filter-out $@,$(MAKECMDGOALS)) sql

migrate-up:
	@docker exec -it cardapp goose -dir ./cmd/migrate/migrations postgres "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" up

migrate-down:
	@docker exec -it cardapp goose -dir ./cmd/migrate/migrations postgres "host=db port=5432 user=postgres password=postgres dbname=postgres sslmode=disable" down
