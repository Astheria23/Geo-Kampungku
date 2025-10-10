# GeoJSON Polyline API (Go Fiber)

Layanan backend untuk mengelola polyline jalan yang disimpan sebagai GeoJSON LineString di MongoDB. Dibangun menggunakan Go, Fiber, dan driver resmi MongoDB.

## Fitur

- Endpoint CRUD untuk polyline jalan (geometri `LineString`)
- Penyimpanan MongoDB dengan skema kompatibel GeoJSON
- Filter berdasarkan wilayah melalui parameter query
- Endpoint pemeriksaan kesehatan (`/health`)

## Mulai Cepat

### Prasyarat

- Go 1.22+
- Instance MongoDB (lokal atau remote)

### Setup

1. Salin contoh file environment dan sesuaikan nilainya jika perlu:

```bash
cp .env.example .env
```

2. Pasang dependensi:

```bash
go mod tidy
```

3. Jalankan server:

```bash
go run ./...
```

Secara bawaan API berjalan pada port `:3000`.

### Isi Data Contoh (Seeder)

Untuk mengisi data awal polyline ke MongoDB, jalankan perintah berikut:

```bash
go run ./cmd/seeder
```

Seeder akan melakukan upsert berdasarkan kombinasi `name` dan `region`, sehingga aman dijalankan berulang kali tanpa menduplikasi data.

## Endpoint API

| Method | Path                  | Deskripsi                                      |
|--------|-----------------------|------------------------------------------------|
| GET    | `/health`             | Mengembalikan status layanan                   |
| POST   | `/api/polylines`      | Membuat polyline baru                          |
| GET    | `/api/polylines`      | Menampilkan daftar polyline (opsional filter `region`)|
| GET    | `/api/polylines/:id`  | Mengambil polyline berdasarkan ID              |
| PUT    | `/api/polylines/:id`  | Memperbarui polyline                           |
| DELETE | `/api/polylines/:id`  | Menghapus polyline                             |

### Contoh Payload
```json
{
  "name": "Main Street",
  "description": "Primary road through downtown",
  "region": "Kecamatan A",
  "geometry": {
    "type": "LineString",
    "coordinates": [
      [110.4321, -7.1234],
      [110.4325, -7.1238],
      [110.4330, -7.1242]
    ]
  }
}
```

## Pengujian dengan Postman

### 1. Atur base URL

Buat collection atau variabel environment di Postman bernama `baseUrl` dan isi dengan alamat server, misalnya:

```
http://localhost:3000
```

Semua request berikut menggunakan variabel ini serta header `Content-Type: application/json` ketika mengirim body.

### 2. Health check

- **Metode:** `GET`
- **URL:** `{{baseUrl}}/health`
- **Respon yang diharapkan:**

```json
{
  "status": "ok"
}
```

### 3. Buat polyline

- **Metode:** `POST`
- **URL:** `{{baseUrl}}/api/polylines`
- **Body (raw JSON):**

```json
{
  "name": "Main Street",
  "description": "Primary road through downtown",
  "region": "Kecamatan A",
  "geometry": {
    "type": "LineString",
    "coordinates": [
      [110.4321, -7.1234],
      [110.4325, -7.1238],
      [110.4330, -7.1242]
    ]
  }
}
```

- **Respon berhasil (201):**

```json
{
  "id": "6524cf7eecf4f6156f6f5a10",
  "name": "Main Street",
  "description": "Primary road through downtown",
  "region": "Kecamatan A",
  "geometry": {
    "type": "LineString",
    "coordinates": [
      [110.4321, -7.1234],
      [110.4325, -7.1238],
      [110.433, -7.1242]
    ]
  },
  "createdAt": "2025-10-10T04:12:06.122Z",
  "updatedAt": "2025-10-10T04:12:06.122Z"
}
```

### 4. Daftar polyline

- **Metode:** `GET`
- **URL:** `{{baseUrl}}/api/polylines`
- **Query opsional:** `?region=Kecamatan%20A`
- **Respon berhasil (200):** array objek polyline. Gunakan filter region untuk membatasi hasil.

### 5. Ambil berdasarkan ID

- **Metode:** `GET`
- **URL:** `{{baseUrl}}/api/polylines/{{polylineId}}`
- **Catatan:** Ganti `{{polylineId}}` dengan `_id` dari respon create/daftar.
- **Respon berhasil (200):** satu objek polyline.

### 6. Perbarui polyline

- **Metode:** `PUT`
- **URL:** `{{baseUrl}}/api/polylines/{{polylineId}}`
- **Body (raw JSON):** kirim payload lengkap untuk mengganti field yang dapat diubah, contohnya:

```json
{
  "name": "Main Street - Updated",
  "description": "Widened road section",
  "region": "Kecamatan A",
  "geometry": {
    "type": "LineString",
    "coordinates": [
      [110.4321, -7.1234],
      [110.4328, -7.1239],
      [110.4333, -7.1245]
    ]
  }
}
```

- **Respon berhasil (200):** objek polyline yang sudah diperbarui.

### 7. Hapus polyline

- **Metode:** `DELETE`
- **URL:** `{{baseUrl}}/api/polylines/{{polylineId}}`
- **Respon berhasil (204):** tanpa body. Request GET selanjutnya akan mengembalikan `404`.

### Respon error umum

| Status | Kondisi                                            | Contoh pesan                   |
|--------|----------------------------------------------------|--------------------------------|
| 400    | JSON tidak valid, field wajib kosong, ID tidak sah | `"invalid JSON payload"`       |
| 404    | Data tidak ditemukan                               | `"polyline not found"`         |
| 500    | Kesalahan database atau server                     | `"mongo: unable to connect"`   |

Saat debugging, cek log server Fiber untuk konteks tambahan.

## Pengujian

Jalankan unit test dengan:

```bash
go test ./...
```

## Struktur Proyek

```
backend/
├── cmd/
│   └── seeder/      # Program seeder untuk data contoh
├── config/          # Helper untuk memuat environment
├── database/        # Logika koneksi MongoDB
├── handlers/        # Handler HTTP Fiber
├── models/          # Model GeoJSON dan domain
├── routes/          # Registrasi route
├── .env.example     # Contoh konfigurasi environment
├── go.mod / go.sum  # Modul Go dan dependensi
└── main.go          # Entry point aplikasi
```
