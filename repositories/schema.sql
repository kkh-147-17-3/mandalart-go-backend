create table users
(
    id              serial
        primary key,
    social_id varchar,
    social_provider varchar(5)
);

create table sheets
(
    id          serial
        primary key,
    owner_id    integer
        references users,
    name        varchar,
    created_at  timestamp,
    modified_at timestamp
);

create table cells
(
    id           serial
        primary key,
    sheet_id     integer
        references sheets,
    goal         varchar,
    color        varchar,
    step         integer not null,
    "order"      integer not null,
    parent_id    integer
        references cells,
    is_completed boolean not null,
    created_at   timestamp,
    modified_at  timestamp,
    owner_id     integer
        references users
);

create table todos
(
    id          serial
        primary key,
    owner_id    integer
        references users,
    cell_id     integer
        references cells,
    content     varchar,
    created_at  timestamp,
    modified_at timestamp
);
