-- +goose Up
-- +goose StatementBegin
create table users (
  uuid UUID primary key,
  first_name varchar(255),
  second_name varchar(255),
  email varchar(255) not null unique,
  created_at timestamp,
  deleted_at timestamp,
  updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users cascade;
-- +goose StatementEnd
