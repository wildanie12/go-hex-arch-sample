## Hexagonal Architecture

contoh sampel implementasi backend arsitektur hexagonal menggunakan golang. 



### Struktur Direktori
```
go-hex-arch-sample
├─ .gitignore
├─ README.md
├─ app
│  ├─ [Layer transport disini...]
│  ├─ grpc
│  │  ├─ grpc.go
│  │  ├─ handlers
│  │  │  └─ product.handler.go
│  │  ├─ mappings
│  │  │  └─ product.mapping.go
│  │  ├─ models
│  │  │  └─ product
│  │  │     ├─ product.pb.go
│  │  │     └─ product_grpc.pb.go
│  │  ├─ proto
│  │  │  └─ product.proto
│  │  ├─ provider.go
│  │  └─ route.go
│  └─ http
│     └─ http-restful-disini-nanti
├─ config
│  ├─ config.go
│  ├─ grpc.go
│  └─ http.go
├─ go.mod
├─ go.sum
├─ main.go
├─ models
│  └─ product.go
├─ services
│  └─ [Layer services disini...]
├─ repositories
│  └─ [Layer repository disini...]
└─ utils
   ├─ color
   │  └─ color.go
   └─ json
      └─ json.go
```