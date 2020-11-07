#!/usr/bin/env bash

set -e

GOPATH="$(go env GOPATH)"
PEBBLE_VERSION="2.3.0"
PEBBLE_CFGFILE="test/config/pebble-config.json"
PEBBLE_PIDFILE="/tmp/pebble.pid"
PEBBLE_LOGFILE="/tmp/pebble.log"
PEBBLE_CHALLTESTSRV_PIDFILE="/tmp/pebble-challtestsrv.pid"
PEBBLE_CHALLTESTSRV_LOGFILE="/tmp/pebble-challtestsrv.log"
PEBBLE_CHALLTESTSRV_DNS_SERVER="127.0.0.1:5553"
PEBBLE_SRC="git@github.com:letsencrypt/pebble.git"
PEBBLE_DIR="src/github.com/letsencrypt/pebble"
PEBBLE_CA_CERT="test/certs/pebble.minica.pem"

# shellcheck disable=SC2086
cd "${GOPATH}"
rm -rf "${PEBBLE_DIR}"
git clone "${PEBBLE_SRC}" "${PEBBLE_DIR}"
cd "${PEBBLE_DIR}"
git checkout "v${PEBBLE_VERSION}"
go install ./...
pebble-challtestsrv -dns01 "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -http01 "" -tlsalpn01 "" > "${PEBBLE_CHALLTESTSRV_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_CHALLTESTSRV_PIDFILE}"
pebble -dnsserver "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -config "${PEBBLE_CFGFILE}" > "${PEBBLE_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_PIDFILE}"
cat << EOS

pebble (and pebble-challtestsrv) started.

pebble PID:              ${PEBBLE_PIDFILE} (PID $(cat ${PEBBLE_PIDFILE}))
pebble Log:              ${PEBBLE_LOGFILE}
pebble-challtestsrv PID: ${PEBBLE_CHALLTESTSRV_PIDFILE} (PID $(cat ${PEBBLE_CHALLTESTSRV_PIDFILE}))
pebble-challtestsrv Log: ${PEBBLE_CHALLTESTSRV_LOGFILE}
Configured DNS server:   ${PEBBLE_CHALLTESTSRV_DNS_SERVER}
Repository directory:    ${GOPATH}/${PEBBLE_DIR}
Config file:             ${GOPATH}/${PEBBLE_DIR}/${PEBBLE_CFGFILE}
Root CA certificate:     ${GOPATH}/${PEBBLE_DIR}/${PEBBLE_CA_CERT}

EOS
