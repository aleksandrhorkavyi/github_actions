FROM golang:1.20-alpine AS builder
# Arguments
ARG APP_NAME=app
# Set env
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
# Prepare directory
RUN mkdir /code
COPY . /code
WORKDIR /code
# Build binary
RUN go mod tidy
RUN go build -mod=readonly -a -o ./bin/app ./cmd/${APP_NAME}

# Finalize
FROM alpine
WORKDIR /go/bin

RUN apk add --no-cache curl

# Application artifacts
COPY --from=builder /code/bin/app .
# System required data
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./app"]
