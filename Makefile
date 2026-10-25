.PHONY: colima-start dev prod-build down logs clean

# Start Colima with adequate resources for M5 Air (4 CPUs, 6GB RAM)
colima-start:
	colima start --cpu 4 --memory 6 --vm-type=vz --vz-rosetta

# Spin up development environment with Docker Compose
dev:
	docker compose -f docker-compose.dev.yml up --build

# Stop containers
down:
	docker-compose down

# Tail backend logs
logs:
	docker-compose logs -f backend

# Clean docker volumes (caution: wipes database data)
clean:
	docker-compose down -v

# Clears all rows from the database tables (keeps schema)
db-clear:
	docker exec -it go_backend go run cmd/database/main.go clear

# Creates the default admin account
db-superuser:
	docker exec -it go_backend go run cmd/database/main.go make-superuser