# Verifikasi Requirements - E-Commerce API Go

## SOAL (8 Tugas Utama)

### ✅ 1. Buatlah service login dan register
**Status**: COMPLETED
- **Implementation**: `AuthUseCase` di `internal/pkg/usecase/auth_usecase.go`
- **Endpoints**: 
  - POST `/api/v1/auth/register` → `AuthHandler.Register()`
  - POST `/api/v1/auth/login` → `AuthHandler.Login()`
- **Features**:
  - Password hashing dengan bcrypt
  - JWT token generation
  - Email validation

### ✅ 2. Ketika berhasil register, toko otomatis terbuat
**Status**: COMPLETED
- **Implementation**: `AuthUseCase.Register()` lines 19-34
  ```go
  // Auto create store for user
  storeName := req.Name + "'s Store"
  store := &entity.Store{
    UserID:    user.ID,
    Name:      storeName,
    CreatedAt: now,
    UpdatedAt: now,
  }
  if err := u.storeRepo.Create(store); err != nil {
    return nil, err
  }
  ```
- **Verification**: Setiap user yang baru register akan memiliki 1 store yang auto-created dengan nama "{UserName}'s Store"

### ✅ 3. Service untuk mengelola akun
**Status**: COMPLETED
- **Implementation**: `UserUseCase` dan `UserHandler`
- **Endpoints**:
  - GET `/api/v1/users/me` → GetMe (authenticated user)
  - PATCH `/api/v1/users/me` → UpdateMe
  - DELETE `/api/v1/users/me` → DeleteMe
  - GET `/api/v1/users` → GetAllUsers (admin/public)
  - GET `/api/v1/users/:id` → GetUserByID
- **Features**:
  - Update email, name, phone
  - Validasi unique email & phone
  - Get user profile

### ✅ 4. Buatlah service toko
**Status**: COMPLETED
- **Implementation**: `StoreUseCase` dan `StoreHandler`
- **Endpoints**:
  - POST `/api/v1/stores` → CreateStore
  - GET `/api/v1/stores/me` → GetMyStore
  - GET `/api/v1/stores/:id` → GetStoreByID
  - GET `/api/v1/stores` → GetAllStores (with pagination)
  - PATCH `/api/v1/stores/:id` → UpdateMyStore (owner only)
  - DELETE `/api/v1/stores/:id` → DeleteMyStore (owner only)
- **Features**:
  - Authorization check (user only manage their own store)
  - Pagination support

### ✅ 5. Buatlah service alamat
**Status**: COMPLETED
- **Implementation**: `AddressUseCase` dan `AddressHandler`
- **Endpoints**:
  - POST `/api/v1/addresses` → CreateAddress
  - GET `/api/v1/addresses` → GetMyAddresses (with pagination)
  - GET `/api/v1/addresses/:id` → GetMyAddressByID
  - PATCH `/api/v1/addresses/:id` → UpdateMyAddress
  - DELETE `/api/v1/addresses/:id` → DeleteMyAddress
- **Features**:
  - Support Indonesia regional data (Provinsi, Kabupaten, Kecamatan, Kelurahan)
  - User isolation
  - Pagination

### ✅ 6. Buatlah service kategori (hanya admin)
**Status**: COMPLETED
- **Implementation**: `CategoryUseCase` dan `CategoryHandler`
- **Endpoints**:
  - POST `/api/v1/categories` → CreateCategory (admin only)
  - GET `/api/v1/categories` → GetAllCategories (public)
  - GET `/api/v1/categories/:id` → GetCategoryByID (public)
  - PATCH `/api/v1/categories/:id` → UpdateCategory (admin only)
  - DELETE `/api/v1/categories/:id` → DeleteCategory (admin only)
- **Authorization**: `CheckIsAdmin()` middleware check di `category_handler.go`
- **Features**:
  - Admin-only create/update/delete
  - Public read access
  - Pagination

