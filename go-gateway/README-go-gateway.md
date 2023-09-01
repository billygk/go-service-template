Template for gateway service

We will use gin as the web framework.

Follwing the follwing directory structure:

```
project/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handlers/
│   │   ├── healthcheck.go
│   │   └── oauth2.go
│   ├── middleware/
│   │   ├── circuitbreaker.go
│   │   ├── logging.go
│   │   ├── metrics.go
│   │   ├── ratelimit.go
│   │   ├── retry.go
│   │   └── tracing.go
│   └── router/
│       └── router.go
├── pkg/
└── vendor/
```

Gateway should have the following features:
- ratelimit
- circuit breaker
- retry
- logging
- metrics
- tracing
- health check
- oauth2 using keycloak


