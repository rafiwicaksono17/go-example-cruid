# Solusi E-Commerce API Go - Dokumentasi Lengkap

## Overview
Solusi ini mengimplementasikan REST API e-commerce lengkap dengan Go menggunakan Fiber web framework, GORM sebagai ORM, dan MySQL sebagai database. Solusi menerapkan Clean Architecture dan memenuhi semua 18 persyaratan yang telah ditetapkan.

## Struktur Project

### Entities (Database Models)
- **User** (`internal/pkg/entity/users_entity.go`)
  - ID, Email (unique), Phone (unique), Name, Password, IsAdmin
  - Relations: Store, Addresses, Products, Transactions

- **Store** (`internal/pkg/entity/store_entity.go`)
  - ID, UserID (unique, auto-create saat register), Name, Address, Phone
  - Relations: Products

- **Address** (`internal/pkg/entity/address_entity.go`)
  - ID, UserID, JudulAlamat, PenerimaNama, PenerimaPhone
  - Provinsi, Kabupaten, Kecamatan, Kelurahan (dengan ID dari API wilayah Indonesia)
  - DetailAlamat, IsDefault

- **Category** (`internal/pkg/entity/category_entity.go`)
  - ID, Name (unique)
  - Relations: Products

- **Product** (`internal/pkg/entity/product_entity.go`)
  - ID, StoreID, CategoryID, Name, Description, Price, Quantity, ImageURL
  - Relations: ProductLogs

- **Transaction** (`internal/pkg/entity/transaction_entity.go`)
  - ID, UserID, AddressID, InvoiceNumber (unique), TotalPrice, Status
  - Relations: ProductLogs

- **ProductLog** (`internal/pkg/entity/product_log_entity.go`)
  - ID, TransactionID, ProductID, Quantity, Price
  - Menyimpan snapshot produk pada saat transaksi

### Models (Request/Response)
- `auth_model.go` - RegisterRequest, LoginRequest, AuthResponse
- `user_model.go` - UpdateUserRequest, GetUsersResponse, UserResponse
- `store_model.go` - CreateStoreRequest, UpdateStoreRequest, GetStoresResponse
- `address_model.go` - CreateAddressRequest, UpdateAddressRequest, GetAddressesResponse
- `category_model.go` - CreateCategoryRequest, UpdateCategoryRequest, GetCategoriesResponse
- `product_model.go` - CreateProductRequest, UpdateProductRequest, GetProductsResponse
- `transaction_model.go` - CreateTransactionRequest, UpdateTransactionStatusRequest, GetTransactionsResponse

### Repositories
Implementasi Data Access Layer dengan interface-based pattern:
- `repository.go` - Interface definitions
- `user_repository.go` - User CRUD operations
- `store_repository.go` - Store CRUD operations
- `address_repository.go` - Address CRUD operations
- `category_repository.go` - Category CRUD operations
- `product_repository.go` - Product CRUD operations + filtering
- `transaction_repository.go` - Transaction CRUD operations

**Features:**
- GetByID, GetAll dengan pagination
- Filtering dan sorting (untuk products)
- Query builder dengan GORM

### Usecases (Business Logic)
- `auth_usecase.go` - Register (dengan auto-create store), Login
- `user_usecase.go` - GetMe, UpdateMe, GetAll, Delete
- `store_usecase.go` - CRUD Store dengan authorization checks
- `address_usecase.go` - CRUD Address dengan user isolation
- `category_usecase.go` - CRUD Category (admin only)
- `product_usecase.go` - CRUD Product dengan file upload support
- `transaction_usecase.go` - Create transaction, update status, ProductLog creation

**Features Keamanan:**
- Authorization checks untuk memastikan user hanya akses data mereka sendiri
- Admin-only operations untuk kategori
- Stock management pada product
- Automatic ProductLog creation saat transaksi

### Handlers (HTTP Controllers)
- `auth_handler.go` - Register, Login endpoints
- `user_handler.go` - GetMe, UpdateMe, GetAll, Delete endpoints
- `store_handler.go` - Store CRUD endpoints
- `address_handler.go` - Address CRUD endpoints
- `category_handler.go` - Category CRUD endpoints (admin only)
- `product_handler.go` - Product CRUD endpoints dengan file upload
- `transaction_handler.go` - Transaction endpoints
- `middleware.go` - JWT validation middleware

