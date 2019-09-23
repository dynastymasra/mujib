#!/bin/bash -e

case $1 in
  "run")
    docker-compose build mujib
    docker-compose run --rm mujib migrate:run
    docker-compose up mujib
    ;;
  "up")
    docker-compose up mujib
    ;;
  *)
    echo "usage: $0 [run|up]"
    exit 1
    ;;
esac