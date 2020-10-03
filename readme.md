#gRPC-web without envoy

#Roadmap:

Install dependencies:
- brew install protobuf
- brew install protoc-gen-grpc-web

Get needed packages:
- go get -u google.golang.org/grpc
- go get -u github.com/golang/protobuf/protoc-gen-go

Generate protobuf:
- protoc -I proto proto/*.proto --proto_path=./proto --go_out=plugins=grpc:./backend/proto
- protoc -I proto proto/*.proto --js_out=import_style=commonjs:./frontend/proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/proto


Frontend:
- npm packages needed for frontEnd:
    - "google-protobuf": "^3.11.3",
    - "grpc-web": "^1.0.7"
    
- cd frontend
- npm install
- npm run dev

Backend:
- cd backend
- go run server.go