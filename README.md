# install dependency

go get github.com/99designs/gqlgen

# bootstrap for create gqlgen.yml , graph/schema.graphqls , server.go

go run github.com/99designs/gqlgen init
go install github.com/99designs/gqlgen@latest

# graphql generate

go run github.com/99designs/gqlgen generate
gqlgen generate

## Architecture Overview

### 🚀 Application Flow

เมื่อมี GraphQL Request เข้ามา ระบบจะทำงานตามลำดับดังนี้:

🌐 HTTP Request (GraphQL Query/Mutation)
↓
🎯 Resolver Layer
↓
🏭 Service Layer (Business Logic)
↓
💾 Repository Layer (Data Access)
↓
🗄️ Database (GORM + PostgreSQL)
↓
📤 Response กลับไป

┌─────────────────┐
│ 🌐 HTTP/GraphQL │ ← Entry point
└─────────────────┘
│
┌─────────────────┐
│ 🎯 Resolver │ ← GraphQL Resolvers (Presentation Layer)
└─────────────────┘
│
┌─────────────────┐
│ 🏭 Service │ ← Business Logic Layer
└─────────────────┘
│
┌─────────────────┐
│ 💾 Repository │ ← Data Access Layer
└─────────────────┘
│
┌─────────────────┐
│ 🗄️ Database │ ← PostgreSQL + GORM
└─────────────────┘

cmd/appv2/
├── main.go # 🚀 Application entry point
internal/appv2/
├── initial/
│ ├── initial.go # 🔧 Dependency injection setup
│ ├── config.go # ⚙️ Database configuration
│ ├── repository.go # 💾 Repository initialization
│ └── service.go # 🏭 Service initialization
├── resolver/
│ ├── resolver.go # 🎯 Resolver struct definition
│ └── schema.resolvers.go # 🎯 GraphQL resolver implementations
├── service/
│ └── user.go # 🏭 Business logic layer
├── repository/
│ └── user.go # 💾 Data access layer
├── entity/
│ └── user.go # 🗃️ Database entities
├── model/
│ └── models_gen.go # 📋 GraphQL models (generated)
└── graph/
├── generated.go # 🔄 Generated GraphQL code
└── schema.graphql # 📝 GraphQL schema definition