cluster/run:
	@kubectl apply -f ./ops/kubernetes

cluster/down:
	@kubectl delete -f ./ops/kubernetes

api/build:
	@docker build -t line-bot-demo_api ./services/api/

api/run:
	@docker run -it --rm -p 8080:8080 --name line-bot-demo_api line-bot-demo_api

trigger/protoc:
	@protoc -I./rpc ./rpc/trigger.proto --go_out=plugins=grpc:./services/trigger/

trigger/build:
	@docker build -t line-bot-demo_trigger ./services/trigger/

trigger/run:
	@docker run -it --rm -p 5000:5000 --name line-bot-demo_trigger line-bot-demo_trigger

trigger/down:
	@docker container stop line-bot-demo_trigger

