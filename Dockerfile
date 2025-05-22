FROM golang:1.24.3 AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY app ./app

RUN CGO_ENABLED=0 GOOS=linux go build -C ./app -o ../server

ARG PORT="4221"

ENV HTTP_PORT=${PORT}

EXPOSE ${PORT}

CMD ./server -port=${HTTP_PORT}
