USE user_service;

CREATE TABLE user (
    id CHAR(36) NOT NULL,
    user_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    pass VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
