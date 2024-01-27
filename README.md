# bruh
A CLI utility to easily use the Golang project structure that I prefer.

## Possible output project structure
```
.
├── apps
│   ├── api
│   │   ├── cmd
│   │   │   └── main.go
│   │   ├── go.mod
│   │   └── internal
│   └── user
│       ├── cmd
│       │   └── main.go
│       ├── go.mod
│       └── internal
├── bruh.yaml
├── go.work
└── libs
    ├── config
    │   ├── go.mod
    │   └── main.go
    └── logger
        ├── go.mod
        └── main.go
```