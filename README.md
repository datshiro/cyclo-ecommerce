# cyclo-ecommerce
Simple ecommerce API server with Golang Echo


# Requirements
- Docker/docker-compose
- Go 1.13++
- Make

#  Run Locally
## Install tools
```
make setup
```

## Start database using Docker
```
make docker/up
```

## init Database
```
make db/up
```

## Seed sample data 
```
make seeder
```


## Start server up
```
make run
```

## To drop  and reset the database
```
make db/down
make db/up
make seeder
```


# Tool I Use
- Sqlboiler: this tool generate models that handle database operations which is much more fast than casual ORM
- Migration: this helps run database migration written in SQL queries

# Project structure

## Design
My code base was design following clean architecture with 2 layers:
- Infras: This layer contains packages involving middlewares, API configurations
- Usecase: This layer is used for business logic handling in which repo packages come along with each target database table

## Directory
- cmd: Implemented necessary executable entry including many server application, seeder runner
- internal: Encapsulated all the source code 
- db: Including SQLBoiler scripts and database migration configures
- scripts: bash scripts to handle operations


# What I have not done
- [ ] Proper API docmentation
- [ ] Cover all the unit tests
