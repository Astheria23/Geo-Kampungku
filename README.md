# Dataset Jalan Kertawangi – Cisarua, KBB

Dokumen ini merangkum isi berkas `kertawangi.geojson`, yaitu jaringan jalan yang dipetakan di Desa Kertawangi, Kecamatan Cisarua, Kabupaten Bandung Barat (Provinsi Jawa Barat).

## Ikhtisar Dataset

- **Sumber**: OpenStreetMap (ekspor *overpass-turbo* `2025-10-17T01:36:58Z`).
- **Jumlah fitur**: 272 ruas jalan dan jalur.
- **Tipe geometri**: seluruhnya `LineString` (polyline).
- **Batas koordinat (bbox)**: `[107.5475062, -6.819106]` hingga `[107.5815871, -6.7696358]` (lon/lat WGS84).
- **Wilayah administrasi**: Kertawangi & sekitarnya di Kecamatan Cisarua, Kabupaten Bandung Barat.

## Statistik Utama

| Statistik                         | Nilai |
|-----------------------------------|------:|
| Total fitur `Feature`             |   272 |
| Fitur dengan nama jalan (`name`)  |    21 |
| Jenis `highway` berbeda           |     9 |
| Nilai `surface` terdata           |  `asphalt` (5), `earth` (1) |
| Fitur dengan atribut `lanes`      |    13 |
| Fitur bertanda `motorcar=private` |    24 |

Distribusi kategori `highway`:

| Kategori       | Jumlah | Persentase ± |
|----------------|-------:|-------------:|
| `residential`  |    126 |       46.3% |
| `service`      |     80 |       29.4% |
| `living_street`|     25 |        9.2% |
| `path`         |     15 |        5.5% |
| `tertiary`     |      8 |        2.9% |
| `primary`      |      5 |        1.8% |
| `footway`      |      5 |        1.8% |
| `unclassified` |      4 |        1.5% |
| `track`        |      4 |        1.5% |

## Atribut Penting pada `properties`

- `@id` – ID resmi objek di OSM (contoh `way/318678694`).
- `highway` – tipe jalan sesuai skema OSM (`residential`, `service`, dll.).
- `name` – nama ruas (tersedia pada 21 fitur; contoh: *Jalan Kolonel Masturi*, *Gang Masjid 2*, *Jalur Pendakian Gunung Burangrang via Komando*).
- `motorcar` / `motorcycle` – pembatasan akses kendaraan bermotor pada sebagian jalur fasilitas.
- `access` – aturan akses umum (misal `private`).
- `lanes` – jumlah lajur (hingga 4 pada ruas utama).
- `surface` – material permukaan (aspal mendominasi ruas berlabel).
- Atribut lain yang kadang muncul: `oneway`, `bridge`, `layer`, `ref`, `maxspeed`, `foot`.

## Contoh Fitur Representatif

```json
{
  "type": "Feature",
  "properties": {
    "@id": "way/318678694",
    "highway": "primary",
    "name": "Jalan Kolonel Masturi",
    "lanes": "2",
    "surface": "asphalt",
    "source": "GPS"
  },
  "geometry": {
    "type": "LineString",
    "coordinates": [
      [107.557082, -6.8171332],
      [107.5572749, -6.816342],
      [107.558808, -6.8133731],
      [107.5692355, -6.8090175],
      [107.5701576, -6.8043291]
    ]
  }
}
```

## Cara Menggunakan Dataset
1. **Visualisasi Peta** – Muat `kertawangi.geojson` di QGIS/ArcGIS atau peta berbasis Leaflet/Mapbox untuk melihat jaringan jalan secara langsung.
2. **Analisis** – Gunakan pustaka GIS (Turf.js, PostGIS, GeoPandas) untuk menghitung panjang ruas, mengelompokkan berdasarkan `highway`, atau mengevaluasi aksesibilitas.
3. **Integrasi Backend** – Paket ini kompatibel dengan payload `LineString` pada API Go Fiber (`POST /api/polylines`). Pilih fitur tertentu, ubah menjadi struktur GeoJSON minimal, kemudian kirimkan ke endpoint.


