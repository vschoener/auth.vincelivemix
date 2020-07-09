# Vince live mix authentication service


This service handle user authentication across the services

For now, it's a poc to play around Golang, modules and try out architecture approach like NestJS or even Symfony.

TODO:
 - [Later] Check if dependency injection is the way to go with Golang or old fashion way should be used
 - Replace all fmt print using a logger
 - Create user database 
 - Log in user using database
 - Refresh token should be added
 - Testing part must be done asap
 - ... ?

 ## Configuration

 ### Development

 Create a new `.env` file from the `.env_sample` root directory and set your own settings there.
 
 Default configs are stored in the dedicated file in the `./config`  and it uses annotation to define their values.


 ## Docker

 ### Docker-compose

 A Docker compose file is provided and contains database connection only for now.

 `docker-compose up` will bring a postgres connection for this project. You can tweak the settings from the `.env`