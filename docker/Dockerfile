FROM golang:1.19-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY *.go ./

COPY . .

RUN go build -o application .

#DEPLOY
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

# Copy executable
COPY --from=builder /app/application /usr/local/bin/application

EXPOSE 5000
CMD ["/usr/local/bin/application"]