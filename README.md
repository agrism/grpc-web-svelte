# [gRPC-web](https://github.com/grpc/grpc-web) *without* `envoy`!
The goal of this project is to allow for use of google's gRPC-web npm package without also implicitly requiring that one uses `envoy` alongside it.
by [@agrism](https://github.com/agriam/)!  

<sub><i>and `README.md` updated by [@zod](https://github.com/zudsniper/)</i></sub>  

## PREREQUISITES
_Standalone CLI tools for the most part._
- `golang` 
- `node / npm`

> As for your package manager, obviously this all depends on OS.  
> **for Mac developers**, just use `Homebrew` for everything.  
> **for Linux devs**, you can easily use `apt` & `npm` together to get all necessary CLI tools and packages.  
> **for Windows devs**    
> ...just admit `batch` and `PowerShell` are doo-doo and install WSL2  

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

---

## **2.** Install Requird `Go` Packages
- `$ go get -u google.golang.org/grpc`
- `$ go get -u github.com/golang/protobuf/protoc-gen-go`

---

## **3.** Generate `protobuf`
- `$ protoc -I proto proto/*.proto --proto_path=./proto --go_out=plugins=grpc:./backend/proto`  
- `$ protoc -I proto proto/*.proto --js_out=import_style=commonjs:./frontend/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/proto`  

---

# `DEPLOYMENT`

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

---

## BACKEND
_Backend setup is very easy, but somewhat limited in functionality at the moment._

```sh
$ cd backend
$ go run server.go
```

---
> `README.md` UPDATED BY  
[![spooky.tf](https://user-images.githubusercontent.com/16076573/192673098-48467c36-2d96-43ca-bc02-5ec993989ceb.gif)](https://spooky.tf/)  
**Jason** at *spooky.tf*  
