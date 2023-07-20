$workspaceRoot = $PSScriptRoot
protoc --proto_path=$workspaceRoot/proto --go_out=service --go-grpc_out=service --python_out=$workspaceRoot/src --python_grpc_out=$workspaceRoot/src --pyi_out=$workspaceRoot/src --mypy_grpc_out=quiet:$workspaceRoot/src $workspaceRoot/proto/essence/*.proto
