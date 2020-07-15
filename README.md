# Vince live mix authentication service

This service handle user authentication across the services

For now, it's a poc to play around Golang, modules and try out architecture approach like NestJS or even Symfony.

TODO:
 - Replace all fmt print using a logger
 - Refresh token should be added
 - Testing part must be done asap
 - ... ?

 ## Configuration

 ### Development

 Create a new `.env` file from the `.env_sample` root directory and set your own settings there.

 Default configs are stored in the dedicated file in the `./config`  and it uses annotation to define their values.


## Dependency Injection

To apply good practice, DI is back with Wire (from google). This DI is generated from code (injector) we write in the `wire.go` file.

So everytime you need to inject a service, think about 2 places:
- Change the Provider of the service you want to add a new one
- If the provider is new, add it to the `wire.go` file

## Testing

Testing has not been covered yet cause I was focus others tasks as I was playing around with the language and see good practice around it.
But don't worry, In a pro situation, they will already exists.

 ## Docker

 ### Docker-compose

 A Docker compose file is provided and contains database connection only for now.

#### Postgres

 `docker-compose run --rm postgres` will bring a postgres connection for this project. You can tweak the settings from the `.env`

#### Migration

You can see the Migration part and use the Makefile, but there is 2 commands available:

- To generate a new migration
`docker-compose run --rm migrate create -ext sql -dir /migrations -seq migration_name`

- To apply migration in the database
`docker-compose run --rm migrate -database postgres://database-change-me -path /migrations up`

