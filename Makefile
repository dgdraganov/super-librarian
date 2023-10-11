

# migrator:
# 	docker-compose up --detach --build migrator

database:
	docker-compose up --build db migrator

# docker-compose up --detach --build migrator

librarian:
	docker-compose up --detach --build librarian

.PHONY: gen
gen:
	goa gen github.com/dgdraganov/super-librarian/design

