set -e

echo "run db migration"
source /app/app.env
/app/migration/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start app"
exec "$@"
