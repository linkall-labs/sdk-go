PWD=$(shell pwd)

generate:
	protoc -I=/Users/wenfeng/personal/workspace/linkall/vanus/proto/include \
		   -I=${PWD} \
		   --go_out=plugins=grpc,paths=source_relative:. cloudevent.proto