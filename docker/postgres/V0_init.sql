create table log
(
    "id" serial not null
        constraint log_pkey
            primary key,
    "user" char(100) not null,
    "route" char(200) not null,
    "method" char(10) not null,
    "payload" jsonb,
    "timestamp" timestamp not null
);
alter table log owner to shorten;

drop table if exists shortener;
-- Create a table with a JSONB column.
create table if not exists shortener
(
    id bigserial not null
        constraint shortener_pk
            primary key,
    "owner" varchar(30) not null,
    "code" varchar(100) not null,
    "link" varchar(1024) not null,
    "description" varchar(250) not null,
    "count" integer,
    "maxcount" integer,
    "createdat" timestamp,
    "updatedat" timestamp,
    "starttime" timestamp,
    "expiresat" timestamp,
    attributes jsonb
);
CREATE INDEX idx_shortener_attributes ON shortener USING gin (attributes);
CREATE UNIQUE INDEX shortener_code_uindex ON shortener (code);
alter table shortener owner to shorten;

create table shortstat
(
    id bigserial not null
        constraint shortstat_pk
            primary key,
    "shortenerid" bigint not null
        constraint shortstat_shortener__fk
            references shortener
            on delete cascade,
    "clientip" varchar(15) not null,
    "useragent" varchar(200) not null,
    referer varchar(200) not null,
    createdat timestamp not null
);

alter table shortstat owner to shorten;

