module github.com/cloudevents/sdk-go/protocol/grpc/v2

go 1.17

require (
	github.com/cloudevents/sdk-go/binding/format/protobuf/v2 v2.12.0
	github.com/cloudevents/sdk-go/v2 v2.12.0
	google.golang.org/grpc v1.51.0
	google.golang.org/protobuf v1.28.1
)

replace (
	github.com/cloudevents/sdk-go/v2 => ../../../v2
	github.com/cloudevents/sdk-go/binding/format/protobuf/v2 => ../../../binding/format/protobuf/v2
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
