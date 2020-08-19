# build stage
FROM golang:1.14.5 AS builder

ENV GO111MODULE=on

WORKDIR /mychat

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/mychat.go

# final stage
FROM alpine:3.12.0

WORKDIR /root/

COPY --from=builder /mychat/main .

EXPOSE 9000

CMD ["./main"]


