#!/bin/bash
set -e
cd `dirname $BASH_SOURCE`

Filename="tuple.go"
Count=30

function maketuple
{
    local i
    if [ $2 == 0 ];then
        echo -n "func MakeTuple$1[" >> "$Filename"
    else
        echo -n "func NewTuple$1[" >> "$Filename"
    fi
    for ((i=1;i<=$1;i++));do
        if [[ $i == 1 ]];then
            echo -n "T$i any" >> "$Filename"
        else
            echo -n ", T$i any" >> "$Filename"
        fi
    done
    echo -n "](" >> "$Filename"
    for ((i=1;i<=$1;i++));do
        if [[ $i == 1 ]];then
            echo -n "v$i T$i" >> "$Filename"
        else
            echo -n ", v$i T$i" >> "$Filename"
        fi
    done
    if [ $2 == 0 ];then
        echo -n ") Tuple$1[" >> "$Filename"
    else
        echo -n ") *Tuple$1[" >> "$Filename"
    fi
    for ((i=1;i<=$1;i++));do
        if [[ $i == 1 ]];then
            echo -n "T$i" >> "$Filename"
        else
            echo -n ", T$i" >> "$Filename"
        fi
    done
    echo "] {" >> "$Filename"
    if [ $2 == 0 ];then
        echo -n "	return Tuple$1[" >> "$Filename"
    else
        echo -n "	return &Tuple$1[" >> "$Filename"
    fi
    for ((i=1;i<=$1;i++));do
        if [[ $i == 1 ]];then
            echo -n "T$i" >> "$Filename"
        else
            echo -n ", T$i" >> "$Filename"
        fi
    done
    echo "]{">> "$Filename"
    for ((i=1;i<=$1;i++));do
            echo "		V$i: v$i," >> "$Filename"
    done
    echo "	}" >> "$Filename"
    echo "}" >> "$Filename"
    echo >> "$Filename"
}
function tuple
{
    local count=$1
    if [[ $1 == 1 ]];then
        echo 'package easygo' > "$Filename"
        echo >> "$Filename"
    fi
    local i

    echo -n "type Tuple$1[" >> "$Filename"
    for ((i=1;i<=$1;i++));do
        if [[ $i == 1 ]];then
            echo -n "T$i any" >> "$Filename"
        else
            echo -n ", T$i any" >> "$Filename"
        fi
    done
    echo "] struct {" >> "$Filename"
    for ((i=1;i<=$1;i++));do
        echo "	V$i T$i" >> "$Filename"
    done
    echo "}" >> "$Filename"
    echo >> "$Filename"

    maketuple $1 0
    maketuple $1 1
}

for ((i=1;i<=$Count;i++));do
    tuple "$i"
done