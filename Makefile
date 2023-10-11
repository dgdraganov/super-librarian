

database:
	docker-compose up --detach --build db migrator

librarian:
	docker-compose up --detach --build librarian

.PHONY: gen
gen:
	goa gen github.com/dgdraganov/super-librarian/design

