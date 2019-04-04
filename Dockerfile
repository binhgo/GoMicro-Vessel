# vessel-service/Dockerfile
FROM golang:latest as builder

WORKDIR /go/src/github.com/binhgo/GoMicro-Vessel

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o GoMicro-Vessel -a -installsuffix cgo main.go repository.go handler.go datastore.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/binhgo/GoMicro-Vessel .

CMD ["./GoMicro-Vessel"]