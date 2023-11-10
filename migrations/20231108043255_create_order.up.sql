CREATE TABLE IF NOT EXISTS orders(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    address VARCHAR NOT NULL,
    coordinate_address_x FLOAT NOT NULL,
    coordinate_address_y FLOAT NOT NULL,
    coordinate_point_x FLOAT,
    coordinate_point_y FLOAT,
    create_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    start_at TIMESTAMPTZ NULL,
    delivery_at TIMESTAMPTZ NULL,
    courier_id BIGINT REFERENCES users(id) ON DELETE SET NULL,
    delivery_status INT
);