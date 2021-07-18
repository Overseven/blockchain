# shellcheck disable=SC2164
cd proto
protoc --go_opt=paths=source_relative --go_out=./ --go-grpc_opt=paths=source_relative --go-grpc_out=./ -I=./ -I=pcommon.proto pcommon.proto node.proto coordinator.proto
echo "Building is finished!"
echo "<Press any key to close>"
read