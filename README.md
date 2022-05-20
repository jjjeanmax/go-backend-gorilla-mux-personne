# go-backend-gorilla-mux-personne

# Requirements:
>Golang, gorilla/mux, Postgres (pq), migrate, bash.

## Functionality for the system administrator:
- Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

- The name mux stands for "HTTP request multiplexer" (https://github.com/gorilla/mux)

## Configuration file
When deploying, copy the `config.json.example` file to the `config.json` file in the directory
`api`.

Example `config.json` file `api/config.json.example`.

## Start

1. Install packages:
  github.com/gorilla/mux
  github.com/lib/pq 
  https://github.com/golang-migrate/migrate for migrations


2. Create config file:

    `config.json`
    
3. Run migrations:

    `cd api`
    
    `bash migration.sh `

4. Start server:

    `go run ./`
    
    
### IN PROGRESS
