# blockless powered vite app

dApps on steroids.

nnApplications should be easy to build. Users should be able to power a network by using a dApp. Operators should be able to launch AVS' with ease. 

This repo is a radical change on building dApplications with Network Neutral Principals. 

What the f' is this?

* Vite Powered Front End using Blockless Web Workers
* Golang Powered Rest Server with a P2P Server Provider Network built in.

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