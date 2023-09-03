#!/bin/sh

set -e
COMMAND=$@

echo 'Waiting for db to start...'
maxTries=10

while [ "$maxTries" -gt 0 ] && ! mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT"; do
    maxTries=$(($maxTries - 1))
    sleep 3
done
echo 
if [ "$maxTries" -le 0]; then
    echo >&2 'error: unable to contact mysql after 10 attempts'
    exit 1
fi

exec $COMMAND