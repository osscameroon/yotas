-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Organisations (
    id INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(150),
    email VARCHAR(150) NOT NULL,
    githubID VARCHAR(150) NOT NULL,
    avatarUrl VARCHAR(300),
    webSite VARCHAR(300),
    `description` VARCHAR(300),

    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    deletedAt TIMESTAMP,

    UNIQUE(email, githubID, walletID),
    PRIMARY KEY(Id)
);

CREATE TABLE IF NOT EXISTS Users (
    id INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(150),
    email VARCHAR(150) NOT NULL,
    githubID VARCHAR(150) NOT NULL,
    avatarUrl VARCHAR(300),

    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    deletedAt TIMESTAMP,

    UNIQUE(email, githubID),
    PRIMARY KEY(Id)
);

CREATE TABLE IF NOT EXISTS Wallets (
    id INT NOT NULL AUTO_INCREMENT,

    walletID VARCHAR(300),
    userID   INT NOT NULL,
    balance INT NOT NULL,

    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    deletedAt TIMESTAMP,

    UNIQUE(walletID),
    PRIMARY KEY(Id),
    FOREIGN KEY (userID) REFERENCES Users(userID),
    FOREIGN KEY (organisationID) REFERENCES Organisations(organisationID)
);

CREATE TABLE IF NOT EXISTS Articles (
    id INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(150),
    `description` TEXT,
    quantity INTEGER,
    yotas INT NOT NULL,
    metadata TEXT,

    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    deletedAt TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id)
);

CREATE TABLE IF NOT EXISTS Orders (
    id INT NOT NULL AUTO_INCREMENT,

    walletID VARCHAR(300),
    articleID INT NOT NULL,
    quantity INTEGER,

    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    deletedAt TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id),
    FOREIGN KEY (walletID) REFERENCES Wallets(walletID),
    FOREIGN KEY (articleID) REFERENCES Articles(articleID)
);

CREATE TABLE IF NOT EXISTS OrganisationsArticles (
    id INT NOT NULL AUTO_INCREMENT,
    organisationID INT NOT NULL,
    articleID INT NOT NULL,

    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    deletedAt TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id),
    FOREIGN KEY (organisationID) REFERENCES Organisations(organisationID),
    FOREIGN KEY (articleID) REFERENCES Articles(articleID)
);

CREATE TABLE IF NOT EXISTS OrganisationsUsers (
    id INT NOT NULL AUTO_INCREMENT,
    organisationID INT NOT NULL,
    userID INT NOT NULL,

    createdAt TIMESTAMP,
    updatedAt TIMESTAMP,
    deletedAt TIMESTAMP,

    UNIQUE(id),
    PRIMARY KEY(Id),
    FOREIGN KEY (organisationID) REFERENCES Organisations(organisationID),
    FOREIGN KEY (userID) REFERENCES Users(userID)
);

CREATE TABLE IF NOT EXISTS Operations (
    id INT NOT NULL AUTO_INCREMENT,

    yotas INT NOT NULL,
    walletID VARCHAR(300),
    operationType ENUM ('buy','receive'),
    operationHash TEXT,

    createdAt TIMESTAMP,

    UNIQUE(operationHash),
    PRIMARY KEY(Id)
    FOREIGN KEY (walletID) REFERENCES Wallets(walletID),
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