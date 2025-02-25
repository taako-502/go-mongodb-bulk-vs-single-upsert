.PHONY: db db-stop benchmark test-benchmark

db:
	docker run -d --name go_mongodb_bulk_vs_single_upsert -p 27017:27017 mongo

db-stop:
	docker stop go_mongodb_bulk_vs_single_upsert || true
	docker rm go_mongodb_bulk_vs_single_upsert || true

benchmark: db-stop
	@$(MAKE) db
	@echo "Running benchmark tests..."
	@go test -bench . -benchmem | tee benchmark_results.txt || true
	@$(MAKE) db-stop
