create table users
(
    id        bigserial primary key,
    email     varchar unique not null,
    password  varchar        not null,
    create_at timestamptz default (now()),
    update_at timestamptz default (now())
);
