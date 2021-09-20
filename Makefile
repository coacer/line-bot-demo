cluster/run:
	@kubectl apply -f ./ops/kubernetes

cluster/down:
	@kubectl delete -f ./ops/kubernetes

local/up:
	@docker-compose up -d

local/down:
	@docker-compose down

local/db-connect:
	@docker-compose exec spanner-cli spanner-cli -p local -i local-instance -d line-bot

# local/yo:
# 	@SPANNER_EMULATOR_HOST=localhost:9010 yo local local-instance line-bot -o ./services/webhook/models -p models && SPANNER_EMULATOR_HOST=localhost:9010 yo local local-instance line-bot -o ./services/channel/models -p models
#
api/build:
	@docker build -t line-bot-demo_api ./services/api/

webhook/protoc:
	@protoc -I./rpc ./rpc/line_webhook.proto --go_out=plugins=grpc:./services/webhook/infrastructure/grpc/

webhook/build:
	@docker build -t line-bot-demo_webhook ./services/webhook/

channel/protoc:
	@protoc -I./rpc ./rpc/channel.proto --go_out=plugins=grpc:./services/channel/

