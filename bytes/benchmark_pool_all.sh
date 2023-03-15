#!/bin/bash
set -e

cd `dirname $BASH_SOURCE`

go test -run none -bench '^BenchmarkPool' -benchmem  -benchtime=2s