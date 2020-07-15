CREATE SCHEMA IF NOT EXISTS payment_system;

CREATE SEQUENCE IF NOT EXISTS payment_system.account_id_seq;

CREATE TABLE IF NOT EXISTS payment_system.account (
    id integer NOT NULL DEFAULT nextval('payment_system.account_id_seq'),
    firstname character varying(64) NOT NULL,
    lastname character varying(64) NOT NULL,
    passportdata character varying(10) NOT NULL UNIQUE,
    phone character varying(64) NOT NULL,
    country character varying(64),
    region character varying(64),
    city character varying(64),
    street character varying(64)
);

ALTER SEQUENCE payment_system.account_id_seq
OWNED BY payment_system.account.id;
