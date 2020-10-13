CREATE SCHEMA IF NOT EXISTS payment_system;

CREATE TABLE IF NOT EXISTS customer (
    uid character varying(64) NOT NULL UNIQUE,
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
	birthdate date NOT NULL default NOW(),
	birthplace character varying(64) NOT NULL
);

CREATE INDEX customer_uid_idx ON customer USING btree (uid);

CREATE INDEX customer_passportnumber_idx ON customer USING btree (passportnumber);
