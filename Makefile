.PHONY: db db-stop run benchmark

db:
	docker run -d --name go_mongodb_bulk_vs_single_upsert -p 27017:27017 mongo

db-stop:
	docker stop mongodb || true
	docker rm mongodb || true

run: db-stop
	@$(MAKE) db
	@echo "Running main.go..."
	@go run main.go || true
	@$(MAKE) db-stop

benchmark: db-stop
	@$(MAKE) db
	@echo "Running main.go..."
	@go run command/benchmark/main.go || true
	@$(MAKE) db-stop
