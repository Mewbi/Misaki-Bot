CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL,
    name VARCHAR(250) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS routines (
    id VARCHAR(36) NOT NULL CHECK (LENGTH(id) > 0) ,
    name VARCHAR(250) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS subscriptions (
    user_id INT NOT NULL,
    PRIMARY KEY (id)
);