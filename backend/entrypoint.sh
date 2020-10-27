#!/bin/sh

echo "dgrapah wt"
if [ "$DATABASE" = "dgraph" ]
then
    echo "Waiting for dgraph..."

    while ! nc -z $DB_HOST 9080; do
      sleep 10
    done

    echo "dgrapah started"
fi
sleep 15
curl -X POST $DB_HOST:$DB_PORT/admin/schema --data-binary '@/usr/src/schema.graphql'

go run server.go

 