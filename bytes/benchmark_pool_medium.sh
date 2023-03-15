#!/bin/bash
set -e

cd `dirname $BASH_SOURCE`

go test -run none -bench '^BenchmarkPoolMedium' -benchmem  -benchtime=1s