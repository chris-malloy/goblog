#!/usr/bin/env bash

psql --username "${DB_USER}" <<-EOSQL

-- Users table for profiles.
create table users
(
   id                    serial                                  not null
     constraint users_pkey
     primary key,
  first_name             varchar,
  last_name              varchar,
  email                  varchar default '' :: character varying not null,
  encrypted_password     varchar default '' :: character varying not null,
  sign_in_count          integer default 0                       not null,
  last_sign_in_at        timestamp,
  created_time_stamp     timestamp                               not null,
  updated_time_stamp     timestamp
);

create index index_users_on_id
    on users(id);

create unique index index_users_on_email
  on users(email);

-- Posts Table
create table posts
(
   id                    serial                                  not null
     constraint posts_pkey
     primary key,
  post                   varchar default '' :: character varying not null,
  poster                 varchar,
  created_at             timestamp                               not null,
  updated_at             timestamp
);
EOSQL