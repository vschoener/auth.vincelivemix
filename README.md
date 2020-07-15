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

## Migration

Migration is not handle by the any app code, but directly from an external library doing all the need for us. This decision saved me time and it was a good call for this project.

### Create new migration

Migration are not created from Entity changes as we could have in NestJs/Symfony/SpringBoot. So you need to write by hand your SQL query.
So generate the file using the make command:

```bash
make new-migration MIGRATION_NAME=name_something
```

Then, write your down and up changes from the 2 new files in `./src/migration` folder you

### Run migration

To run migration, use the following command:

```bash
make run-migration
```

This required a valid database connection using by default the container in the docker-compose file. If you wish to apply this change in a different environment, change the env variables in `DATABASE_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable` or just set a `DATABASE_URL` that should be overridden

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

