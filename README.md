# Dokumentasi GeoJSON – Kertawangi, Cisarua (KBB)

Repositori ini menyertakan berkas `kertawangi.geojson` yang berisi jaringan jalan di wilayah Kertawangi, Kecamatan Cisarua, Kabupaten Bandung Barat, Provinsi Jawa Barat. Dokumen ini menjelaskan struktur GeoJSON, header yang digunakan, serta sintaks penting di dalam berkas tersebut.

## Ringkasan Lokasi

- **Wilayah**: Desa/Kelurahan Kertawangi, Kecamatan Cisarua, Kabupaten Bandung Barat.
- **Koordinat perkiraan**: garis bujur 107.55°–107.57° BT, garis lintang -6.81°– -6.79° LS.
- **Sumber data**: OpenStreetMap melalui *overpass-turbo* per tanggal `2025-10-17T01:36:58Z`.

## Pengantar GeoJSON

GeoJSON adalah format berbasis JSON untuk menyimpan dan bertukar data geospasial. Beberapa konsep kunci:

- **FeatureCollection**: struktur utama yang menampung sekumpulan *feature*.
- **Feature**: representasi satu entitas spasial (misalnya satu ruas jalan) yang terdiri dari `properties` (atribut) dan `geometry` (bentuk geometri).
- **Geometry**: tipe geometri seperti `Point`, `LineString`, `Polygon`. Dalam berkas ini mayoritas berupa `LineString` karena merepresentasikan jalur/jalan.
- **Koordinat**: selalu berformat `[longitude, latitude]` dalam sistem koordinat WGS84 (EPSG:4326).

Struktur umum GeoJSON:

```json
{
	"type": "FeatureCollection",
	"features": [
		{
			"type": "Feature",
			"properties": { ... },
			"geometry": {
				"type": "LineString",
				"coordinates": [ [lon, lat], ... ]
			}
		}
	]
}
```

## Header dan Metadata pada `kertawangi.geojson`

| Field         | Nilai contoh                                                   | Penjelasan                                                                                       |
|---------------|----------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| `type`        | `FeatureCollection`                                            | Menandakan berkas berisi kumpulan feature.                                                        |
| `generator`   | `overpass-turbo`                                              | Alat yang digunakan untuk mengekstraksi data dari OpenStreetMap.                                  |
| `copyright`   | `The data included ... under ODbL.`                           | Informasi lisensi ODbL dari OpenStreetMap.                                                        |
| `timestamp`   | `2025-10-17T01:36:58Z`                                        | Waktu ekspor data terakhir.                                                                       |
| `features`    | Array berisi objek `Feature`                                  | Setiap feature mewakili satu ruas jalan atau elemen infrastruktur lain di Kertawangi.              |

## Struktur Feature

Setiap feature memiliki dua bagian utama:

1. **`properties`** – atribut non-spasial seperti:
	 - `@id`: ID unik OSM (contoh `way/318678694`).
	 - `name`: nama jalan (misal `Jalan Kolonel Masturi`).
	 - `highway`: klasifikasi jalan (`primary`, `residential`, `service`, dll.).
	 - Atribut lain: `lanes`, `surface`, `oneway`, `ref`, hingga pengaturan khusus (misal `cycleway:both`).
2. **`geometry`** – representasi spasial:
	 - `type`: mayoritas `LineString` (jalur polyline). Ada kemungkinan `Point` atau `Polygon` bila dataset diperluas.
	 - `coordinates`: daftar titik `[longitude, latitude]` yang membentuk geometri.

Contoh feature dari `kertawangi.geojson`:

```json
{
	"type": "Feature",
	"properties": {
		"@id": "way/318678694",
		"highway": "primary",
		"name": "Jalan Kolonel Masturi",
		"lanes": "2",
		"surface": "asphalt"
	},
	"geometry": {
		"type": "LineString",
		"coordinates": [
			[107.557082, -6.8171332],
			[107.5572749, -6.816342],
			[107.5577432, -6.8145058],
			[107.558808, -6.8133731],
			[107.5692355, -6.8090175],
			[107.5701576, -6.8043291]
		]
	}
}
```

### Catatan Sintaks `LineString`

- Koordinat disusun berurutan sepanjang jalur jalan dari titik awal hingga akhir.
- Nilai berupa angka desimal (floating point) dalam derajat.
- *LineString* minimal memiliki dua titik; semakin banyak titik, semakin detail kurva jalan di peta.

## Ringkasan Jenis Jalan di Dataset

Walaupun dataset sangat besar, sebagian besar fitur dapat dikategorikan:

| Kategori `highway` | Deskripsi singkat                                             | Contoh nama jalan                          |
|--------------------|----------------------------------------------------------------|---------------------------------------------|
| `primary`          | Jalan utama penghubung antarwilayah                            | Jalan Kolonel Masturi                       |
| `secondary`        | Jalan arteri tingkat menengah                                  | Jalan Terusan Kolonel Masturi               |
| `tertiary`         | Jalan kolektor di dalam kecamatan/kelurahan                    | Jalan Sersan Bajuri                         |
| `residential`      | Jalan lingkungan permukiman                                    | Jalan Kampung Gunung Mas, Jalan Puspa Raya  |
| `service`          | Jalan akses terbatas (misal ke fasilitas atau area parkir)     | Jalan menuju restoran/fasilitas wisata      |
| `track` / `path`   | Jalur pertanian atau pejalan kaki                              | Jalur kebun atau setapak kampung            |

