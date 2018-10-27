build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/teamtools/tp-pos-service proto/pos/pos.proto