-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS organisations (
    id INT NOT NULL AUTO_INCREMENT,

    `name` VARCHAR(150),
    email VARCHAR(150) NOT NULL,
    github_id VARCHAR(150) NOT NULL,
    avatar_url VARCHAR(300),
    web_site VARCHAR(300),
    `description` VARCHAR(300),

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(email, github_id),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,

    `name` VARCHAR(150),
    email VARCHAR(150) NOT NULL,
    github_id VARCHAR(150) NOT NULL,
    github_token VARCHAR(300) NOT NULL,
    avatar_url VARCHAR(300),

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(email, github_id),
    PRIMARY KEY(Id)
);

CREATE TABLE IF NOT EXISTS wallets (
    id INT NOT NULL AUTO_INCREMENT,

    wallet_id VARCHAR(300) NOT NULL,
    user_id   INT NOT NULL,
    organisation_id   INT NOT NULL,
    balance INT NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id, wallet_id),
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES Users (id),
    FOREIGN KEY (organisation_id) REFERENCES Organisations (id)
);

CREATE TABLE IF NOT EXISTS articles (
    id INT NOT NULL AUTO_INCREMENT,

    `name` VARCHAR(150),
    `description` TEXT,
    quantity INTEGER,
    price INT NOT NULL,
    metadata TEXT,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id)
);

CREATE TABLE IF NOT EXISTS orders (
    id INT NOT NULL AUTO_INCREMENT,

    wallet_id VARCHAR(300),
    article_id INT NOT NULL,
    quantity INTEGER,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id),
    FOREIGN KEY (wallet_id) REFERENCES Wallets(wallet_id),
    FOREIGN KEY (article_id) REFERENCES Articles(id)
);

CREATE TABLE IF NOT EXISTS organisations_articles (
    id INT NOT NULL AUTO_INCREMENT,

    organisation_id INT NOT NULL,
    article_id INT NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id),
    FOREIGN KEY (organisation_id) REFERENCES Organisations (id),
    FOREIGN KEY (article_id) REFERENCES Articles (id)
);

CREATE TABLE IF NOT EXISTS organisations_users (
    id INT NOT NULL AUTO_INCREMENT,

    organisation_id INT NOT NULL,
    user_id INT NOT NULL,
    active BOOLEAN,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id),
    FOREIGN KEY (organisation_id) REFERENCES Organisations (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
);

CREATE TABLE IF NOT EXISTS operations (
    id INT NOT NULL AUTO_INCREMENT,

    amount INT NOT NULL,
    wallet_id VARCHAR(300),
    operation_type ENUM ('buy','receive'),
    approved       BOOLEAN,
    operation_hash TEXT,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(operation_hash),
    PRIMARY KEY(id),
    FOREIGN KEY (wallet_id) REFERENCES Wallets(wallet_id),
);

CREATE TABLE IF NOT EXISTS pictures (
    id  INT NOT NULL AUTO_INCREMENT,

    organisation_id INT NOT NULL,

    alt_text TEXT,
    original TEXT,
    small TEXT,
    medium TEXT,
    large TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(id),
    FOREIGN KEY (organisation_id) REFERENCES Organisations(id),
);

CREATE TABLE IF NOT EXISTS articles_pictures(
    id INT NOT NULL AUTO_INCREMENT,

    picture_id INT NOT NULL,
    article_id INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(id),
    FOREIGN KEY (article_id) REFERENCES articles(id),
    FOREIGN KEY (picture_id) REFERENCES pictures(id),
);

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE organisations;
DROP TABLE users;
DROP TABLE wallets;
DROP TABLE articles;
DROP TABLE orders;
DROP TABLE organisations_articles;
DROP TABLE organisations_users;
DROP TABLE operations;
DROP TABLE pictures;
DROP TABLE articles_pictures;
-- +goose StatementEnd
