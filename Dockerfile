
FROM golang:1.21 AS build 

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./bin/librarian cmd/librarian/*.go \
 && go build -o ./bin/migrator cmd/migrator/*.go

FROM ubuntu:22.04

WORKDIR /bin

COPY --from=build /app/bin/librarian ./librarian
COPY --from=build /app/bin/migrator ./migrator
COPY --from=build /app/internal/migrations ./migrations

EXPOSE 9205
CMD ["/bin/librarian"]