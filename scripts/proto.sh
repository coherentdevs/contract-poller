protoc -I=. --go_opt=paths=source_relative --go_out=protos/go protos/shared/*.proto
protoc -I=. --go_opt=paths=source_relative --go_out=protos/go protos/evm/node_client/*.proto
