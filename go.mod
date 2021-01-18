module github.com/iotaledger/autopeering-sim

go 1.14

require (
	github.com/GeertJohan/go.rice v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.2
	github.com/iotaledger/hive.go v0.0.0-20200424160103-9d9bfc1fe24f
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.6.1
)

replace github.com/iotaledger/hive.go => ../hive.go
