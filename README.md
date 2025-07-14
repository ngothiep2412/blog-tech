# Blog Tech Platform

A simple tech blog platform built with Go, gRPC and MySQL for learning backend development.

## Features

- User registration and login
- Create and publish tech articles
- Like and share articles
- Categories and tags
- REST API for clients
- gRPC for internal services

## Tech Stack

- **Backend**: Go
- **Database**: MySQL
- **API**: REST + gRPC
- **ORM**: GORM
- **Auth**: JWT

## Quick Start

### 1. Setup Database
```bash
# Run MySQL with Docker
docker run --name mysql-blog \
  -e MYSQL_ROOT_PASSWORD=YOUR_PASSWORD \
  -e MYSQL_DATABASE=blog_tech \
  -p 3307:3306 \
  -d mysql:8.0
```

### 2. Clone and Install
```bash
git clone <your-repo>
cd blog-tech
go mod download
```

### 3. Generate Proto Files
```bash
# Install buf
go install github.com/bufbuild/buf/cmd/buf@latest

# Generate protobuf
buf generate
```

### 4. Run Server
```bash
go run main.go
```

## API Endpoints

### REST API (Port 8080)
```http
POST /api/v1/auth/register    # User registration
POST /api/v1/auth/login       # User login
GET  /api/v1/users/profile    # Get profile
PUT  /api/v1/users/profile    # Update profile
```

### gRPC (Port 50051)
```protobuf
# Internal service communication
rpc GetUserById()      # Get user info
rpc GetUserProfile()   # Get detailed profile
rpc UpdateUserStats()  # Update user statistics
```

## Project Structure

```
blog-tech/
├── internal/
│   └── users/
│       ├── biz/           # Business logic
│       ├── model/         # Data models
│       ├── repository/    # Database layer
│       ├── transport/
│       │   ├── api/       # REST handlers
│       │   └── rpc/       # gRPC handlers
│       └── proto/         # Protobuf files
├── common/                # Shared utilities
├── main.go
└── README.md
```

## Configuration

Create `.env` file:
```bash
go build -o blog-tech
./blog-tech outenv > .env
```

## Testing

```bash
# Test registration
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","email":"test@example.com","password":"123456","full_name":"Test User"}'

# Test login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"123456"}'
```

## License

MIT License