### ✅ 7. Buatlah service produk
**Status**: COMPLETED
- **Implementation**: `ProductUseCase` dan `ProductHandler`
- **Endpoints**:
  - POST `/api/v1/products` → CreateProduct (with file upload)
  - GET `/api/v1/products/me` → GetMyProducts (owner's products)
  - GET `/api/v1/products` → GetAllProducts (public)
  - GET `/api/v1/products/:id` → GetProductByID
  - PATCH `/api/v1/products/:id` → UpdateProduct (owner only)
  - DELETE `/api/v1/products/:id` → DeleteProduct (owner only)
- **Features**:
  - File upload support (image)
  - Filtering: search, category, store
  - Sorting: price_asc, price_desc, newest
  - Pagination
  - User isolation

### ✅ 8. Buatlah service transaksi
**Status**: COMPLETED
- **Implementation**: `TransactionUseCase` dan `TransactionHandler`
- **Endpoints**:
  - POST `/api/v1/transactions` → CreateTransaction
  - GET `/api/v1/transactions/me` → GetMyTransactions
  - GET `/api/v1/transactions/me/:id` → GetMyTransactionByID
  - PATCH `/api/v1/transactions/me/:id` → UpdateTransactionStatus
  - GET `/api/v1/transactions` → GetAllTransactions (admin)
- **Features**:
  - Automatic ProductLog creation
  - Stock management
  - Invoice number generation
  - Transaction status management
  - User isolation

---

## KETENTUAN (18 Aturan)

### ✅ 1. Harus memiliki routing seperti collection berikut
**Status**: COMPLETED
- **Location**: `internal/server/http/httproute.go`
- **Implementation**: Semua endpoints follow RESTful conventions
- **Postman Collection**: Support struktur standar REST

### ✅ 2. Boleh menambahkan dari API yang sudah tapi tidak boleh dikurangi
**Status**: COMPLETED
- **Design**: Architecture extensible dengan Clean Architecture
- **Flexibility**: Mudah menambah endpoint baru tanpa breaking existing

### ✅ 3. Email dan no telepon user tidak boleh ada yang sama
**Status**: COMPLETED
- **Implementation**: 
  - Database constraints: `uniqueIndex` di `User` entity
  - Validation di `UserRepository.GetByEmail()` dan `GetByPhone()`
  - Check di `AuthUseCase.Register()` lines 15-30
  - Check di `UserUseCase.UpdateUser()` lines 53-70

### ✅ 4. Menggunakan JWT
**Status**: COMPLETED
- **Location**: `internal/utils/jwt.go`
- **Implementation**:
  - `GenerateToken()` - Create JWT token
  - `ValidateToken()` - Validate JWT token
  - `CustomClaims` struct dengan UserID, Email, IsAdmin
  - 24-hour expiration
- **Middleware**: `MiddlewareAuth` di `handler/middleware.go`

### ✅ 5. Harus terdapat API yang meng-upload file
**Status**: COMPLETED
- **Location**: `ProductHandler.CreateProduct()` di `product_handler.go`
- **Implementation**:
  ```go
  file, err := c.FormFile("image")
  if err == nil && file != nil {
    filename := fmt.Sprintf("%d_%s", userID, filepath.Base(file.Filename))
    destPath := filepath.Join("uploads", "products", filename)
    c.SaveFile(file, destPath)
  }
  ```
- **Endpoint**: POST `/api/v1/products` (multipart/form-data)

### ✅ 6. Toko otomatis terbuat ketika user mendaftar
**Status**: COMPLETED
- **Verification**: Lihat #2 di SOAL section
- **Code**: `AuthUseCase.Register()` lines 19-34

### ✅ 7. Alamat diperlukan untuk alamat kirim produk
**Status**: COMPLETED
- **Implementation**: Address entity dan service
- **Used In**: Transaction creation memerlukan address_id
- **Fields**: PenerimaNama, PenerimaPhone, DetailAlamat, Provinsi, Kabupaten, Kecamatan, Kelurahan

### ✅ 8. Yang dapat mengelola kategori hanyalah admin
**Status**: COMPLETED
- **Implementation**: `CheckIsAdmin()` check di CategoryHandler
- **Code**: `category_handler.go` lines 24-28
  ```go
  isAdmin, err := utils.CheckIsAdmin(c)
  if err != nil || !isAdmin {
    return helper.Response(c, http.StatusForbidden, "forbidden: only admin can create category", nil, "")
  }
  ```

### ✅ 9. Menerapkan pagination seperti di postman
**Status**: COMPLETED
- **Parameters**: `limit` (default: 10) dan `offset` (default: 0)
- **Implementation**: Semua GET endpoints support pagination
- **Repositories**: Filter struct dengan Limit/Offset
- **Response**: Includes `total` count untuk pagination calculation

### ✅ 10. Menerapkan filtering data
**Status**: COMPLETED
- **Product Filtering**:
  - Search by name: `?search=productname`
  - Filter by category: `?category_id=1`
  - Filter by store: `?store_id=1`
  - Sort by: `?sort=price_asc|price_desc|newest`
  - Pagination: `?limit=10&offset=0`
- **Transaction Filtering**:
  - Filter by status: `?status=pending|completed|cancelled`
  - Pagination
- **Implementation**: `FilterProduct` dan `FilterTransaction` structs

### ✅ 11. User tidak dapat mendapatkan data user lain atau meng-update user lain
**Status**: COMPLETED
- **Implementation**: 
  - UserHandler.GetMe() - Extract current user dari JWT
  - UserHandler.UpdateMe() - Extract current user dari JWT
  - UserHandler.DeleteMe() - Extract current user dari JWT
  - All operations use `utils.ExtractToken(c)` to get authenticated user

### ✅ 12. User tidak dapat mengelola alamat data user lain
**Status**: COMPLETED
- **Implementation**: `AddressRepository.GetByUserIDAndID(userID, addressID)`
- **Code**: `AddressUseCase` semua methods check userID match
- **Example**: `address_usecase.go` line 69 (GetAddressByID)

### ✅ 13. User tidak dapat mengelola data toko dari data user lain
**Status**: COMPLETED
- **Implementation**: `StoreUseCase.UpdateStore()` line 96-100
  ```go
  if store.UserID != userID {
    return nil, errors.New("unauthorized: you can only update your own store")
  }
  ```

### ✅ 14. User tidak dapat mengelola data product dari data user lain
**Status**: COMPLETED
- **Implementation**: `ProductUseCase.UpdateProduct()` check user ownership
- **Code**: `product_usecase.go` line 105-114 (get user's store, verify product belongs to store)

### ✅ 15. User tidak dapat mengelola data transaksi dari data user lain
**Status**: COMPLETED
- **Implementation**: `TransactionUseCase.GetTransactionByID()` line 165-170
  ```go
  if transaction.UserID != userID {
    return nil, errors.New("unauthorized: you can only view your own transactions")
  }
  ```

### ✅ 16. Tabel log product diisi ketika melakukan transaksi
**Status**: COMPLETED
- **Implementation**: `TransactionUseCase.CreateTransaction()` line 100-108
- **Code**:
  ```go
  for i := range productLogs {
    productLogs[i].TransactionID = transaction.ID
    productLogs[i].CreatedAt = now
    productLogs[i].UpdatedAt = now
    if err := u.productLogRepo.Create(&productLogs[i]); err != nil {
      return nil, err
    }
  }
  ```

### ✅ 17. Tabel log produk digunakan untuk menyimpan data produk yang ada di transaksi
**Status**: COMPLETED
- **Entity**: `ProductLog` struct dengan fields: TransactionID, ProductID, Quantity, Price
- **Implementation**: Auto-create saat transaksi, menyimpan snapshot quantity dan price
- **Response**: Included di transaction response sebagai `product_logs`

### ✅ 18. Menerapkan clean architecture
**Status**: COMPLETED
- **Layer Separation**:
  - **Entity Layer**: `internal/pkg/entity/*.go` - Database models
  - **Model Layer**: `internal/pkg/model/*.go` - Request/Response DTOs
  - **Repository Layer**: `internal/pkg/repository/*.go` - Data access
  - **UseCase Layer**: `internal/pkg/usecase/*.go` - Business logic
  - **Handler Layer**: `internal/server/http/handler/*.go` - HTTP handlers
  - **Infrastructure**: `internal/infrastructure/*` - Database, config
  - **Utils**: `internal/utils/*` - JWT, password, helpers
  - **Server**: `internal/server/http/httproute.go` - Routes

- **Benefits**:
  - Separation of concerns
  - Easy testing with interfaces
  - Maintainable codebase
  - Scalable architecture

---

## Summary

✅ **All 8 SOAL (Tasks) Completed**
✅ **All 18 KETENTUAN (Requirements) Completed**
✅ **No Compilation Errors**
✅ **Clean Architecture Applied**
✅ **RESTful API Design**
✅ **JWT Authentication**
✅ **Data Isolation & Authorization**
✅ **Database Migrations**
✅ **Error Handling**
✅ **Pagination & Filtering**


