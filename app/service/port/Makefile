SERVICE=port
PROTO_FILE=pb/$(SERVICE)/pb.proto

gen-proto:
	protoc -I/usr/local/include -I. \
        -I$(GOPATH)/src \
        --proto_path=${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
        --validate_out="lang=go:." \
        --go_out=plugins=grpc:. \
        $(PROTO_FILE)

.PHONY: gen-proto