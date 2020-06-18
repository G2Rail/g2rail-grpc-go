gen:
	protoc lib/proto/*.proto --go_out=plugins=grpc:g2rail --grpc-gateway_out=:g2rail

client:
	go run cmd/client/main.go -address ${ADDRESS} -apikey ${APIKEY} -secret ${SECRET}

.PHONY: gen client