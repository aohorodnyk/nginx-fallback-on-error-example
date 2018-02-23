FROM golang:1.10-alpine
WORKDIR /tmp/ws
COPY main.go /opt/ws/main.go
RUN go build -o /tmp/ws/main /opt/ws/main.go \
    && chmod +x ./main

FROM alpine:latest
WORKDIR /opt/ws
COPY --from=0 /tmp/ws/main ./main
CMD ["./main"]
