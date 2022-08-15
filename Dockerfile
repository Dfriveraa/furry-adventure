# FROM golang:1.19-alpine
# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download
# COPY *.go ./

# COPY . .

# RUN go build -o application .
# EXPOSE 5000
# CMD ["/app/application"]

FROM golang:1.16-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

# Hot reloading mod
RUN go get -u github.com/cosmtrek/air && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 8080
EXPOSE 2345

ENTRYPOINT ["air"]

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o my-great-program -a .

### Production
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/my-great-program /usr/local/bin/my-great-program
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/my-great-program"]