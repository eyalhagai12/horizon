# syntax=docker/dockerfile:1
FROM golang:1.21 AS build
WORKDIR /app

COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

COPY --from=build /app/app /app/app
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/app"]
