# Lysithea

<img src="https://cdn.discordapp.com/attachments/1164921098425602108/1208851345080188999/wp-ghost.png?ex=65e4c98c&is=65d2548c&hm=fd0d386e1cad045b3e479b2108fb9f377e7b8b624a67bf327fed44c1b1d43b2a&">

## Getting Started

Haii, is my project, in this project about implementation ELK (Elasticsearch, Logstash, Kibana) stack. Before you running this project you must have :

- [Golang](https://go.dev/dl)
- [PostgreSQL](https://www.postgresql.org/download/)
- [API Client](https://www.postman.com/downloads/)
- [Docker](https://docs.docker.com/engine/install/)

## Tech Stack

- Golang
- PostgreSQL
- GORM (ORM)
- Docker
- Elasticsearch
- Kibana
- Logstash
- JWT (JSON Web Token)
- [Validator](https://github.com/go-playground/validator)
- Mux (Web Framework)
- Opentelemetry
- Elasticsearch APM

## Installation

Clone this project on your local laptop

```
git clone https://github.com/rualnugrh/lysithea
```

Go to project folder

```
cd lysithea
```

Install go modules

```
go mod tidy
```

## Running

Copy `.env.example` to `.env`

```
cp .env.example .env
```

Make sure your PostgresQL is running and you have it setup, and then migrate struct

```
go cmd/main.go migrate
```

After running migration you can seed data to db

```
go cmd/main.go seeder
```

And last, you can running HTTP Server

```
go cmd/main.go serve
```

## Docs

You can see docs in url `APP_URL:APP_PORT/docs/`

## Structure Directory

This project stucture i use, hmm It looks like this structure is familiar to your eyes

```
.
├── Dockerfile
├── LICENSE
├── README.md
├── cmd
│   └── main.go
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── infrastructure
│   ├── elk
│   │   ├── Dockerfile
│   │   ├── logstash.yml
│   │   ├── pipeline.yml
│   │   ├── postgresql-42.7.1.jar
│   │   └── pusher.conf
│   └── nginx
│       └── nginx.conf
└── internal
    ├── config
    │   ├── app.go
    │   ├── db.go
    │   └── elasticsearch.go
    ├── entity
    │   ├── domain
    │   │   ├── category.go
    │   │   ├── order.go
    │   │   ├── product.go
    │   │   └── user.go
    │   └── web
    │       ├── response.go
    │       └── response_data.go
    ├── http
    │   ├── category.go
    │   ├── order.go
    │   ├── product.go
    │   └── user.go
    ├── middleware
    │   ├── jwt.go
    │   ├── token.go
    │   └── validation.go
    ├── repository
    │   ├── category.go
    │   ├── order.go
    │   ├── product.go
    │   └── user.go
    ├── route
    │   ├── category.go
    │   ├── order.go
    │   ├── product.go
    │   └── user.go
    ├── service
    │   ├── category.go
    │   ├── order.go
    │   ├── product.go
    │   └── user.go
    └── util
        ├── migration.go
        ├── pagination.go
        ├── seeder.go
        └── uuid.go
```

## LICENSE

This project is licensed under the terms of the [MIT](./LICENSE) license.
