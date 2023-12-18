CREATE TABLE IF NOT EXISTS sellers_categories(
  id SERIAL PRIMARY KEY,
  seller_id INT NOT NULL REFERENCES sellers(id),
  category_id  INT NOT NULL REFERENCES categories(id)
);