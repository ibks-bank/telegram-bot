-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id       bigserial primary key not null,
    username text unique           not null check ( username != '' ),
    token    text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
