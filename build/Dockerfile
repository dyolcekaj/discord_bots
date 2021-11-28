FROM golang:alpine as builder

ARG COMMAND_NAME=spongebob

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o discord-bot ./cmd/${COMMAND_NAME}

FROM scratch

WORKDIR /app

EXPOSE 8080/tcp

COPY --from=builder /app/discord-bot /usr/bin/

ENTRYPOINT ["discord-bot"]