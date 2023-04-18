# [gRPC-web](https://github.com/grpc/grpc-web) *without* `envoy`!  
This project allows for the use of Google's `gRPC-web` `npm` package without implicitly requiring that one also use `envoy` alongside it.   

[![author  - agrism](https://img.shields.io/badge/author_-agrism-2ea44f?logo=github&logoColor=%23ffffff&style=plastic)](https://github.com/agrism)  

## PREREQUISITES
_Standalone CLI tools for the most part._  
![Golang](https://img.shields.io/badge/GoLang-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![npm - nodejs](https://img.shields.io/badge/Node.js-43853D?style=for-the-badge&logo=node.js&logoColor=white)

> As for your package manager, obviously this all depends on OS.  
> **for Mac developers**, just use `Homebrew` for everything.  
> **for Linux devs**, you can easily use `apt` & `npm` together to get all necessary CLI tools and packages.  
> **for Windows devs**    
> <sup>...just admit `batch` and `PowerShell` are doo-doo and install WSL2</sup>  

# `GETTING STARTED`
This is the barebones installation process which I went through to configure my application environment correctly. 

## **1.** Install Dependencies 
**MacOS** via `Homebrew`
- `$ brew install protobuf`
- `$ brew install protoc-gen-grpc-web`

**Windows** WITH WSL / MSYS2  
  and
**LINUX** *(Tested: Debian 11.3)*  
> _the `-g` tag is used liberally here in order to do our best to guarantee that the package executable will end up referenced within the user $PATH variable._  

- `$ npm install protoc-gen-grpc-web -g` <sub><i>(does this have to be Global or not?)</i></sub>  
- `$ npm install google-protobuf -g`
- `$ npm install grpc-web -g`

## **2.** Install Required `Go` Packages
- `$ go get -u google.golang.org/grpc`
- `$ go get -u github.com/golang/protobuf/protoc-gen-go`

## **3.** Generate `protobuf`
- `$ protoc -I proto proto/*.proto --proto_path=./proto --go_out=plugins=grpc:./backend/proto`  
- `$ protoc -I proto proto/*.proto --js_out=import_style=commonjs:./frontend/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/proto`  

---  

<br>  

# `DEPLOYMENT`
here we discuss how to deploy the webserver [frontend](#frontend) via `npm`, as well as the `Golang` based [backend](#backend). 

## FRONTEND
_required `npm` packages & minimum versions_
(alongside `Sveltekit` obviously)  
  - "google-protobuf": "^3.11.3",
  - "grpc-web": "^1.0.7"
---

### Example Deployment
_example of FRONTEND deployment, with an output folder named as specified and utilized in the previous sections._
```sh
$ cd frontend
$ npm install
$ npm run dev
```

## BACKEND
_Backend setup is very easy, but somewhat limited in functionality at the moment._

```sh
$ cd backend
$ go run server.go
```

---  
[![author  - agrism](https://img.shields.io/badge/author_-agrism-2ea44f?logo=github&logoColor=%23ffffff&style=plastic)](https://github.com/agrism)
[![README  - zudsniper](https://img.shields.io/badge/README.md_-zudsniper-abdab8?logo=github&logoColor=%23ffffff&style=plastic)](https://github.com/zudsniper)
