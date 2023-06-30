protoc -I ./proto --go_out=./service --go-grpc_out=./service ./proto/*.proto

python -m grpc_tools.protoc -I./proto --pyi_out=./src/proto/typings --python_out=./src/proto --grpc_python_out=./src/proto ./proto/*.proto
