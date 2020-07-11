CREATE SCHEMA IF NOT EXISTS payment_system;

CREATE TABLE IF NOT EXISTS payment_system.account (
    id integer NOT NULL,
    firstname character varying(64) NOT NULL,
    lastname character varying(64) NOT NULL,
    passportdata character varying(64) NOT NULL,
    phone character varying(64) NOT NULL,
    address text
);
