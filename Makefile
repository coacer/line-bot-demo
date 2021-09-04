protoc:
	@protoc -I./rpc ./rpc/bot-server.proto --go_out=plugins=grpc:./services/bot-server/src/
