FROM golang:1.17 AS builder
WORKDIR /build
ADD main.go main.go
ADD go.mod go.mod
ADD cmd cmd
ADD pkg pkg

# Fetch dependencies
RUN go mod download

# Build image as a truly static Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /diffuser -a -tags netgo -ldflags '-s -w' .

FROM alpine:3.15

LABEL org.opencontainers.image.title="Diffuser"
LABEL maintainer="Evance Soumaoro <evanxg852000@gmail.com>"
LABEL org.opencontainers.image.vendor=""
LABEL org.opencontainers.image.licenses="AGPL-3.0"

COPY --from=builder /diffuser /diffuser
ENTRYPOINT ["/diffuser"]
