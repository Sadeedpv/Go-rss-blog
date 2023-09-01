-- +goose Up

CREATE TABLE Users(
    id UUID PRIMARY KEY,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    Username TEXT NOT NULL
);


-- +goose Down

DROP TABLE Users;