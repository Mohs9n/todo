-- +goose Up

CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  text TEXT NOT NULL,
  done BOOLEAN NOT NULL,
  user_id INTEGER NOT NULL REFERENCES users(id)
);

-- +goose Down
DROP TABLE todos;