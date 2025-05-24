# clothes_order
uniformes-app/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   ├── uniform.go
│   │   ├── upload.go
│   │   └── quote.go
│   ├── models/
│   │   ├── uniform.go
│   │   ├── order.go
│   │   └── pricing.go
│   ├── services/
│   │   ├── uniform_service.go
│   │   ├── pricing_service.go
│   │   └── image_service.go
│   └── database/
│       ├── migrations/
│       └── connection.go
├── web/
│   ├── static/
│   │   ├── css/
│   │   ├── js/
│   │   └── images/
│   │       ├── linhas/
│   │       ├── pecas/
│   │       ├── modelos/
│   │       └── tecidos/
│   └── templates/
│       ├── base.html
│       ├── configurador.html
│       └── components/
├── uploads/
└── go.mod
