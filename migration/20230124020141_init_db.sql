-- +goose Up
-- +goose StatementBegin
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE history (
    "id" SERIAL PRIMARY KEY,
    "account_uuid" uuid NOT NULL,
    "amount" NUMERIC(19, 4) NOT NULL,
    "amount_before" NUMERIC(19, 4) NOT NULL,
    "amount_after" NUMERIC(19, 4) NOT NULL,
    "date_time" timestamp NOT NULL,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE transactions (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "account_uuid" uuid NOT NULL,
    "amount" NUMERIC(19, 4) NOT NULL,
    "date_time" timestamp NOT NULL,
    "flag_error" bool default false,
    "error_detail" text,
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE account (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "created_at" timestamp NOT NULL DEFAULT (now()),
    "updated_at" timestamp NOT NULL DEFAULT (now())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS history;
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd
