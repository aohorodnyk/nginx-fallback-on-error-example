FROM golang:1.10-alpine

WORKDIR /tmp/ws

COPY main.go /opt/ws/main.go

RUN go build -o /tmp/ws/main /opt/ws/main.go

FROM alpine:latest

WORKDIR /opt/ws

COPY --from=0 /tmp/ws/main ./main
RUN chmod +x ./main

CMD ["./main"]
