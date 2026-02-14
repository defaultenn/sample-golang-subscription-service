-- +goose Up
-- +goose StatementBegin
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS subscriptions_id_seq;

-- Table Definition
CREATE TABLE "public"."subscriptions" (
    "id" int8 NOT NULL DEFAULT nextval('subscriptions_id_seq'::regclass),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "service_name" text,
    "price" int8,
    "user_id" text,
    "start_date" timestamptz,
    "end_date" timestamptz,
    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."subscriptions";

DROP SEQUENCE IF EXISTS subscriptions_id_seq;
-- +goose StatementEnd
