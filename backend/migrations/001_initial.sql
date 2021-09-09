-- +goose Up

-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS organisations_seq;
CREATE TABLE IF NOT EXISTS organisations (
    id INTEGER NOT NULL DEFAULT NEXTVAL ('organisations_seq'),

    name VARCHAR(150),
    email VARCHAR(150) NOT NULL UNIQUE,
    github_id VARCHAR(150) NOT NULL UNIQUE,
    avatar_url VARCHAR(300),
    web_site VARCHAR(300),
    description VARCHAR(300),

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),

    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS users_seq;
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL DEFAULT NEXTVAL ('users_seq'),

    name VARCHAR(150),
    email VARCHAR(150) NOT NULL UNIQUE,
    github_id VARCHAR(150) NOT NULL UNIQUE,
    github_token VARCHAR(300) NOT NULL,
    avatar_url VARCHAR(300),

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),

    PRIMARY KEY(id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS wallets_seq;
CREATE TABLE IF NOT EXISTS wallets (
    id INTEGER NOT NULL DEFAULT NEXTVAL ('wallets_seq'),

    wallet_id VARCHAR(300) UNIQUE NOT NULL,
    user_id   INTEGER NOT NULL,
    organisation_id   INTEGER NOT NULL,
    balance INTEGER NOT NULL,

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),


    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (organisation_id) REFERENCES organisations (id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS articles_seq;
CREATE TABLE IF NOT EXISTS articles (
    id INTEGER NOT NULL DEFAULT NEXTVAL ('articles_seq'),

    name VARCHAR(150),
    description TEXT,
    quantity INTEGER,
    price INTEGER NOT NULL,
    metadata TEXT,

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),

    PRIMARY KEY(id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS orders_seq;
CREATE TABLE IF NOT EXISTS orders (
    id INTEGER NOT NULL DEFAULT NEXTVAL ('orders_seq'),

    wallet_id VARCHAR(300) NOT NULL,
    total_amount INTEGER NOT NULL,
    state VARCHAR(300) NOT NULL,
    decision TEXT,

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),


    PRIMARY KEY(id),
    FOREIGN KEY (wallet_id) REFERENCES wallets(wallet_id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS orders_articles_seq;
CREATE TABLE IF NOT EXISTS orders_articles(
    id INTEGER NOT NULL DEFAULT NEXTVAL ('orders_articles_seq'),

    order_id INTEGER NOT NULL,
    article_id INTEGER NOT NULL,
    article_price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),

    PRIMARY KEY(id),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (article_id) REFERENCES articles(id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS organisations_articles_seq;
CREATE TABLE IF NOT EXISTS organisations_articles (
    id INTEGER NOT NULL DEFAULT NEXTVAL ('organisations_articles_seq'),

    organisation_id INTEGER NOT NULL,
    article_id INTEGER NOT NULL,

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),


    PRIMARY KEY(id),
    FOREIGN KEY (organisation_id) REFERENCES organisations (id),
    FOREIGN KEY (article_id) REFERENCES articles (id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS organisations_users_seq;
CREATE TABLE IF NOT EXISTS organisations_users (
    id INTEGER NOT NULL DEFAULT NEXTVAL ('organisations_users_seq'),

    organisation_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    active BOOLEAN,

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),


    PRIMARY KEY(id),
    FOREIGN KEY (organisation_id) REFERENCES organisations (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS operations_seq;
CREATE TABLE IF NOT EXISTS operations (
    id INTEGER NOT NULL DEFAULT NEXTVAL ('operations_seq'),

    amount INTEGER NOT NULL,
    description TEXT,
    wallet_id VARCHAR(300),
    operation_type TEXT,
    approved       BOOLEAN,
    operation_hash TEXT UNIQUE,

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),


    PRIMARY KEY(id),
    FOREIGN KEY (wallet_id) REFERENCES wallets(wallet_id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS pictures_seq;
CREATE TABLE IF NOT EXISTS pictures (
    id  INTEGER NOT NULL DEFAULT NEXTVAL ('pictures_seq'),

    organisation_id INTEGER NOT NULL,

    alt_text TEXT,
    original TEXT,
    small TEXT,
    medium TEXT,
    large TEXT,

    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),


    PRIMARY KEY(id),
    FOREIGN KEY (organisation_id) REFERENCES organisations(id)
);
-- +goose StatementEnd


-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS articles_pictures_seq;
CREATE TABLE IF NOT EXISTS articles_pictures(
    id INTEGER NOT NULL DEFAULT NEXTVAL ('articles_pictures_seq'),

    picture_id INTEGER NOT NULL,
    article_id INTEGER NOT NULL,
    created_at TIMESTAMP(0),
    updated_at TIMESTAMP(0),
    deleted_at TIMESTAMP(0),


    PRIMARY KEY(id),
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (picture_id) REFERENCES pictures(id)
);
-- +goose StatementEnd

-- DROP TABLE IF EXISTS organisations;
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS wallets;
-- DROP TABLE IF EXISTS articles;
-- DROP TABLE IF EXISTS orders;
-- DROP TABLE IF EXISTS organisations_articles;
-- DROP TABLE IF EXISTS organisations_users;
-- DROP TABLE IF EXISTS operations;
-- DROP TABLE IF EXISTS pictures;
-- DROP TABLE IF EXISTS articles_pictures;
-- +goose Down
