# Sistem Informasi GoLang

Proyek ini adalah sebuah aplikasi RESTful API yang dibangun menggunakan Go dengan framework Fiber. Aplikasi ini menyediakan fitur untuk mengelola pengguna dan item dalam basis data, termasuk registrasi, login, dan manajemen item.

## Struktur Folder

Berikut adalah struktur folder dari proyek ini:

```txt
Sistem-Infromasi_GoLang/
│
├── main.go
├── controllers/
│   ├── auth_controller.go
│   ├── item_controller.go
│   └── user_controller.go
├── models/
│   ├── item_model.go
│   └── user_model.go
├── middlewares/
│   └── auth_middleware.go
├── services/
│   ├── auth_service.go
│   ├── item_service.go
│   └── user_service.go
├── utils/
│   ├── jwt.go
│   └── password.go
├── tests/
│   ├── user_test.go
│   └── item_test.go
└── go.mod
```

## Deskripsi File

**`main.go`**

File utama yang menjalankan aplikasi. Mengatur koneksi ke database dan mendefinisikan rute-rute API.

**`controllers/`**

Berisi logika kontrol untuk menangani permintaan HTTP. Terdapat tiga file controller:

- auth_controller.go: Mengelola registrasi dan login pengguna.
- item_controller.go: Mengelola item yang dimiliki pengguna.
- user_controller.go: Mengelola informasi pengguna.

**`models/`**

Berisi definisi model untuk pengguna dan item:

- item_model.go: Model untuk item.
- user_model.go: Model untuk pengguna.

**`middlewares/`**

Berisi middleware untuk autentikasi:

- auth_middleware.go: Memeriksa token JWT untuk setiap permintaan yang membutuhkan autentikasi.

**`services/`**

Berisi logika bisnis yang berinteraksi dengan model:

- auth_service.go: Mengelola logika autentikasi.
- item_service.go: Mengelola logika untuk item.
- user_service.go: Mengelola logika untuk pengguna.

**`utils/`**

Berisi fungsi utilitas seperti pengelolaan token JWT dan hashing password:

- jwt.go: Mengelola pembuatan dan validasi token JWT.
- password.go: Mengelola hashing password.

**`tests/`**

Berisi pengujian unit untuk memastikan fungsionalitas aplikasi:

- user_test.go: Pengujian untuk layanan pengguna.
- item_test.go: Pengujian untuk layanan item.

## Konfigurasi

`go.mod` File ini mendefinisikan modul Go dan dependensi yang diperlukan:

```go
module restful-api

go 1.23.2

require (
    github.com/gofiber/fiber/v2 v2.52.5
    github.com/golang-jwt/jwt/v5 v5.0.0
    golang.org/x/crypto v0.23.0
    gorm.io/driver/mysql v1.5.7
    gorm.io/gorm v1.25.12
)
```

## Cara Menjalankan Proyek

1. Pastikan Anda telah menginstal Go, MySQL, dan Redis.
2. Buat database baru di MySQL dengan nama database_si.
3. Jalankan perintah berikut untuk menginstal dependensi:

```bash
go mod tidy
```

4. Jalankan aplikasi dengan perintah:

```bash
go run main.go
```

## Rute API

Berikut adalah rute-rute utama yang tersedia dalam aplikasi:

### Autentikasi Pengguna

- POST `/registrasi`: Mendaftar pengguna baru.
- POST `/login`: Masuk sebagai pengguna.

### Manajemen Pengguna (Memerlukan autentikasi)

- GET `/api/user`: Mendapatkan informasi pengguna.
- PUT `/api/user`: Memperbarui informasi pengguna.
- DELETE `/api/user`: Menghapus pengguna.

### Manajemen Item (Memerlukan autentikasi)

- GET `/api/item`: Mendapatkan daftar item milik pengguna.
- POST `/api/item`: Menambahkan item baru.
- PUT `/api/item/`:id: Memperbarui item berdasarkan ID.
- DELETE `/api/item/`:id: Menghapus item berdasarkan ID.

## Kesimpulan
Proyek ini merupakan implementasi sederhana dari RESTful API menggunakan Go, menyediakan fungsionalitas dasar untuk manajemen pengguna dan item dengan keamanan melalui autentikasi token JWT.
