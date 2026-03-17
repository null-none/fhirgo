# FHIR GO - Golang Implementation of FHIR Resources

This package represents the FHIR data structure implemented in pure golang.

## Supported versions

| Version | Package |
|---------|---------|
| STU3    | `github.com/null-none/fhirgo/STU3` |
| R4      | `github.com/null-none/fhirgo/R4` |
| R5      | `github.com/null-none/fhirgo/R5` |
| R6      | `github.com/null-none/fhirgo/R6` |

## Running tests

Run tests for all versions:

```bash
go test ./...
```

Run tests for a specific version:

```bash
go test ./R4/...
go test ./R5/...
go test ./R6/...
```

Run with verbose output:

```bash
go test ./R4/... -v
```

Run a specific test:

```bash
go test ./R4/resources/... -run TestPatientMarshal
```

## Docker

```bash
docker system prune -a
docker volume prune
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
docker rmi $(docker images -a -q)
```

```bash
docker-compose up
```