#!/bin/sh
set -eu

if [ ! -r /run/secrets/postgres_password ]; then
  echo "database password secret is unavailable" >&2
  exit 1
fi
POSTGRES_PASSWORD=$(cat /run/secrets/postgres_password)
export POSTGRES_PASSWORD

exec goose postgres "user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB_NAME} host=${POSTGRES_HOST} port=${POSTGRES_PORT} sslmode=${POSTGRES_SSLMODE}" up
