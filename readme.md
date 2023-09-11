# Ports & Adapters Architecture

![Application](https://github.com/Kartochnik010/hex-arch/assets/78559811/95c5abe2-24f7-454e-96b6-3b0e47124323)
### Primary driving adapters - HTTP, gRPC, cli
### Secondary driven adapters - DB, redis
### The goal is to decouple so the adapters can be switched easily. Let’s say change HTTP to CLI or MySQL to Postgres.

| Layer | Description |
| --- | --- |
| Domain | core business logic. |
| Application | will orchestrate the use of the domain code as well as Adapter requests from framework layer to domain layer |
| Framework | logic for outside components(adapters). |

```
.
├── Dockerfile
├── cmd
│   └── main.go
├── docker-compose.yaml
├── go.mod
├── go.sum
├── grpc_entrypoint.sh
├── internal
│   ├── adapters
│   │   ├── app
│   │   │   └── api
│   │   │       └── api.go
│   │   ├── core
│   │   │   ├── arithmetic
│   │   │   ├── arithmetic.go
│   │   │   └── arithmetic_test.go
│   │   ├── framework
│   │   │   ├── left
│   │   │   │   └── grpc
│   │   │   │       ├── pb
│   │   │   │       │   ├── arithmetic_svc_grpc.pb.go
│   │   │   │       │   └── number_msg.pb.go
│   │   │   │       ├── proto
│   │   │   │       │   ├── arithmetic_svc.proto
│   │   │   │       │   └── number_msg.proto
│   │   │   │       ├── rpc.go
│   │   │   │       ├── rpc_test.go
│   │   │   │       └── server.go
│   │   │   └── right
│   │   │       └── db
│   │   │           └── db.go
│   │   └── right
│   └── ports
│       ├── app.go
│       ├── core.go
│       ├── framework_left.go
│       └── framework_right.go
├── sqlite.db
└── testdb
    └── init.sql
```

