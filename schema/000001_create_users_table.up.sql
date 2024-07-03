CREATE TABLE IF NOT EXISTS managers (
    id            serial       primary key,
    name          varchar(255) not null,
    managername   varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          varchar(7)   not null default 'manager',
    created_at    timestamp    not null default now(),

    CONSTRAINT role_manager CHECK (role IN ('admin', 'manager'))
);

CREATE TABLE IF NOT EXISTS people (
    id              serial       primary key,
    manager_id      integer      not null references managers(id),
    surname         varchar(255) not null default '',
    name            varchar(255) not null default '',
    patronymic      varchar(255) not null default '',
    address         varchar(255) not null default '',
    passport_serie  varchar(255) not null,
    passport_number varchar(255) not null,
    created_at      timestamp    not null default now()
);

CREATE UNIQUE INDEX people_passport_idx ON people (passport_serie, passport_number);

CREATE TABLE IF NOT EXISTS tasks (
    id              serial       primary key,
    importance      varchar(4)   not null default 'low',
    status          varchar(10)  not null default 'planed',
    description     text,

    CONSTRAINT importance_task CHECK (importance IN ('high', 'low'))
    CONSTRAINT status_task CHECK (status IN ('planed', 'accepted', 'completed'))
);


CREATE TABLE IF NOT EXISTS people_tasks (
    id           serial       primary key,
    person_id    integer      not null references people(id) on delete cascade ,
    task_id      integer      not null references tasks(id)  on delete cascade ,
);

CREATE TABLE IF NOT EXISTS tracker (
    id           serial       primary key,
    task_id      integer      not null references tasks(id) on delete cascade,
    created_at   timestamp without time zone default now(),
    finished_at  timestamp without time zone,
);