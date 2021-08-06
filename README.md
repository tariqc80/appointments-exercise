# Simple Appointment Availability API

Example API Server written in golang poroviding endpoints to create, check, and cancel appointments.

Uses:

 - http web framework Gin - https://github.com/gin-gonic/gin
 - database/sql & postgres driver pq

Setup to run locally using Docker
using the images -golang:latest & postgres:13-buster

## Requires
Docker Version 20+


## Endpoints

`POST /schedule/create`
`POST /schedule/available`
`POST /schedule/cancel`

all endpoints take the same two parameters in application/x-www-form-urlencoded

`start` - datetime in RFC3339 format `2021-08-11T10:00:00Z`
`duration` - integer


## Configure and run
Clone and from inside the repo directory run
`docker-compose up -d`

## Create the database
The statements to setup the database are in `sql/up.sql`

Use psql cli from your host machine to create the database
`psql -h 127.0.0.1 -U postgres -f sql/up.sql`
password is `postgrespassword`

Connect to the db container using docker-compose exec and run the commands
`docker-compose exec db bash`
`psql -U postgres`
Then paste the contents of `sql/up.sql` into the psql client.

Or you can copy the contents of `sql/up.sql` into your favorite postgres database client and run the commands to create the database.


