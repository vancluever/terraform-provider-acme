#!/usr/bin/env bash

case "$(uname)" in
  "Darwin")
    brew services stop memcached || exit 0
    ;;

  *)
    echo "stopping unsupported on this system, please stop manually."
esac
