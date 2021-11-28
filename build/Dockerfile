FROM golang:alpine as builder

ARG COMMAND_NAME=spongebob

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./cmd/${COMMAND_NAME}

FROM scratch

WORKDIR /app

COPY --from=builder /app/${COMMAND_NAME} /usr/bin/

ENTRYPOINT ["${COMMAND_NAME}"]