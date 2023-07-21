$workspaceRoot = $PSScriptRoot
protoc --proto_path=$workspaceRoot/proto --go_out=service --go-grpc_out=service $workspaceRoot/proto/essence/*.proto
python -m grpc_tools.protoc --proto_path=$workspaceRoot/proto --pyi_out=$workspaceRoot/src --mypy_grpc_out=quiet:$workspaceRoot/src --python_out=$workspaceRoot/src --grpc_python_out=$workspaceRoot/src $workspaceRoot/proto/essence/*.proto
