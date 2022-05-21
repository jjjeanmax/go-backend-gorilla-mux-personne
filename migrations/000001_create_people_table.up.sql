CREATE TABLE IF NOT EXISTS country(
    "country_id" SERIAL PRIMARY KEY,
    "name_country" varchar(50) NOT NULL UNIQUE,
    "continent" varchar(50) NOT NULL,
    "capital" varchar(50) NOT NULL
);
COMMIT;

CREATE TABLE IF NOT EXISTS person(
    "person_id" SERIAL PRIMARY KEY,
    "first_name" text NOT NULL,
    "last_name" text NOT NULL,
    "birth_day" date not NULL,
    "in_life" boolean,
    "country_id" bigint NOT NULL,
    "registry" timestamp with time zone NOT NULL DEFAULT (now()),
    UNIQUE("first_name","last_name","birth_day")
);
COMMIT;

ALTER TABLE person ADD FOREIGN KEY(country_id) REFERENCES country(country_id) on delete CASCADE;

CREATE SEQUENCE country_sequence
  start 1
  increment 1;

CREATE SEQUENCE persone_sequence
  start 1
  increment 1;