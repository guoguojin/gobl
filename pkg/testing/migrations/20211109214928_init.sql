-- +goose Up
-- +goose StatementBegin
create table users (
    user_name varchar(255) not null,
    age int not null
);

insert into users (user_name, age) values ('John Doe', 26), ('Jane Doe', 24);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
