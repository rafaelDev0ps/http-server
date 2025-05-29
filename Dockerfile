FROM golang:1.24.3 AS build

WORKDIR /app

COPY go.mod ./

COPY app ./app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -C ./app -o /usr/local/bin/server

# FROM gcr.io/distroless/static

# COPY --from=build --chown=nonroot:nonroot /usr/local/bin/server /usr/local/bin/server

ARG PORT="4221"

ENV HTTP_PORT=${PORT}

EXPOSE ${PORT}

ENTRYPOINT server -port=${HTTP_PORT}
