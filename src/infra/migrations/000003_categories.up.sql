DROP TABLE IF EXISTS categories;

CREATE TABLE IF NOT EXISTS categories (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

INSERT INTO "categories" ("id", "name") VALUES
(1,	'pintor'),
(2,	'eletricista'),
(3,	'encanador'),
(4,	'carpinteiro'),
(5,	'pedreiro'),
(6,	'jardineiro'),
(7,	'marceneiro'),
(8,	'mecanico'),
(9,	'borracheiro'),
(10,	'limpeza'),
(11,	'arquiteto'),
(13,	'serralheiro'),
(14,	'informatica'),
(15,	'eletronica'),
(16,	'baba');