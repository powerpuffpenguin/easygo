#!/bin/bash
set -e

cd `dirname $BASH_SOURCE`

go test -run none -bench '^(BenchmarkReadString)|(BenchmarkWriteByte)|(BenchmarkWriteRune)|(BenchmarkBufferNotEmptyWriteRead)|(BenchmarkBufferFullSmallReads)|(BenchmarkBufferWriteBlock)$' -benchmem  -benchtime=1s