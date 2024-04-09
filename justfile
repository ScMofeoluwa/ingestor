# Load .env
set dotenv-load

gen-migration m:
  migrate create -ext sql -dir internal/database/migrations -seq {{m}}
  
migrate-up:
  migrate -path internal/database/migrations -database $DATABASE_URL -verbose up

migrate-down:
  migrate -path internal/database/migrations -database $DATABASE_URL -verbose down
  
sqlc:
  sqlc generate
