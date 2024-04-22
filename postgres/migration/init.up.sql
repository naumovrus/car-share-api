create table "user"
(
    id       bigserial
        unique,
    uuid     varchar(512) not null
        primary key,
    email    varchar(255) not null
        unique,
    password varchar(255) not null
);



create table user_credentials
(
    id          bigserial
        unique,
    client_uuid varchar(512) not null
        unique
        references "user"
            on delete cascade,
    is_verifyed boolean default false,
    data        jsonb        not null
        unique
);


create table car
(
    id   bigserial
        unique,
    uuid varchar(512) not null
        primary key,
    data jsonb        not null
);



create table status
(
    id   serial
        unique,
    name varchar(55)
);


create table "order"
(
    id          bigserial
        unique,
    uuid        varchar(512) not null
        primary key,
    car_uuid    varchar(512) not null
        unique
        references car
            on delete cascade,
    client_uuid varchar(512) not null
        unique
        references "user"
            on delete cascade,
    status_id   integer
        references status (id)
            on delete cascade
);

