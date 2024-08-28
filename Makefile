include ./config/.env
.DEFAULT_GOAL:= help

help:
	@echo "build -builds the project"
	@echo "run -runs the project withou building it"
	@echo "host -starts the server on virtual host"
	@echo "clean -removes the binary file"
	@echo "help - displays this help message"
	@echo "migrate_create -creates a new migration"
	@echo "migrate_up -runs the migration up"
	@echo "migrate_down -runs the migration down"

build:
	@echo "Building the project"
	go mod tidy
	go build -o ./mvc cmd/main.go
	@echo "Build complete"

run:
	@echo "Running the project"
	go run cmd/main.go

host:
	@echo "Starting the server on virtual host"
	@echo "Server started on mvc.libmansys.local"
	go run cmd/main.go

clean:
	@echo "Removing the binary file"
	rm -f ./mvc



migration_up:
		@read -p \"Enter version by which you want to up the db: \" v; \\
		migrate -path database/migration/ -database \"mysql://root:$(DB_PASS)@tcp(localhost:3306)/lib_db?\" -verbose up	$$v
migration_down:
		@read -p \"Enter version by which you want to down the db: \" v; \\
		migrate -path database/migration/ -database \"mysql://root:$(DB_PASS)@tcp(localhost:3306)/lib_db?\" -verbose down $$v
migration_fix:
		@read -p \"Enter version: \" v; \\
		migrate -path database/migration/ -database \"mysql://root:$(DB_PASS)@tcp(localhost:3306)/lib_db?\" force $$v
