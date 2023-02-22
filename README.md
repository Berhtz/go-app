# Auth system 
## Version. 0.1.0

## Installation:
1. Clone repository
2. Add .env file with db data in this format:
    PG_PASS = "password"
    PG_DB = "database_name"
    PG_USER = "username"
    PG_HOST = "localhost"
    PG_PORT = "5432"
    Postgres_URL="postgres://${PG_USER}:${PG_PASS}@${PG_HOST}:${PG_PORT}/${PG_DB}"
3. Edit .env file
4. Run cd cmd/auth
5. Run go run main.go
