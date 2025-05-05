FROM golang:1.23.4 AS build
WORKDIR /app

COPY . .

RUN ls -a

RUN go mod download -x

RUN go mod tidy 

CMD ["go", "run", "."]
