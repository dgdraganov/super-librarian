# 
API for CRUD operations with Books.

# super-librarian

A simple API for CRUD operations on books. In order to achieve that the service exposes 5 endpoints:

GET     /book/{id}
GET     /books/{skip}/{take}
POST    /book
PATCH   /book
DELETE  /book/{id}


## How to run?

The first step is to prepare the database:

```
    make database
```

The above make target will first start the `db` service from the `docker-compose.yaml` file and then will run the `migrator` service. The role of the `migrator` is to make sure the needed `Books` table is created and ready to be used. After the `db` is up and running the following make target will start the service itself - `super-librarian`:

```
    make librarian
```

The `docker` setup uses `dev.env` file to lead the needed environment variables. Service will be run by default (from `dev.env` file) on port `9205`.

## How to use?

Obviously requests to the above endpoints can be sent from anywhere. For ease of use there is a Postman exported collection added to the project - `Super-Librarian.postman_collection.json`

## What about tests?

No tests. :()

