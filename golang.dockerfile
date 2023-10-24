FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/rwas2505/go-chi/
WORKDIR /go/src/github.com/rwas2505/go-chi
RUN go mod download
COPY . /go/src/github.com/rwas2505/go-chi
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/ryanGoChi github.com/rwas2505/go-chi

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/rwas2505/go-chi/build/ryanGoChi /usr/bin/ryanGoChi
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/ryanGoChi"]