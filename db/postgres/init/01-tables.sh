#!/usr/bin/env bash

psql --username "${DB_USER}" <<-EOSQL

-- Posts Table
create table posts
(
   id                    serial                                  not null
     constraint users_pkey
     primary key,
  post                   varchar default '' :: character varying not null,
  poster                 varchar,
  created_at             timestamp                               not null,
  updated_at             timestamp
);

EOSQL