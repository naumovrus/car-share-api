CREATE TABLE "user" if not exists (
    id bigserial not null UNIQUE,
    uuid varchar(512) PRIMARY KEY not null UNIQUE,
    email varchar(255) not null unique,
    password varchar(255) not null,
)

