#cac
Compliance as code 

## Codegen from OpenAPI schema
This section walks through how to codegen types defined in openapi specification for compliance as code server.

1. Install the tool:
Install openapi code generator tooling

```sh
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```

2. Generate server stubs + types:
Generate server stubs and types

```sh
make generate
```
