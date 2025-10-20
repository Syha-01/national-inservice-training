include .envrc

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application (development mode - allows all CORS origins)
.PHONY: run/api
run/api:
	@echo '--Running application (development mode)'
	@go run ./cmd/api -port=4000 -env=development -limiter-burst=5 -limiter-rps=2 -limiter-enabled=true

## run/api/prod: run the cmd/api application (production mode - restricted CORS)
.PHONY: run/api/prod
run/api/prod:
	@echo '--Running application (production mode)'
	@go run ./cmd/api -port=4000 -env=production \
		-cors-trusted-origins="https://police.gov.bz https://training.police.gov.bz"

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## test: run all tests
.PHONY: test
test:
	@echo 'Running tests...'
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	@echo 'Running tests with coverage...'
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# ==================================================================================== #
# CORS TESTING
# ==================================================================================== #

## test/cors/preflight: test CORS preflight request
.PHONY: test/cors/preflight
test/cors/preflight:
	@echo 'Testing CORS preflight request...'
	@curl -i -X OPTIONS http://localhost:4000/v1/healthcheck \
		-H "Origin: http://localhost:3000" \
		-H "Access-Control-Request-Method: GET"

## test/cors/get: test CORS GET request
.PHONY: test/cors/get
test/cors/get:
	@echo 'Testing CORS GET request...'
	@curl -i http://localhost:4000/v1/healthcheck \
		-H "Origin: http://localhost:3000"

## test/cors/post: test CORS POST request
.PHONY: test/cors/post
test/cors/post:
	@echo 'Testing CORS POST request...'
	@curl -i -X POST http://localhost:4000/v1/healthcheck \
		-H "Origin: http://localhost:3000" \
		-H "Content-Type: application/json" \
		-d '{"test": "data"}'

# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	@go build -ldflags='-s' -o=./bin/api ./cmd/api
	@GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/api ./cmd/api

# ==================================================================================== #
# DATABASE
# ==================================================================================== #

## db/psql: connect to the database using psql
.PHONY: db/psql
db/psql:
	psql ${TRAINING_DB_DSN}

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	@migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	@migrate -path ./migrations -database ${TRAINING_DB_DSN} up

## db/migrations/down: apply all down database migrations
.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running down migrations...'
	@migrate -path ./migrations -database ${TRAINING_DB_DSN} down
