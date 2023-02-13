# PetPal API

Backend API server for the PetPal project written in Go. Uses Gin framework.

## Running locally

Download a Google Service Account key from the PetPal firebase console, and set the path in your environment

```
export GOOGLE_APPLICATION_CREDENTIALS=relative/path/to/key.json
export DB_HOST=<db-host>
export DB_PORT=<db-port>
export DB_NAME=<db-name>
export DB_USER=<db-username>
export DB_PASS=<db-password>
go run .
```

## Viewing The Docs

SwaggerUI endpoint coming soon

To see the UI for the OApi Documentation

1. Download the [OpenAPI (Swagger) Editor for VSCode](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi)
2. Select the following from the Command Palette
   `OpenAPI: show preview using Swagger UI`
.
...
