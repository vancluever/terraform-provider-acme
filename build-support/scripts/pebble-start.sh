#!/usr/bin/env bash

set -e

GOPATH="$(go env GOPATH)"
PEBBLE_VERSION="2.3.0"
# config files are relative to script dir
PEBBLE_CFGFILE="../pebblecfg/basic.json"
PEBBLE_PIDFILE="/tmp/pebble.pid"
PEBBLE_LOGFILE="/tmp/pebble.log"
# config files are relative to script dir
PEBBLE_EAB_CFGFILE="../pebblecfg/eab.json"
PEBBLE_EAB_PIDFILE="/tmp/pebble-eab.pid"
PEBBLE_EAB_LOGFILE="/tmp/pebble-eab.log"
PEBBLE_CHALLTESTSRV_PIDFILE="/tmp/pebble-challtestsrv.pid"
PEBBLE_CHALLTESTSRV_LOGFILE="/tmp/pebble-challtestsrv.log"
PEBBLE_CHALLTESTSRV_DNS_SERVER="0.0.0.0:5553"
PEBBLE_SRC="https://github.com/letsencrypt/pebble.git"
PEBBLE_DIR="src/github.com/letsencrypt/pebble"
PEBBLE_CA_CERT="test/certs/pebble.minica.pem"

# Calculate path names
BASIC_CFG="$(realpath "$(dirname "$0")"/${PEBBLE_CFGFILE})"
EAB_CFG="$(realpath "$(dirname "$0")"/${PEBBLE_EAB_CFGFILE})"

if [ "$1" == "--install" ]; then
  cd "${GOPATH}"
  rm -rf "${PEBBLE_DIR}"
  git clone "${PEBBLE_SRC}" "${PEBBLE_DIR}"
  cd "${PEBBLE_DIR}"
  git checkout "v${PEBBLE_VERSION}"
  go install ./...
else
  cd "${GOPATH}/${PEBBLE_DIR}"
fi

pebble-challtestsrv -dns01 "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -http01 "" -tlsalpn01 "" > "${PEBBLE_CHALLTESTSRV_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_CHALLTESTSRV_PIDFILE}"
# Basic Pebble instance
pebble -dnsserver "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -config "${BASIC_CFG}" > "${PEBBLE_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_PIDFILE}"
# EAB pebble instance
pebble -dnsserver "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -config "${EAB_CFG}" > "${PEBBLE_EAB_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_EAB_PIDFILE}"
cat << EOS

pebble instances (and pebble-challtestsrv) started.

pebble PID:              ${PEBBLE_PIDFILE} (PID $(cat ${PEBBLE_PIDFILE}))
pebble Log:              ${PEBBLE_LOGFILE}
pebble PID (EAB):        ${PEBBLE_EAB_PIDFILE} (PID $(cat ${PEBBLE_EAB_PIDFILE}))
pebble Log (EAB):        ${PEBBLE_EAB_LOGFILE}
pebble-challtestsrv PID: ${PEBBLE_CHALLTESTSRV_PIDFILE} (PID $(cat ${PEBBLE_CHALLTESTSRV_PIDFILE}))
pebble-challtestsrv Log: ${PEBBLE_CHALLTESTSRV_LOGFILE}
Configured DNS server:   ${PEBBLE_CHALLTESTSRV_DNS_SERVER}
Repository directory:    ${GOPATH}/${PEBBLE_DIR}
Config file:             ${BASIC_CFG}
Config file (EAB):       ${EAB_CFG}
Root CA certificate:     ${GOPATH}/${PEBBLE_DIR}/${PEBBLE_CA_CERT}

EOS
