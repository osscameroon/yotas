-- +goose Up

-- +goose StatementBegin

-- SQLINES DEMO *** egin
DROP TABLE IF EXISTS organisations;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS wallets;
DROP TABLE IF EXISTS articles;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS organisations_articles;
DROP TABLE IF EXISTS organisations_users;
DROP TABLE IF EXISTS operations;
DROP TABLE IF EXISTS pictures;
DROP TABLE IF EXISTS articles_pictures;

-- +goose StatementEnd

-- +goose Down