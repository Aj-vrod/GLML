# GLML

Girls' love media library

Original idea by @mariblan

## Local Development

1. Start a postgres container using Docker

   `docker-compose up -d`

2. Export the DB_DSN env variable

   `export DB_DSN="postgres://postgres:glml-postgres@localhost:5432/glml?sslmode=disable"`

3. Run the server

   `go run main.go`
