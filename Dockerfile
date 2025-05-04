FROM golang:1.21 AS build
WORKDIR /app

COPY . .

RUN ls -a

RUN go download

RUN go mod tidy 

CMD ["go", "run", "."]
