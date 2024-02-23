[<- README](../../README.md#development-guide)

# Development Guide

## Acceptance criteria for an issue
- Tests are updated to include the described test cases.
- All tests are passing.
- No linter errors occur.
- All changes to the public API are documented in go docs.
- The release notes are updated.

## Commands

Run tests using:

```
go test --cover ./...
```

Run godoc using:

```
godoc -http=:6060
```

Run the linter using:

```
docker run --rm -v ${pwd}:/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v
```

Run gofmt in the project root using:

```
gofmt -s -w .
```