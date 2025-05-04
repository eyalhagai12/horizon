FROM golang:1.21 AS build
WORKDIR /app

COPY go.mod go.sum ./
# RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Run stage
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

COPY --from=build /app/server /app/server
WORKDIR /app

EXPOSE 8080
ENTRYPOINT ["/app/server"
