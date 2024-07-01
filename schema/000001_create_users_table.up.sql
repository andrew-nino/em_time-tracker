CREATE TABLE IF NOT EXISTS users (
    id              serial       not null unique,
    passportSerie   varchar(255) not null unique,
    passportNumber  varchar(255) not null unique,
    created_at      timestamp    not null default now()
);