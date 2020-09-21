FROM golang:latest AS builder
WORKDIR /go/src/github.com/tonyskapunk/yctf/
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -a -o yctf .

FROM alpine:latest
RUN addgroup -S app \
    && adduser -S -g app app \
    && apk add --no-cache ca-certificates
WORKDIR /app
ADD /templates ./templates
ADD flags.json .
ADD LICENSE .
COPY --from=builder /go/src/github.com/tonyskapunk/yctf/yctf .
RUN chown -R app:app .
USER app
LABEL org.opencontainers.image.source https://github.com/tonyskapunk/yctf
CMD ["./yctf"]
