Go Shorten!
===============

This is a simple URL shortener written in GoLang.

Currently it supports only a CRUD API interface

See more at docs/swagger.json

## Build project

```
make
```

## Test Project

```
make test
```

## Build WebUI

```
make build-webui
```

## Develop Backend

I highly recommend to start the project with air

A config (.air.conf) is located at the root directory.

Start the backend with:

```
./air -c air.conf
``` 

You will find the server at http://localhost:8080 and the Swagger-UI at http://localhost:8080/swagger/index.html

## Develop WebUI

The UI is located in webui/ and based on create-react-app

During the build process it is automatically bundled with the resulting executable.

To start it during development use:

```
npm --prefix webui install
npm --prefix webui start    
```

A browser window should open automatically. Otherwise please visit http://localhost:3000/

A proxy is already configured and points to the backend at  http://localhost:8080

# Feature Roadmap

* [x] Get Description automatically by parsing the given URL (API ready at: /api/url/meta/?url=...) (2021-04-28)
* [ ] Generate QR Code for an URL
    * [x] API (ready at: /api/url/qrcode/?url=...)
    * [ ] UI
* [ ] Add more backends:
    * [x] File storage (2021-05-06)
    * [ ] MySQL
    * [ ] Key-Value-Store
* [ ] Add more authentication methods:
    * [ ] Add more providers (Facebook, Slack, GitHub)
    * [ ] OAuth2 JWT
    * [ ] Username:Password
    * [ ] Special Admin Header

