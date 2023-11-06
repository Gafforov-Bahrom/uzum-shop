CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    login VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    coordinate_address_x FLOAT NOT NULL,
    coordinate_address_y FLOAT NOT NULL
);