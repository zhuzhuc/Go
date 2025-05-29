#!/bin/sh
set -e

echo "===== Debug info start ====="
id
ls -ld /app
ls -la /app/server
echo "Starting Nginx..."
nginx -g "daemon on;"
sleep 2

cd /app
echo "Starting backend service..."
/app/server/server &

wait
