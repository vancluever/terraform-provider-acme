#!/usr/bin/env bash

OP="$1"
REC="$2"
DATA="$3"

set -e


case "${OP}" in
  present)
    if [ -e "/tmp/pebble-challtest-dns.lock" ]; then
      echo "error: pebble-challtest-dns already running for $(cat /tmp/pebble-challtest-dns.lock)"
      echo "delete the lockfile if it is stale, or wait for cleanup to run and remove the lockfile."
      exit 1
    fi

    echo -n "pid: $$ domain: ${REC}" > /tmp/pebble-challtest-dns.lock
    trap "rm /tmp/pebble-challtest-dns.lock" ERR

    curl -q -X POST -d "{\"host\":\"${REC}\", \"value\": \"${DATA}\"}" \
      http://localhost:8055/set-txt
    ;;
  cleanup)
    curl -q -X POST -d "{\"host\":\"${REC}\"}" \
      http://localhost:8055/clear-txt
    rm /tmp/pebble-challtest-dns.lock
    ;;
  *)
    echo "error: invalid command ${OP}">&2
    exit 1
esac
