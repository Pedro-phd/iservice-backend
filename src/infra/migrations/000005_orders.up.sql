CREATE TABLE IF NOT EXISTS orders(
  id SERIAL PRIMARY KEY,
  seller_id INT NOT NULL REFERENCES sellers(id),
  client_id INT NOT NULL REFERENCES clients(id),
  status TEXT NOT NULL
);