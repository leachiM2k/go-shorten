create table log
(
    id serial not null
        constraint log_pkey
            primary key,
    pnum char(8) not null,
    route char(200) not null,
    method char(10) not null,
    payload jsonb,
    timestamp timestamp not null
);
alter table log owner to shorten;

drop table if exists shortener;
-- Create a table with a JSONB column.
create table if not exists shortener
(
    id bigserial not null
        constraint shortener_pk
            primary key,
    "owner" char(20) not null,
    "code" char(100) not null,
    "count" integer,
    "maxcount" integer,
    "createdat" timestamp,
    "updatedat" timestamp,
    "starttime" timestamp,
    "expiresat" timestamp,
    attributes jsonb
);
CREATE INDEX idx_shortener_attributes ON shortener USING gin (attributes);
alter table shortener owner to shorten;
