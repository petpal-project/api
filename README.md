# PetPal API

Backend API server for the PetPal project written in Go. Uses Gin framework.

## Running locally

Local development environment consists of a docker container network with a Postgres DB instance and the API.
Create a .env file from the template:
```
cp .env-template .env
```
And then populate it. To build the images:
```
docker build
```
and to run the network:
```
docker compose up -d
```
The API will be available at `localhost:3000`, and the Postgres DB at `localhost:5432`

When finished:
```
docker compose down
```
## Viewing The Docs

SwaggerUI endpoint coming soon

To see the UI for the OApi Documentation

1. Download the [OpenAPI (Swagger) Editor for VSCode](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi)
2. Select the following from the Command Palette
   `OpenAPI: show preview using Swagger UI`
.
...
