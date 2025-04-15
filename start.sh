#!/bin/sh

set -e

echo "run db migration"
/app/migrate -database "$DB_SOURCE" -path /app/migration -verbose up

echo "start the app"
exec "$@"