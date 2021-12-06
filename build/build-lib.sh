#!/bin/bash

PLCTAG_VERSION=2.4.5

echo "=============== Building libplctag ==============="
pushd $(pwd)
mkdir /tmp/libplctag
cd /tmp/libplctag
wget https://github.com/kyle-github/libplctag/archive/v$PLCTAG_VERSION.tar.gz
tar xzf v$PLCTAG_VERSION.tar.gz

cd libplctag-$PLCTAG_VERSION && mkdir -p build && cd build && cmake .. && make && sudo make install
rm -fr /tmp/libplctag

cp /usr/local/lib/* /artifact/lib/

echo "=============== listing libs ==============="
ls -la /usr/local/lib
echo "=============== listing artifacts ==============="
ls -la /artifact

popd
