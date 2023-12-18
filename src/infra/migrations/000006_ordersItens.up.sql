CREATE TABLE IF NOT EXISTS orders_items(
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL REFERENCES orders(id),
  product_id INT NOT NULL REFERENCES products(id),
  quantity INT NOT NULL,
  price MONEY NOT NULL
);