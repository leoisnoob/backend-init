grpc:
	protoc -I./api/ \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
	--go_out=paths=source_relative,plugins=grpc:./api/ \
	--govalidators_out=paths=source_relative:./api/ service.proto mod.proto;

	protoc -I./api/ \
	-I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
	--grpc-gateway_out=./api/ \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=false \
	service.proto

	# protoc -I./api \
	# -I ${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/ \
	# --swagger_out=logtostderr=true:. service.proto
	# sed -i -e "s/HITLBack_//g" ./service.swagger.json
	# rm service.swagger.json-e
	# docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli generate \
	# -i /local/service.swagger.json \
	# -l html2 \
	# -o /local/public/

serve:
	go run main.go serve
migrate:
	go run main.go migrate
name:='registry.sensetime.com/ucg_test/example-backend:v0.0.6'
docker:
	docker build -t ${name} .
	docker push ${name}
	ssh root@viper "cd /etc/kubernetes/apps/hitl-console/console; ./upgrade.sh back"