### Routes
File `internal/server/http/httproute.go` mengdefine semua routes:
- Auth routes (no middleware): /api/v1/auth/register, /api/v1/auth/login
- Protected routes (dengan middleware): /api/v1/users/*, /api/v1/stores/*, etc.

## Persyaratan yang Dipenuhi

### Soal (8 Tugas Utama)
✅ 1. **Service login dan register** - AuthUseCase, AuthHandler
✅ 2. **Toko otomatis terbuat saat register** - Logic di AuthUseCase.Register() dengan auto-create Store
✅ 3. **Service mengelola akun** - UserUseCase, UserHandler dengan GetMe, UpdateMe, Delete
✅ 4. **Service toko** - StoreUseCase, StoreHandler dengan full CRUD
✅ 5. **Service alamat** - AddressUseCase, AddressHandler dengan full CRUD
✅ 6. **Service kategori (admin only)** - CategoryUseCase, CategoryHandler dengan role check
✅ 7. **Service produk** - ProductUseCase, ProductHandler dengan file upload
✅ 8. **Service transaksi** - TransactionUseCase, TransactionHandler dengan ProductLog

### Ketentuan (18 Aturan)
✅ 1. **Routing sesuai collection** - Semua endpoints follow RESTful conventions
✅ 2. **Boleh tambah tidak boleh kurangi API** - Architecture extensible
✅ 3. **Email & phone unique** - Database constraints + validation di UserRepository
✅ 4. **Menggunakan JWT** - JWTtoken generation di utils/jwt.go, middleware auth
✅ 5. **API upload file** - Product handler support multipart form upload (image)
✅ 6. **Toko otomatis saat register** - AutoCreate di AuthUseCase
✅ 7. **Alamat untuk kirim produk** - Address entity dengan fields lengkap
✅ 8. **Hanya admin kelola kategori** - CheckIsAdmin middleware di category handler
✅ 9. **Pagination** - Semua GET endpoints support limit/offset parameters
✅ 10. **Filtering data** - Products support search, category filter, sorting (price, newest)
✅ 11. **User isolation** - Semua operations check user ownership
✅ 12. **User tidak akses alamat user lain** - Authorization check di AddressUseCase
✅ 13. **User tidak kelola toko user lain** - Authorization check di StoreUseCase
✅ 14. **User tidak kelola produk user lain** - Authorization check di ProductUseCase
✅ 15. **User tidak kelola transaksi user lain** - Authorization check di TransactionUseCase
✅ 16. **ProductLog diisi saat transaksi** - Auto-creation di TransactionUseCase
✅ 17. **ProductLog menyimpan snapshot produk** - Fields Quantity, Price disimpan
✅ 18. **Clean Architecture** - Layer separation: Entity → Model → Repository → UseCase → Handler

## Teknologi & Dependencies

```
Go 1.25
Fiber v2.41.0 - Web Framework
GORM v1.31.1 - ORM
MySQL driver - Database
JWT (golang-jwt/jwt) - Authentication
Crypto (golang.org/x/crypto) - Password hashing
Validator (go-playground/validator) - Validation
Viper - Configuration management
Zerolog - Logging
```

## API Endpoints

### Auth (Public)
- `POST /api/v1/auth/register` - Register user (auto-create store)
- `POST /api/v1/auth/login` - Login user (return JWT token)

### Users (Protected)
- `GET /api/v1/users/me` - Get current user profile
- `PATCH /api/v1/users/me` - Update current user
- `DELETE /api/v1/users/me` - Delete current user
- `GET /api/v1/users` - Get all users (admin)
- `GET /api/v1/users/:id` - Get user by ID

### Stores (Protected)
- `POST /api/v1/stores` - Create store
- `GET /api/v1/stores/me` - Get my store
- `PATCH /api/v1/stores/:id` - Update store (owner only)
- `DELETE /api/v1/stores/:id` - Delete store (owner only)
- `GET /api/v1/stores` - Get all stores (with pagination)
- `GET /api/v1/stores/:id` - Get store by ID

### Addresses (Protected)
- `POST /api/v1/addresses` - Create address
- `GET /api/v1/addresses` - Get my addresses (with pagination)
- `GET /api/v1/addresses/:id` - Get my address
- `PATCH /api/v1/addresses/:id` - Update my address
- `DELETE /api/v1/addresses/:id` - Delete my address

### Categories (Protected, Admin Only)
- `POST /api/v1/categories` - Create category (admin)
- `PATCH /api/v1/categories/:id` - Update category (admin)
- `DELETE /api/v1/categories/:id` - Delete category (admin)
- `GET /api/v1/categories` - Get all categories (public)
- `GET /api/v1/categories/:id` - Get category by ID (public)

### Products (Protected)
- `POST /api/v1/products` - Create product (with image upload)
- `GET /api/v1/products/me` - Get my products (with pagination & filtering)
- `GET /api/v1/products` - Get all products (with pagination, filtering, sorting)
- `GET /api/v1/products/:id` - Get product by ID
- `PATCH /api/v1/products/:id` - Update product (owner only)
- `DELETE /api/v1/products/:id` - Delete product (owner only)

**Query Parameters for Products:**
- `limit` (default: 10) - Pagination limit
- `offset` (default: 0) - Pagination offset
- `search` - Search by product name
- `category_id` - Filter by category
- `store_id` - Filter by store
- `sort` - Sort by: newest (default), price_asc, price_desc

### Transactions (Protected)
- `POST /api/v1/transactions` - Create transaction
- `GET /api/v1/transactions/me` - Get my transactions (with status filter)
- `GET /api/v1/transactions/me/:id` - Get my transaction detail
- `PATCH /api/v1/transactions/me/:id` - Update transaction status
- `GET /api/v1/transactions` - Get all transactions (admin, with status filter)

## Key Features

### 1. JWT Authentication
- Token generation saat login/register
- 24-hour expiration
- Payload: user_id, email, is_admin
- Middleware validation di semua protected routes

### 2. User Isolation
- Semua operations mengecek user ownership
- User tidak bisa akses/modifikasi data user lain
- User tidak bisa akses product/store/address/transaksi user lain

### 3. Role-Based Access Control
- Admin-only operations: create/update/delete kategori
- Admin bisa view all transactions

### 4. File Upload
- Support image upload untuk products
- Simpan di `uploads/products/` directory
- Filename: `{userID}_{originalFilename}`

### 5. Product Stock Management
- Automatic stock reduction saat transaksi dibuat
- Validation untuk stock availability
- Rollback jika transaksi gagal

### 6. Transaction with Product Logs
- Auto-creation ProductLog saat transaksi
- Menyimpan snapshot: product quantity, price
- Invoice number auto-generation: `INV-{userID}-{timestamp}`

### 7. Filtering & Pagination
- Semua GET endpoints mendukung limit/offset
- Product support advanced filtering: search, category, store, sorting
- Transaction support status filter

## Cara Menjalankan

### 1. Setup Database
```bash
# Create database
CREATE DATABASE evermos;

# Update .env dengan database credentials
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=evermos
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Run Application
```bash
go run ./app/main.go
# API akan berjalan di http://localhost:PORT (sesuai .env)
```

### 4. Database Migrations
Migrations otomatis berjalan saat aplikasi start via `mysql.RunMigration()`

## Testing dengan Postman

### 1. Register
```
POST /api/v1/auth/register
{
  "email": "user@example.com",
  "phone": "081234567890",
  "name": "User Name",
  "password": "password123"
}
```
Response: JWT Token + User Info + Auto-created Store

### 2. Login
```
POST /api/v1/auth/login
{
  "email": "user@example.com",
  "password": "password123"
}
```

### 3. Create Product (dengan image upload)
```
POST /api/v1/products
Headers: Authorization: Bearer {token}
Body (form-data):
- category_id: 1
- name: "Product Name"
- description: "Description"
- price: 50000
- quantity: 10
- image: [binary file]
```

### 4. Create Transaction
```
POST /api/v1/transactions
Headers: Authorization: Bearer {token}
Body:
{
  "address_id": 1,
  "products": [
    {"product_id": 1, "quantity": 2},
    {"product_id": 2, "quantity": 1}
  ]
}
```

## Error Handling

Semua endpoints return consistent error format:
```json
{
  "code": 400,
  "message": "error message",
  "data": null,
  "error": "detailed error"
}
```

Common errors:
- 400 Bad Request - Validation error, invalid input
- 401 Unauthorized - Missing/invalid JWT token
- 403 Forbidden - Insufficient permissions (e.g., not admin)
- 404 Not Found - Resource tidak ditemukan
- 500 Internal Server Error - Server error

## Clean Architecture Benefits

1. **Separation of Concerns** - Entity, Model, Repository, UseCase, Handler terpisah jelas
2. **Testability** - Interface-based repository memudahkan unit testing
3. **Maintainability** - Mudah menambah fitur baru tanpa mengubah existing code
4. **Scalability** - Architecture memudahkan scaling dan refactoring
5. **Reusability** - UseCase logic bisa digunakan dari berbagai handler

## Future Improvements

1. Add caching layer (Redis) untuk products
2. Add message queue (RabbitMQ) untuk async transaction processing
3. Add email notification saat order created
4. Add payment gateway integration
5. Add order tracking dengan real-time updates (WebSocket)
6. Add admin dashboard untuk manage semua resources
7. Add advanced search dengan Elasticsearch
8. Add rate limiting untuk API endpoints

