#!/bin/zsh

if [[ -z $1 ]]; then
  echo "Usage: $0 proto_module"
  echo "Usage: $0 ssh"
  exit 1
fi

proto_module=$1
#handler_dir=$2

#protoc --proto_path=proto \
#       --go_out=server/rpc/${handler_dir}/proto \
#       --go_opt=paths=source_relative \
#       --go_opt=Mssh.proto=server/rpc/${handler_dir}/proto \
#       --go-grpc_out=server/rpc/${handler_dir}/proto \
#       --go-grpc_opt=paths=source_relative \
#       --go-grpc_opt=Mssh.proto=server/rpc/${handler_dir}/proto \
#       proto/${proto_module}.proto

#protoc --proto_path=proto \
#       --go_out=client/rpc/${handler_dir}/proto \
#       --go_opt=paths=source_relative \
#       --go_opt=Mssh.proto=client/rpc/${handler_dir}/proto \
#       --go-grpc_out=client/rpc/${handler_dir}/proto \
#       --go-grpc_opt=paths=source_relative \
#       --go-grpc_opt=Mssh.proto=client/rpc/${handler_dir}/proto \
#       proto/${proto_module}.proto

protoc  -I=./proto \
        --go_out=server/rpc/${proto_module} \
        --go_opt=paths=source_relative \
        --go-grpc_out=.. \
        --go-grpc_opt=Mssh.proto=.. \
        ./proto/${proto_module}.proto
