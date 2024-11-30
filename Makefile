run:
	go run cmd/server/main.go

swagger:
	swag init --output ./docs --generalInfo ./cmd/server/main.go

migrate-up:
	migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path internal/db/migrations up

migrate-down:
	migrate -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path internal/db/migrations down

tidy:
	go mod tidy

test:
	go test ./...
