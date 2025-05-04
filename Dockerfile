FROM golang:1.21 AS build
WORKDIR /app

COPY . .

# DEBUG: List contents to confirm files copied
RUN echo "Contents of /app:" && ls -al /app

# DEBUG: Print go.mod to see what's going wrong
RUN echo "go.mod:" && cat /app/go.mod

# Now attempt tidy
RUN go mod tidy
