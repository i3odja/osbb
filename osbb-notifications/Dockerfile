FROM golang:1.14 as modules

WORKDIR /go/src/osbb
COPY . .

ENV GO111MODULE=on

RUN go build -mod=vendor -o ./osbb-notifications

CMD ["/go/src/osbb/osbb-notifications"]
