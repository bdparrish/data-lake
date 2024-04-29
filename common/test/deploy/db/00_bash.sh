#!/bin/sh

set -e

psql -v --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER webserver WITH PASSWORD 'testpassword';
EOSQL
