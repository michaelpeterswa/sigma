all: down run

run:
	docker compose up --build

down: 
	docker compose down