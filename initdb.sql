\c food_encyclopedia;

CREATE TABLE food (
  id SERIAL PRIMARY KEY,
  species VARCHAR(256),
  description VARCHAR(1024)
);