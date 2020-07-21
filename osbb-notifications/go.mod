module github.com/i3odja/osbb/notifications

go 1.14

replace github.com/i3odja/osbb/contracts => ./../osbb-contracts

require (
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.2
	github.com/i3odja/osbb/contracts v0.0.0-20200718175326-fcc017c8c09b
	google.golang.org/grpc v1.30.0
)
