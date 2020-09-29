module DebugService

go 1.12

require (
	github.com/Shopify/sarama v1.26.4
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1 // indirect
	github.com/mapgoo-lab/atreus v0.1.26
	github.com/prometheus/common v0.10.0
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	google.golang.org/genproto v0.0.0-20200620020550-bd6e04640131
	google.golang.org/grpc v1.29.1
	gopkg.in/jcmturner/gokrb5.v7 v7.5.0
)

replace google.golang.org/grpc v1.29.1 => google.golang.org/grpc v1.26.0
