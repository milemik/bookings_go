POSTGRES_NAME=bookings
POSTGRES_DB_NAME=bookings
POSTGRES_PASSWORD=pass123
POSTGRES_USER=mile

CHECK_POSTGRES_CREATED="$(docker images|grep $POSTGRES_NAME)"

echo "Running/Creating postgresDB for developing"
docker run --rm --name $POSTGRES_NAME -v "$PWD/postgresdb_volume":/var/lib/postgresql/data -e POSTGRES_PASSWORD=$POSTGRES_PASSWORD -e POSTGRES_USER=$POSTGRES_USER -e POSTGRES_DB=$POSTGRES_DB_NAME -p 5432:5432 postgres