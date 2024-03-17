.PHONY: db db-stop run benchmark

db:
	docker run -d --name mongodb -p 27017:27017 mongo

db-stop:
	docker stop mongodb || true
	docker rm mongodb || true

run:
	@$(MAKE) db
	@echo "Running main.go..."
	@go run main.go || true
	@$(MAKE) db-stop

benchmark:
	@$(MAKE) db
	@echo "Running main.go..."
	@go run command/benchmark/main.go || true
	@$(MAKE) db-stop
