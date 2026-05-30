# Quick Start Guide - E-Commerce API

## Prerequisites
- Go 1.25+
- MySQL 5.7+
- Postman (untuk testing API)

## Setup & Run

### 1. Clone & Navigate
```bash
cd go-example-cruid
```

### 2. Setup Database
```sql
-- Login ke MySQL
mysql -u root -p

-- Create database
CREATE DATABASE evermos;
EXIT;
```

### 3. Configure Environment
Buat/update file `.env`:
```env
# Application
APP_NAME=evermos-api
APP_HOST=localhost
APP_HTTPPORT=8080
APP_VERSION=1.0

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=evermos

# JWT Secret
JWT_SECRET=your_secret_key_here
```

### 4. Install Dependencies
```bash
go mod download
go mod tidy
```

### 5. Run Application
```bash
go run ./app/main.go
```

Server akan start di `http://localhost:8080`

Output yang expected:
```
2024/05/29 10:00:00 API running at localhost:8080
```

## Testing API

### Base URL
```
http://localhost:8080/api/v1
```

### 1. Register User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "phone": "081234567890",
    "name": "John Doe",
    "password": "password123"
  }'
```

**Response**:
```json
{
  "code": 201,
  "message": "register success",
  "data": {
    "token": "eyJhbGc...",
    "user": {
      "id": 1,
      "email": "user@example.com",
      "phone": "081234567890",
      "name": "John Doe",
      "is_admin": false
    }
  }
}
```

### 2. Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

Save token dari response untuk requests berikutnya.

### 3. Get User Profile
```bash
curl -X GET http://localhost:8080/api/v1/users/me \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 4. Get My Store
```bash
curl -X GET http://localhost:8080/api/v1/stores/me \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 5. Create Category (Admin)
Pertama, set user sebagai admin di database:
```sql
UPDATE users SET is_admin = true WHERE id = 1;
```

Kemudian create category:
```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Electronics"
  }'
```

### 6. Create Address
```bash
curl -X POST http://localhost:8080/api/v1/addresses \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "judul_alamat": "Rumah",
    "penerima_nama": "John Doe",
    "penerima_phone": "081234567890",
    "provinsi": "DKI Jakarta",
    "provinsi_id": "31",
    "kabupaten": "Jakarta Pusat",
    "kabupaten_id": "3101",
    "kecamatan": "Gambir",
    "kecamatan_id": "310101",
    "kelurahan": "Gambir",
    "kelurahan_id": "3101011",
    "detail_alamat": "Jl. Sudirman No. 123",
    "is_default": true
  }'
```

### 7. Create Product (with Image Upload)
```bash
curl -X POST http://localhost:8080/api/v1/products \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -F "category_id=1" \
  -F "name=Laptop Gaming" \
  -F "description=High Performance Gaming Laptop" \
  -F "price=15000000" \
  -F "quantity=5" \
  -F "image=@/path/to/image.jpg"
```

### 8. Get All Products (with Filtering)
```bash
# Get all products with pagination
curl -X GET "http://localhost:8080/api/v1/products?limit=10&offset=0"

# Search products
curl -X GET "http://localhost:8080/api/v1/products?search=laptop"

# Filter by category
curl -X GET "http://localhost:8080/api/v1/products?category_id=1"

# Sort by price ascending
curl -X GET "http://localhost:8080/api/v1/products?sort=price_asc"

# Sort by newest
curl -X GET "http://localhost:8080/api/v1/products?sort=newest"
```

### 9. Create Transaction
```bash
curl -X POST http://localhost:8080/api/v1/transactions \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "address_id": 1,
    "products": [
      {"product_id": 1, "quantity": 2},
      {"product_id": 2, "quantity": 1}
    ]
  }'
```

### 10. Get My Transactions
```bash
curl -X GET "http://localhost:8080/api/v1/transactions/me?limit=10&offset=0" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"

# Filter by status
curl -X GET "http://localhost:8080/api/v1/transactions/me?status=pending" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## File Upload

### Save Products Images
Images akan di-save di: `uploads/products/{userID}_{filename}`

Pastikan folder ada:
```bash
mkdir -p uploads/products
```

## Common Issues & Solutions

### Issue: "unauthorized"
**Solution**: Pastikan token di-include di header:
```
Authorization: Bearer {token}
```

### Issue: "invalid database credentials"
**Solution**: Check `.env` file:
- DB_HOST, DB_USER, DB_PASSWORD, DB_NAME harus sesuai
- MySQL server harus running
- Database `evermos` harus sudah di-create

### Issue: "unique constraint failed"
**Solution**: 
- Email/phone sudah terdaftar, gunakan yang baru
- Atau user sudah membuat category dengan nama yang sama

### Issue: "permission denied (admin only)"
**Solution**: Update user menjadi admin di database:
```sql
UPDATE users SET is_admin = true WHERE id = {user_id};
```

## Project Structure
```
go-example-cruid/
├── app/
│   └── main.go                 # Entry point
├── internal/
│   ├── helper/                 # Helper functions
│   ├── infrastructure/
│   │   ├── container/          # Dependency injection
│   │   └── mysql/              # Database setup & migrations
│   ├── pkg/
│   │   ├── entity/             # Database models
│   │   ├── model/              # Request/Response DTOs
│   │   ├── repository/         # Data access layer
│   │   ├── usecase/            # Business logic
│   │   └── controller/         # (legacy)
│   ├── server/
│   │   └── http/
│   │       ├── handler/        # HTTP handlers
│   │       └── httproute.go    # Routes
│   └── utils/                  # Utilities (JWT, password, etc)
├── uploads/                    # File uploads directory
├── .env                        # Environment variables
├── docker-compose.yaml         # Docker compose
├── dockerfile                  # Docker image
├── go.mod                      # Dependencies
├── go.sum                      # Dependencies checksum
└── README.md                   # Documentation
```

## Useful Commands

### Run Application
```bash
go run ./app/main.go
```

### Build Binary
```bash
go build -o evermos-api ./app/main.go
./evermos-api
```

### Run Tests
```bash
go test ./...
```

### Format Code
```bash
go fmt ./...
```

### Lint Code
```bash
golangci-lint run ./...
```

## Database Schema

Aplikasi akan auto-create tables:
- users
- stores  
- addresses
- categories
- products
- transactions
- product_logs

## API Documentation

Lihat file `SOLUTION_DOCUMENTATION.md` untuk dokumentasi lengkap API.
Lihat file `REQUIREMENTS_VERIFICATION.md` untuk verifikasi requirements.

## Support

Untuk pertanyaan atau issues, lihat dokumentasi lengkap di project directory.

