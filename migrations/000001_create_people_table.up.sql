CREATE TABLE IF NOT EXISTS country(
    "country_id" bigserial PRIMARY KEY,
    "name" text NOT NULL,
    "continent" varchar(50) NOT NULL,
    "capital" varchar(50) NOT NULL
);
COMMIT;

CREATE TABLE IF NOT EXISTS person(
    "person_id" bigserial PRIMARY KEY,
    "first_name" text NOT NULL,
    "last_name" text NOT NULL,
    "birth_day" date not NULL,
    "in_life" boolean,
    "country_id" bigint NOT NULL,
    "registry" timestamp NOT NULL DEFAULT (now())
);
COMMIT;

ALTER TABLE person ADD FOREIGN KEY(country_id) REFERENCES country(country_id)
