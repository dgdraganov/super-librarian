

compose:
	docker-compose up --build

compose-detached:
	docker-compose up --build

migrator:
	docker-compose up --detach --build migrator

database:
	docker-compose up --detach --build db

librarian:
	docker-compose up --detach --build librarian



