FROM golang:1.21-alpine

ARG EXPOSE_PORT

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy
RUN go build -o cmd/main

EXPOSE ${EXPOSE_PORT}
RUN ./cmd/main migrate && ./cmd/main seeder

CMD ["./cmd/main", "serve"]