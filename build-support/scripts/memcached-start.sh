#!/usr/bin/env bash

case "$(uname)" in
  "Darwin")
    brew list memcached > /dev/null 2>&1 || brew install memcached
    brew services run memcached
    ;;

  *)
    # Assuming Ubuntu as that's what our CI runs on. YMMV here, might
    # need to expand this into a separate function for distribution
    # detection if need be.
    sudo apt-get update && sudo apt-get -y install memcached && sudo /etc/init.d/memcached start
    ;;
esac
