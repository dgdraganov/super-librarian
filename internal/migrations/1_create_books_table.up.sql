

CREATE TABLE Books (
    Id int NOT NULL AUTO_INCREMENT,
    author varchar(100) NOT NULL,
    title varchar(255) NOT NULL,
    bookCover varchar(1024) NOT NULL,
    publishedAt DATE NOT NULL,
    PRIMARY KEY (id)
);

