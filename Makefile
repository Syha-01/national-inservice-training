.PHONY: run/api
run/api:
	@echo '--Running aaplication'
	@go run ./cmd/api -port=3000 -env=production