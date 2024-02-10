[<- back](../DOCS.md)

# Development Guide

## Commands

Run the linter using:

```
docker run --rm -v ${pwd}:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v
```

Run gofmt in the project root using:

```
gofmt -s -w .
```