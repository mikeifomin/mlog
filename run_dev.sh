#!/bin/bash

export ADMIN_DIR=`pwd`"/front/dist"
export PG_HOST=localhost
export PG_USER=postgres
export PG_PASS=postgres
export PG_NAME=postgres
export PG_PORT=3576

gin --appPort 8092 --port 3030 -i -x front -x .postgress_data -x vendor --build ./cmd/server/ run
