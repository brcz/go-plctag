#!/bin/bash
echo "=============== Building libplctag ==============="
OLDPWD=$(pwd)
SCRIPTDIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
cd ${SCRIPTDIR}/..
pwd

cd third_party/libplctag
mkdir -p build
cd build
cmake ..
make
cp libplctag.so /artifact/libplctag.so

cd $OLDPWD