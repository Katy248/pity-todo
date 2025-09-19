# Pity todo

Web app runs on port :5900

## Build

```bash
go build -o bin/pity-todo .
```

## Test

```bash
go test -v ./...
```

## Docker

Build docker image

```bash
docker build -t pity-todo .
```

Run docker image

```bash
docker run -p 5900:5900 pity-todo
```
