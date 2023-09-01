# go-service-a

Template service A

We will use gin as the web framework.

endpoints exposed:
- GET /api/v1/service-a/
- GET /api/v1/service-a/auth-user-1
- GET /api/v1/service-a/auth-user-2

Follwing the follwing directory structure:

```
project/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   ├── auth_user_1.go
│   │   ├── auth_user_2.go
│   │   └── service_a.go
│   ├── middleware/
│   │   └── middleware.go
│   └── router/
│       └── router.go
├── pkg/
├── vendor/
├── go.mod
└── go.sum
```

