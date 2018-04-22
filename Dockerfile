# Start by building the application.
FROM golang:1 as builder

RUN apt-get update && apt-get install -y --no-install-recommends upx

WORKDIR /go/src/gitlab.com/Simple-Bank

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep && dep ensure -v -vendor-only

RUN  CGO_ENABLE=0 GOOS=linux go build \
 -tags netgo \
 -installsuffix netgo,cgo \
 -v -a \
 -ldflags '-s -w -extldflags "-static"' \
 -o app

FROM gcr.io/distroless/base
WORKDIR /go/src/gitlab.com/Simple-Bank
COPY --from=builder /go/src/gitlab.com/Simple-Bank .
ENTRYPOINT ["./app"]
EXPOSE 3000