# blockless powered vite app

* Vite Powered Front End 
* Golang Powered Rest Server

# dev

To build the basics : 

* go 1.21
* node 1.18

To build more : 

* Swift for OSX 14+

```bash
make dev
```

Starts a development server.

* Vite Server http://localhost:5173/ (messages swallowed for cleaner ux)
* Development Server http://localhost:8080


* Install protoc-js 
```
npm install -g protoc-gen-js
```

# production

`make build` will produce a single executible containing the `vite` app, with the golang backed api server.

```bash
./myapp
```

## build

```
Usage of ./myapp:
  -headless
        Run in headless mode without opening the browser
  -port int
        Provide a listenable port e.g. 8080
```
## try it

two ways to join the reference network!

1. Grab a binary for your system here

https://github.com/dmikey/go-vite-app/releases/tag/latest

2. Give a docker a spin! 

```bash
docker run -p 8080:8080 ghcr.io/dmikey/go-vite-app:v0.0.4
```