module github.com/i3odja/osbb/webapi

go 1.14

replace github.com/i3odja/osbb/contracts => ./../osbb-contracts

require (
	github.com/i3odja/osbb/contracts v0.0.0-20200718175326-fcc017c8c09b
	google.golang.org/grpc v1.30.0
)
