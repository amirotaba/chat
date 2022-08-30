chat in terminal on golang with mongo and nats.
to run project, enter your database config in utils/utils.go and then
```
cd cmd && go run main.go
```
```
├── cmd
│   └── main.go
├── domain
│   ├── domain.go
│   ├── nats.go
│   └── user.go
├── go.mod
├── go.sum
├── internal
│   ├── nats
│   │   ├── handler
│   │   │   └── natsHandler.go
│   │   ├── PubSub
│   │   │   └── nats.go
│   │   ├── repository
│   │   │   └── mongo
│   │   │       └── db.go
│   │   └── usecase
│   │       └── nats.go
│   └── user
│       ├── handler
│       │   └── userHandler.go
│       ├── repository
│       │   └── mongo
│       │       └── db.go
│       └── usecase
│           └── user.go
└── utils
    └── utils.go
```
