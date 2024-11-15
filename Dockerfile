# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21.5 AS build-stage

WORKDIR /be

RUN mkdir -p -m 0700 ~/.ssh && touch ~/.ssh/known_hosts && ssh-keyscan github.com >> ~/.ssh/known_hosts

# RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

ENV GOPROXY=https://goproxy.io,direct
RUN git config --global url.git@github.com:.insteadOf https://github.com/ && cat ~/.gitconfig && export GOPRIVATE=github.com/tiket/*

COPY go.mod go.sum ./
RUN --mount=type=ssh go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./puxing-be ./cmd/app/main.go

FROM alpine:3.19.1 AS app

WORKDIR /

COPY --from=build-stage ./be/puxing-be ./app
COPY .env /

EXPOSE 8000

ENTRYPOINT ["/app"]
