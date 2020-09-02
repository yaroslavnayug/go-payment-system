CREATE SCHEMA IF NOT EXISTS payment_system;

CREATE SEQUENCE IF NOT EXISTS payment_system.customer_uid_seq;

CREATE TABLE IF NOT EXISTS payment_system.customer (
    uid integer NOT NULL DEFAULT nextval('payment_system.customer_uid_seq'),
    generatedid character varying(64) NOT NULL UNIQUE,
    firstname character varying(64) NOT NULL,
    lastname character varying(64) NOT NULL,
    email character varying(64),
    phone character varying(64) NOT NULL,
    country character varying(64) NOT NULL,
    region character varying(64) NOT NULL,
    city character varying(64) NOT NULL,
    street character varying(64) NOT NULL,
    building character varying(10) NOT NULL,
    passportnumber character varying(10) NOT NULL UNIQUE,
	passportissuedate date NOT NULL,
	passportissuer character varying(255) NOT NULL,
	birthdate date NOT NULL,
	birthplace character varying(64) NOT NULL
);

ALTER SEQUENCE payment_system.customer_uid_seq
OWNED BY payment_system.customer.uid;
