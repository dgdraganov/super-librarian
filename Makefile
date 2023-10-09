

migrator:
	docker-compose up --detach --build migrator

database:
	docker-compose up --detach --build db
	docker-compose up --detach --build migrator

librarian:
	docker-compose up --detach --build librarian



