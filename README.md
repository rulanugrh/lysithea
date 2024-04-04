# Lysithea

<img src="https://raw.githubusercontent.com/D3Ext/aesthetic-wallpapers/main/images/purple-girl.png">

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

## Before Running
You must running db, elasticsearch, kibana, and logstash.

```
docker compose up -d db
docker compose up -d elasticsearch
docker compose up -d kibana
docker compose up -d logstash
```

If you dont have docker you can install docker, see [Docker Docs](https://docs-docker-com.translate.goog/engine/install).

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
├── internal
│   ├── config
│   │   ├── app.go
│   │   ├── db.go
│   │   ├── elasticsearch.go
│   │   └── otelapm.go
│   ├── entity
│   │   ├── domain
│   │   │   ├── category.go
│   │   │   ├── order.go
│   │   │   ├── product.go
│   │   │   └── user.go
│   │   └── web
│   │       ├── response.go
│   │       └── response_data.go
│   ├── http
│   │   ├── category.go
│   │   ├── order.go
│   │   ├── product.go
│   │   └── user.go
│   ├── middleware
│   │   ├── cors.go
│   │   ├── jwt.go
│   │   ├── token.go
│   │   └── validation.go
│   ├── mocks
│   │   ├── category_repository_mock.go
│   │   └── user_repository_mock.go
│   ├── repository
│   │   ├── category.go
│   │   ├── order.go
│   │   ├── product.go
│   │   └── user.go
│   ├── route
│   │   ├── category.go
│   │   ├── order.go
│   │   ├── product.go
│   │   └── user.go
│   ├── service
│   │   ├── category.go
│   │   ├── order.go
│   │   ├── product.go
│   │   └── user.go
│   └── util
│       ├── migration.go
│       ├── pagination.go
│       ├── seeder.go
│       └── uuid.go
└── tests
    ├── category_repository_test.go
    └── user_repository_test.go
```

## LICENSE

This project is licensed under the terms of the [MIT](./LICENSE) license.
