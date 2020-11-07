#!/usr/bin/env bash

OP="$1"
REC="$2"
DATA="$3"

set -e

case "${OP}" in
  present)
    curl -q -X POST -d "{\"host\":\"${REC}\", \"value\": \"${DATA}\"}" \
      http://localhost:8055/set-txt
    ;;
  cleanup)
    curl -q -X POST -d "{\"host\":\"${REC}\"}" \
      http://localhost:8055/clear-txt
    ;;
  *)
    echo "error: invalid command ${OP}">&2
    exit 1
esac
