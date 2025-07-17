CREATE TYPE NameType AS ENUM ('nickname', 'alternate', 'fullname', 'translation')

CREATE TABLE IF NOT EXISTS characters_index (
  id                 SERIAL PRIMARY KEY,
  name               TEXT NOT NULL,
  aliases            JSONB,
  species            TEXT[] NOT NULL,

  origin_id          INT NOT NULL REFERENCES origins_index(id),
  franchise_id       INT NOT NULL REFERENCES franchises_index(id),
  tags               TEXT[],

  -- Added for proto only
  first_appearance   TEXT,
  appears_in         TEXT[],
  date_added         TIMESTAMPTZ
)

CREATE TABLE IF NOT EXISTS franchises_index (
  id    SERIAL PRIMARY KEY,
  name  TEXT NOT NULL,
  year  INT NOT NULL,
  owner TEXT NOT NULL
)

CREATE TABLE IF NOT EXISTS origins_index (
  id                 SERIAL PRIMARY KEY,
  name               TEXT NOT NULL,
  country_of_origin  TEXT NOT NULL,
  year               INT NOT NULL
)
