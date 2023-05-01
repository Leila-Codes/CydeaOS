module cydeaos/game-manager

go 1.19

require (
	cydeaos/libs v0.0.0
	github.com/segmentio/kafka-go v0.4.39
	github.com/sirupsen/logrus v1.9.0
	github.com/tjarratt/babble v0.0.0-20210505082055-cbca2a4833c1
)

replace cydeaos/libs v0.0.0 => ./../libs

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.27.6 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
)
