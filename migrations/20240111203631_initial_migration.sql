-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders(
    order_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255) NOT NULL UNIQUE, -- Check either track number should be unique
    entry VARCHAR(255) NOT NULL, -- Check either entry should be unique
    delivery jsonb NOT NULL,
    payment jsonb NOT NULL,
    items jsonb NOT NULL,
    locale VARCHAR(10) NOT NULL,
    internal_signature VARCHAR(255), -- Check length of the attribute
    customer_id VARCHAR(255) NOT NULL,
    delivery_service VARCHAR(255) NOT NULL,
    shard_key VARCHAR(10) NOT NULL,
    sm_id INTEGER NOT NULL,
    date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    oof_shard VARCHAR(10) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd