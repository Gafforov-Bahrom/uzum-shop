ALTER TABLE products
ADD COLUMN description VARCHAR NOT NULL DEFAULT '',
ADD COLUMN price BIGINT NOT NULL DEFAULT 0,
ADD COLUMN count BIGINT NOT NULL DEFAULT 0
;