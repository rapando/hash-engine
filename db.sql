create table passwords (
    id bigint unsigned primary key,
    protocol enum('md5', 'sha1', 'sha2', 'sha256'),
    raw_password varbinary(512) not null,
    hashed_password varbinary(512) not null,
    created datetime(6) not null default current_timestamp,
    unique key(`raw_password`, `hashed_password`)
);