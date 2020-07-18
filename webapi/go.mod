module github.com/i3odja/osbb/webapi

go 1.14

replace github.com/i3odja/osbb/contracts => ./../contracts

require (
	github.com/i3odja/osbb/contracts v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.30.0
)
