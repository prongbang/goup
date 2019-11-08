# Goup

Simple upload file to server with golang

## Run

```
make start
```

## Endpoint

- Request

```
POST http://localhost:4000/api/v1/upload

Fields:
- file=filename.jpg
- platforms=Android
```

- Response

```json
{
  "filePath": "http://localhost:4000/public/assets/images/BloC.png"
}
```