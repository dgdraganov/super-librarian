
FROM golang:1.21 AS build 

WORKDIR /librarian

COPY . .

RUN go mod download
RUN go build -o /librarian/app cmd/librarian/main.go


FROM ubuntu:22.04

WORKDIR /librarian

# todo: add certificates
COPY --from=build /librarian/app /librarian/app

EXPOSE 9205
CMD ["/librarian/app"]