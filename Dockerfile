# Build stage
FROM golang:1.23.4 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o server .

# Run stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

COPY --from=build /app/server /server

EXPOSE 8080
ENTRYPOINT ["/server"]
