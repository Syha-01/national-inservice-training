.PHONY: run/api
run/api:
	@echo '--Running aaplication'
	@go run ./cmd/api -port=4000 -env=production