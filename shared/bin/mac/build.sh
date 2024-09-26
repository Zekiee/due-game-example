#!/bin/bash

PROTOPATH=$1/shared/pb


#with rpcx
echo "proto building"
for dir in $PROTOPATH/pb_*; do
  (
  baseDir=`basename $dir`
# by default, protoc-gen-go
  protoc --proto_path="$PROTOPATH/" --go_out $PROTOPATH/../../../ --rpcx_out=$PROTOPATH/../../../ --js_out=import_style=commonjs,binary:. $PROTOPATH/$baseDir/*.proto
  )
done