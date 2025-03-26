#!/usr/bin/env bash

set -e

GOPATH="$(go env GOPATH)"
PEBBLE_VERSION="2.8.0-vancluever3"
# config files are relative to script dir
PEBBLE_CFGFILE="../pebblecfg/basic.json"
PEBBLE_PIDFILE="/tmp/pebble.pid"
PEBBLE_LOGFILE="/tmp/pebble.log"
# config files are relative to script dir
PEBBLE_EAB_CFGFILE="../pebblecfg/eab.json"
PEBBLE_EAB_PIDFILE="/tmp/pebble-eab.pid"
PEBBLE_EAB_LOGFILE="/tmp/pebble-eab.log"
# config files are relative to script dir
PEBBLE_PROFILE_CFGFILE="../pebblecfg/alt-profile.json"
PEBBLE_PROFILE_PIDFILE="/tmp/pebble-profile.pid"
PEBBLE_PROFILE_LOGFILE="/tmp/pebble-profile.log"
PEBBLE_CHALLTESTSRV_PIDFILE="/tmp/pebble-challtestsrv.pid"
PEBBLE_CHALLTESTSRV_LOGFILE="/tmp/pebble-challtestsrv.log"
PEBBLE_CHALLTESTSRV_DNS_SERVER="0.0.0.0:5553"
PEBBLE_SRC="https://github.com/vancluever/pebble.git"
PEBBLE_DIR="src/github.com/letsencrypt/pebble"
PEBBLE_CA_CERT="test/certs/pebble.minica.pem"

# Calculate path names
BASIC_CFG="$(realpath "$(dirname "$0")"/${PEBBLE_CFGFILE})"
EAB_CFG="$(realpath "$(dirname "$0")"/${PEBBLE_EAB_CFGFILE})"
PROFILE_CFG="$(realpath "$(dirname "$0")"/${PEBBLE_PROFILE_CFGFILE})"

# Enable alternate roots
export PEBBLE_ALTERNATE_ROOTS="1"

if [ "$1" == "--install" ]; then
  INSTALL="yes"
fi

if [ ! -d "${GOPATH}/${PEBBLE_DIR}" ] || [ "$(cd "${GOPATH}/${PEBBLE_DIR}" && git rev-parse HEAD)" != "$(cd "${GOPATH}/${PEBBLE_DIR}" && git rev-list -n 1 "tags/v${PEBBLE_VERSION}")" ]; then
  echo "pebble source code missing or incorrect version, forcing install."
  INSTALL="yes"
fi

if [ "${INSTALL}" == "yes" ]; then
  cd "${GOPATH}"
  rm -rf "${PEBBLE_DIR}"
  git clone "${PEBBLE_SRC}" "${PEBBLE_DIR}"
  cd "${PEBBLE_DIR}"
  git checkout "v${PEBBLE_VERSION}"
  go install ./...
else
  cd "${GOPATH}/${PEBBLE_DIR}"
  if [ ! -x "${GOPATH}/bin/pebble" ] || [ ! -x "${GOPATH}/bin/pebble-challtestsrv" ]; then
    echo "rebuilding ${GOPATH}/bin/pebble and ${GOPATH}/bin/pebble-challtestsrv from cache."
    go install ./...
  fi
fi

if [ ! -x "${GOPATH}/bin/pebble" ] || [ ! -x "${GOPATH}/bin/pebble-challtestsrv" ]; then
  echo "${GOPATH}/bin/pebble or ${GOPATH}/bin/pebble-challtestsrv missing; error happened in installation.">&2
  exit 1
fi

pebble-challtestsrv -dns01 "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -http01 "" -tlsalpn01 "" > "${PEBBLE_CHALLTESTSRV_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_CHALLTESTSRV_PIDFILE}"
# Basic Pebble instance
pebble -dnsserver "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -config "${BASIC_CFG}" > "${PEBBLE_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_PIDFILE}"
# EAB pebble instance
pebble -dnsserver "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -config "${EAB_CFG}" > "${PEBBLE_EAB_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_EAB_PIDFILE}"
# alt profile pebble instance
pebble -dnsserver "${PEBBLE_CHALLTESTSRV_DNS_SERVER}" -config "${PROFILE_CFG}" > "${PEBBLE_PROFILE_LOGFILE}" 2>&1 &
echo -n $! > "${PEBBLE_PROFILE_PIDFILE}"
cat << EOS

pebble instances (and pebble-challtestsrv) started.

pebble PID:              ${PEBBLE_PIDFILE} (PID $(cat ${PEBBLE_PIDFILE}))
pebble Log:              ${PEBBLE_LOGFILE}
pebble PID (EAB):        ${PEBBLE_EAB_PIDFILE} (PID $(cat ${PEBBLE_EAB_PIDFILE}))
pebble Log (EAB):        ${PEBBLE_EAB_LOGFILE}
pebble PID (Profile):    ${PEBBLE_PROFILE_PIDFILE} (PID $(cat ${PEBBLE_PROFILE_PIDFILE}))
pebble Log (Profile):    ${PEBBLE_PROFILE_LOGFILE}
pebble-challtestsrv PID: ${PEBBLE_CHALLTESTSRV_PIDFILE} (PID $(cat ${PEBBLE_CHALLTESTSRV_PIDFILE}))
pebble-challtestsrv Log: ${PEBBLE_CHALLTESTSRV_LOGFILE}
Configured DNS server:   ${PEBBLE_CHALLTESTSRV_DNS_SERVER}
Repository directory:    ${GOPATH}/${PEBBLE_DIR}
Config file:             ${BASIC_CFG}
Config file (EAB):       ${EAB_CFG}
Config file (Profile):   ${PROFILE_CFG}
Root CA certificate:     ${GOPATH}/${PEBBLE_DIR}/${PEBBLE_CA_CERT}

EOS
