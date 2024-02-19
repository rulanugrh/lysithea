FROM golang:1.21-alpine

ARG EXPOSE_PORT

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy
RUN go build ./cmd/main.go
RUN ./main migrate && ./main seeder
EXPOSE ${EXPOSE_PORT}

CMD [ "./main", "serve" ]