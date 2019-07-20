protoc -I=proto --go_out=plugins=grpc:./grpc ./proto/echo.proto
protoc -I=proto --go_out=plugins=grpc:./grpc ./proto/ceshi.proto
@pause