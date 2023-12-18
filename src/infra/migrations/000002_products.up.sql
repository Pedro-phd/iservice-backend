CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  seller_id INT REFERENCES sellers(id) NOT NULL,
  name TEXT NOT NULL,
  price MONEY NOT NULL,
  description TEXT NOT NULL,
  image TEXT NOT NULL,
  active BOOLEAN NOT NULL
);