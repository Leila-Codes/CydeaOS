module cydeaos/api

go 1.19

require (
	cydeaos/libs v0.0.0
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/segmentio/kafka-go v0.4.39
	github.com/sirupsen/logrus v1.9.0
)

replace cydeaos/libs => ./../libs

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
