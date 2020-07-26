module github.com/i3odja/osbb/webapi

go 1.14

replace github.com/i3odja/osbb/contracts => ./../osbb-contracts

replace github.com/i3odja/osbb/shared => ./../osbb-shared

require (
	github.com/i3odja/osbb/contracts v0.0.0-20200718175326-fcc017c8c09b
	github.com/i3odja/osbb/shared v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.6.0
	google.golang.org/grpc v1.30.0
)
