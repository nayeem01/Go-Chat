#!/bin.bash

set -e

echo "Waiting for database to start ..."

psql -v ON_ERROR_STOP=1 --usernaem "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

    DROP DATABASE IF EXISTS chat;
    CREATE DATABASE chat;

EOSQL 