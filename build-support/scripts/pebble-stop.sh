#!/usr/bin/env bash

PEBBLE_PIDFILE="/tmp/pebble.pid"
PEBBLE_EAB_PIDFILE="/tmp/pebble-eab.pid"
PEBBLE_CHALLTESTSRV_PIDFILE="/tmp/pebble-challtestsrv.pid"

PEBBLE_ERROR="false"
PEBBLE_EAB_ERROR="false"
PEBBLE_CHALLTESTSRV_ERROR="false"

PEBBLE_PID=""
if [ -f "${PEBBLE_PIDFILE}" ]; then
  PEBBLE_PID="$(cat "${PEBBLE_PIDFILE}")"
fi

PEBBLE_EAB_PID=""
if [ -f "${PEBBLE_EAB_PIDFILE}" ]; then
  PEBBLE_EAB_PID="$(cat "${PEBBLE_EAB_PIDFILE}")"
fi

PEBBLE_CHALLTESTSRV_PID=""
if [ -f "${PEBBLE_CHALLTESTSRV_PIDFILE}" ]; then
  PEBBLE_CHALLTESTSRV_PID="$(cat "${PEBBLE_CHALLTESTSRV_PIDFILE}")"
fi

if [ -z "${PEBBLE_PID}" ] && [ -z "${PEBBLE_EAB_PID}" ] && [ -z "${PEBBLE_CHALLTESTSRV_PID}" ]; then
  echo "no pebble instances nor pebble-challtestsrv are running; do not need to stop.">&2
  exit 0
fi

# pebble
if [ -n "${PEBBLE_PID}" ]; then
  if [ "$(ps -p "${PEBBLE_PID}" -o comm=)" != "pebble" ]; then
    echo "error: stale PID file ${PEBBLE_PIDFILE}; PID ${PEBBLE_PID} not found or is not \"pebble\".">&2
    PEBBLE_ERROR="true"
  fi

  if [ "${PEBBLE_ERROR}"  != "true" ]; then
    kill "${PEBBLE_PID}" && \
      echo "pebble (PID ${PEBBLE_PID}) stopped." && \
      rm "${PEBBLE_PIDFILE}"
  fi
fi

# pebble (EAB)

if [ -n "${PEBBLE_EAB_PID}" ]; then
  if [ "$(ps -p "${PEBBLE_EAB_PID}" -o comm=)" != "pebble" ]; then
    echo "error: stale PID file ${PEBBLE_EAB_PIDFILE}; PID ${PEBBLE_EAB_PID} not found or is not \"pebble\".">&2
    PEBBLE_EAB_ERROR="true"
  fi

  if [ "${PEBBLE_EAB_ERROR}"  != "true" ]; then
    kill "${PEBBLE_EAB_PID}" && \
      echo "pebble (PID ${PEBBLE_EAB_PID}) stopped." && \
      rm "${PEBBLE_EAB_PIDFILE}"
  fi
fi

# pebble-challtestsrv
if [ -n "${PEBBLE_CHALLTESTSRV_PID}" ]; then
  if [ "$(ps -p "${PEBBLE_CHALLTESTSRV_PID}" -o comm=)" != "pebble-challtestsrv" ]; then
    echo "error: stale PID file ${PEBBLE_CHALLTESTSRV_PIDFILE}; PID ${PEBBLE_CHALLTESTSRV_PID} not found or not \"pebble-challtestsrv\".">&2
    PEBBLE_CHALLTESTSRV_ERROR="true"
  fi

  if [ "${PEBBLE_CHALLTESTSRV_ERROR}"  != "true" ]; then
    kill "${PEBBLE_CHALLTESTSRV_PID}" && \
      echo "pebble-challtestsrv (PID ${PEBBLE_CHALLTESTSRV_PID}) stopped." && \
      rm "${PEBBLE_CHALLTESTSRV_PIDFILE}"
  fi
fi

if [ "${PEBBLE_ERROR}" == "true" ] || [ "${PEBBLE_EAB_ERROR}" == "true" ] || [ "${PEBBLE_CHALLTESTSRV_ERROR}" == "true" ]; then
  exit 1
fi
