# Soal Test Tahap 1 — Golang

Satu executable: nomor soal dipilih lewat **argumen** (`1`–`4`), masukan dibaca dari **stdin**, keluaran ke **stdout**. Saat Anda menjalankan dari terminal tanpa pipe, petunjuk singkat boleh ditulis ke **stderr** (lihat [stdin dan terminal](#stdin-dan-terminal)).

## Prasyarat

- [Go](https://go.dev/dl/) **1.22** atau lebih baru (`go version`)

---

## Menjalankan

Di akar repository (`go.mod`):

```bash
go run . <1|2|3|4>
```

| Argumen | Keterangan |
|--------|------------|
| `1` | Pencocokan string (case insensitive, kelompok duplikat pertama) |
| `2` | Kembalian kasir + pecahan uang |
| `3` | Validasi kurung `<>{}[]` |
| `4` | Cuti karyawan |

Tanpa argumen, atau bukan `1`–`4`: pesan penggunaan / error di **stderr**.

**Jangan** menjalankan hanya satu berkas, misalnya `go run soal1.go`. `func main` ada di `main.go`; semua `soal*.go` ikut dikompilasi sebagai satu paket `main`. Perintah itu memicu error: `function main is undeclared in the main package`.

### Build binary

```bash
go build -o soaltest .
./soaltest 1
```

---

## stdin dan terminal

Program **selalu membaca dari stdin**. Jika Anda hanya mengetik `go run . 2` dan tidak mengirim data (tanpa pipe / redirect), proses akan **menunggu** sampai stdin ditutup:

- **macOS / Linux:** akhiri input dengan **Ctrl+D**
- **Windows (CMD):** **Ctrl+Z** lalu Enter

Jika stdin adalah **keyboard** (bukan pipe/file), program dapat mencetak satu baris petunjuk ke **stderr** supaya tidak terlihat “macet” (lihat `stdin_hint.go`).

**Disarankan untuk uji cepat:** gunakan pipe:

```bash
printf '4\nabcd\nacbd\naaab\nacbd\n' | go run . 1
```

---

## Format input & output per soal

### Soal 1

**stdin**

1. Baris pertama: bilangan bulat `N`.
2. `N` baris berikutnya: tiap baris satu string (`s1` … `sN`).

**stdout**

- Ada minimal dua string yang sama (**tanpa membedakan huruf besar/kecil**): satu baris, nomor baris (1-based) dalam **kelompok pertama** yang mengandung duplikat (semua indeks dengan nilai yang sama dengan kelompok itu), dipisah spasi, urut naik.
- Tidak ada duplikat: satu baris `false`.

### Soal 2

**stdin**

- Dua baris tidak kosong pertama: **total belanja**, lalu **uang dibayar**.
- Boleh angka saja (`700649`) atau teks seperti `Rp 700.649`: semua **digit** pada baris digabung untuk mendapat nilai.

**stdout**

- Jika dibayar &lt; total: satu baris `False, kurang bayar`.
- Jika cukup: beberapa baris — kembalian sebenarnya, kembalian dibulatkan **ke bawah** ke kelipatan **Rp100**, lalu `Pecahan uang:` dan tiap pecahan yang terpakai (`lembar` untuk nominal ≥ Rp1.000, `koin` untuk 500 / 200 / 100).

### Soal 3

**stdin**

- Satu baris string; hanya karakter `<>{}[]`, panjang 1–4096.

**stdout**

- Satu baris: `TRUE` atau `FALSE`.

### Soal 4

**stdin** (empat baris)

1. Cuti bersama (bilangan bulat).
2. Tanggal bergabung: `YYYY-MM-DD`.
3. Tanggal mulai rencana cuti: `YYYY-MM-DD`.
4. Durasi (hari).

**stdout**

- Jika boleh: satu baris `True`.
- Jika tidak: baris `False`, lalu baris `Alasan: …`.

---

## Contoh perintah (pipe)

```bash
# Soal 1 → contoh: 2 4
printf '4\nabcd\nacbd\naaab\nacbd\n' | go run . 1

# Soal 2 → kembalian & pecahan (format bahasa Indonesia)
printf '700649\n800000\n' | go run . 2

# Soal 3 (kurung seimbang)
printf '[{}<>]\n' | go run . 3

# Soal 4
printf '2\n2024-01-01\n2025-06-01\n3\n' | go run . 4
```

Input dari file:

```bash
go run . 1 < input_soal1.txt
```

---

## Struktur berkas

| Berkas | Peran |
|--------|--------|
| `main.go` | Argumen `1`–`4`, memanggil `RunSoal*` |
| `stdin_hint.go` | Petunjuk ke stderr jika stdin terminal |
| `soal1.go` | `CocokkanString`, `RunSoal1` |
| `soal2.go` | `HitungKembalian`, `RunSoal2` |
| `soal3.go` | `Validasi`, `RunSoal3` |
| `soal4.go` | `EvaluasiCuti`, `RunSoal4` |
| `go.mod` | Modul Go |

---

## Kompilasi & tes

```bash
go build -o soaltest .
go test ./...
```

*(Belum ada berkas `*_test.go`; `go test` memverifikasi paket dapat dikompilasi.)*

---

## Dokumen resmi vs kode

Jika PDF / soal resmi mendefinisikan format input/output lain, sesuaikan fungsi `RunSoal1` … `RunSoal4` di `soal*.go`, atau sesuaikan fungsi inti (`CocokkanString`, `HitungKembalian`, dll.) tanpa mengubah kontrak CLI `go run . <n>` kecuali memang diperlukan.
