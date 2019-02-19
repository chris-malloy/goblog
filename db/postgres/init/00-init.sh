#!/usr/bin/env bash

# Don't screw with any of this.
psql --username "${POSTGRES_USER}" <<-EOSQL
  create role ${DB_USER} with password '${DB_PASS}' nosuperuser nocreatedb nocreaterole login;
  create database ${DB_NAME} with owner ${DB_USER};
  revoke all on database ${DB_NAME} from public;
  grant usage on schema public to ${DB_USER};
  grant connect on database ${DB_NAME} to ${DB_USER};
  grant all on database ${DB_NAME} to ${DB_USER};
EOSQL