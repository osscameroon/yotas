-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Organisations (
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
    PRIMARY KEY(Id)
);

CREATE TABLE IF NOT EXISTS Users (
    id INT NOT NULL AUTO_INCREMENT,

    `name` VARCHAR(150),
    email VARCHAR(150) NOT NULL,
    github_id VARCHAR(150) NOT NULL,
    avatar_url VARCHAR(300),

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(email, github_id),
    PRIMARY KEY(Id)
);

CREATE TABLE IF NOT EXISTS Wallets (
    id INT NOT NULL AUTO_INCREMENT,

    wallet_id VARCHAR(300) NOT NULL,
    user_id   INT NOT NULL,
    organisation_id   INT NOT NULL,
    balance INT NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id, wallet_id),
    PRIMARY KEY(Id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id),
    FOREIGN KEY (organisation_id) REFERENCES Organisations(organisation_id)
);

CREATE TABLE IF NOT EXISTS Articles (
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

CREATE TABLE IF NOT EXISTS Orders (
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
    FOREIGN KEY (article_id) REFERENCES Articles(article_id)
);

CREATE TABLE IF NOT EXISTS OrganisationsArticles (
    id INT NOT NULL AUTO_INCREMENT,

    organisation_id INT NOT NULL,
    article_id INT NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id),
    FOREIGN KEY (organisation_id) REFERENCES Organisations(organisation_id),
    FOREIGN KEY (article_id) REFERENCES Articles(article_id)
);

CREATE TABLE IF NOT EXISTS OrganisationsUsers (
    id INT NOT NULL AUTO_INCREMENT,

    organisation_id INT NOT NULL,
    user_id INT NOT NULL,

    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id),
    FOREIGN KEY (organisation_id) REFERENCES Organisations(organisation_id),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

CREATE TABLE IF NOT EXISTS Operations (
    id INT NOT NULL AUTO_INCREMENT,

    amount INT NOT NULL,
    wallet_id VARCHAR(300),
    operationType ENUM ('buy','receive'),
    approuved BOOLEAN,
    operationHash TEXT,

    created_at TIMESTAMP,

    UNIQUE(operationHash),
    PRIMARY KEY(Id)
    FOREIGN KEY (wallet_id) REFERENCES Wallets(wallet_id),
);

-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE Organisations;
DROP TABLE Users;
DROP TABLE Wallets;
DROP TABLE Articles;
DROP TABLE Orders;
DROP TABLE OrganisationsArticles;
DROP TABLE OrganisationsUsers;
DROP TABLE Operations;
-- +goose StatementEnd