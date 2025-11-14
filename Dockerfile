FROM golang:alpine AS builder
WORKDIR /usr/src/app
COPY . .
RUN go build -o server

FROM alpine:latest
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/server ./server
ENTRYPOINT [ "/usr/src/app/server" ]
