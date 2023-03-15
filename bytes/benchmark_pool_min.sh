#!/bin/bash
set -e

cd `dirname $BASH_SOURCE`

go test -run none -bench '^BenchmarkPoolMin' -benchmem  -benchtime=1s