module github.com/i3odja/osbb/notifications

go 1.14

replace github.com/i3odja/osbb/contracts => ./../osbb-contracts

replace github.com/i3odja/osbb/shared => ./../osbb-shared

require (
	github.com/gorilla/mux v1.7.4
	github.com/gorilla/websocket v1.4.2
	github.com/i3odja/osbb/contracts v0.0.0-20200718175326-fcc017c8c09b
	github.com/i3odja/osbb/shared v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
	google.golang.org/grpc v1.30.0
)
