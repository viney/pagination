#!/bin/sh

PWDDIR="$(dirname "`readlink -f $0`")"
PAGING="$(dirname "$PWDDIR")"

export GOPATH=$PAGING

echo "exec sql start"
psql test -h 127.0.0.1 -p 4932 -U viney -f $PWDDIR/test.sql
echo "exec sql start"

cd $PWDDIR/service
go clean
go build
echo "go clean and build complete"
./service
