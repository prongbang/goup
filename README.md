# Goup

Simple upload file to server with golang

## Install

```
go get github.com/prongbang/goup
```

## How to use

- .env

```
HOST=http://localhost:4000/
```

- main.go

```golang
package main

import (
	"github.com/labstack/echo"
	"github.com/prongbang/goup/pkg/file"
)

func main() {
	e := echo.New()

	e.Static("/public", "public")
	file.ProvideRoute().Initial(e)

	e.Logger.Fatal(e.Start(":4000"))
}
```

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
```

- Response

```json
{
  "filePath": "http://localhost:4000/public/assets/images/BloC.png"
}
```