#!/bin/zsh

if [[ -z $1 ]]; then
  echo "Usage: $0 proto_module"
  echo "Usage: $0 ssh"
  exit 1
fi

proto_module=$1
#handler_dir=$2
cd $(dirname $0)/..
protoc  -I=./proto \
        --go_out=proto \
        --go_opt=paths=source_relative \
        --rpcx_out=proto \
        --rpcx_opt=paths=source_relative \
        ./proto/${proto_module}.proto

