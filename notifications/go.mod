module github.com/i3odja/osbb/notifications

go 1.14

replace github.com/i3odja/osbb/contracts => ./../contracts

require (
	github.com/i3odja/osbb/contracts v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	google.golang.org/grpc v1.30.0
)
