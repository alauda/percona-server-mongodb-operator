#!/bin/sh

set -o errexit
set -o xtrace

# show id info in container log
echo "container user: $(id)"

install -o "$(id -u)" -g "$(id -g)" -m 0755 -D /ps-entry.sh /data/db/ps-entry.sh
install -o "$(id -u)" -g "$(id -g)" -m 0755 -D /mongodb-healthcheck /data/db/mongodb-healthcheck
chown 1001:0 /data/db
