# Build stage
FROM golang:1.21 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app . && ls -l app

# Run stage
FROM alpine:3.18
RUN apk add --no-cache ca-certificates tzdata && adduser -D -g '' appuser

COPY --from=build /app/app /app/app
RUN chmod +x /app/app
WORKDIR /app
USER appuser

EXPOSE 8080
ENTRYPOINT ["/app/app"]
