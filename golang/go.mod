module github.com/apache/dubbo-samples/golang

require (
	github.com/apache/dubbo-go v1.0.1-0.20200407122856-ed2fa0c9c71f
	github.com/apache/dubbo-go-hessian2 v1.4.0
	github.com/golang/protobuf v1.3.2
	github.com/pkg/errors v0.8.1
	google.golang.org/grpc v1.22.1
)

replace github.com/apache/dubbo-go v1.0.1-0.20200407122856-ed2fa0c9c71f => github.com/hxmhlt/dubbo-go v1.0.1-0.20200407122856-ed2fa0c9c71f

go 1.13
