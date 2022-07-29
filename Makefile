migrate:
	migrate -path ./migration -database 'postgres://postgres:postgres@0.0.0.0:5430/postgres?sslmode=disable' up