CREATE TABLE IF NOT EXISTS managers (
    id            serial       not null unique,
    name          varchar(255) not null,
    managername   varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          varchar(7)   not null default 'manager',
    created_at    timestamp    not null default now(),

    CONSTRAINT role_manager CHECK (role IN ('admin', 'manager'))
);