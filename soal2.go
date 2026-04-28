// Soal 2 — Hitung kembalian, dibulatkan ke bawah Rp100 (integer kelipatan 100).
// Pecahan: 100.000 … 100. Uang kertas ≥1.000 sebagai lembar; 500, 200, 100 sebagai koin.
// Jika bayar < total: False, kurang bayar.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// pecahanTersedia urutan nominal turun (rupiah).
var pecahanTersedia = []int64{
	100_000, 50_000, 20_000, 10_000, 5_000, 2_000, 1_000, 500, 200, 100,
}

// HasilKembalian menyimpan detail output kasir.
type HasilKembalian struct {
	KurangBayar          bool
	KembalianAsli        int64
	KembalianDibulatkan  int64
	JumlahPerNominal     []struct {
		Nominal int64
		Jumlah  int
	}
}

// HitungKembalian menghitung kembalian dan rincian pecahan (greedy pada urutan nominal resmi).
func HitungKembalian(totalBelanja, dibayar int64) HasilKembalian {
	if dibayar < totalBelanja {
		return HasilKembalian{KurangBayar: true}
	}

	kembalianAsli := dibayar - totalBelanja
	dibulatkan := (kembalianAsli / 100) * 100

	var rinci []struct {
		Nominal int64
		Jumlah  int
	}
	sisa := dibulatkan
	for _, u := range pecahanTersedia {
		if u <= 0 || sisa < u {
			continue
		}
		n := int(sisa / u)
		if n <= 0 {
			continue
		}
		rinci = append(rinci, struct {
			Nominal int64
			Jumlah  int
		}{Nominal: u, Jumlah: n})
		sisa -= int64(n) * u
	}

	return HasilKembalian{
		KurangBayar:         false,
		KembalianAsli:       kembalianAsli,
		KembalianDibulatkan: dibulatkan,
		JumlahPerNominal:    rinci,
	}
}

// formatTitikRibu menulis bilangan bulat dengan pemisah ribuan '.' (gaya Indonesia).
func formatTitikRibu(n int64) string {
	if n < 0 {
		return "-" + formatTitikRibu(-n)
	}
	s := strconv.FormatInt(n, 10)
	if len(s) <= 3 {
		return s
	}
	var bagian []string
	for len(s) > 3 {
		bagian = append([]string{s[len(s)-3:]}, bagian...)
		s = s[:len(s)-3]
	}
	bagian = append([]string{s}, bagian...)
	return strings.Join(bagian, ".")
}

func satuannya(nominal int64) string {
	if nominal >= 1000 {
		return "lembar"
	}
	return "koin"
}

// ambilAngkaDariBaris mengambil digit beruntun dari satu baris (mis. dari teks "Rp 700.649").
func ambilAngkaDariBaris(barisan string) (nilai int64, ok bool) {
	var digit strings.Builder
	for _, r := range barisan {
		if r >= '0' && r <= '9' {
			digit.WriteRune(r)
		}
	}
	if digit.Len() == 0 {
		return 0, false
	}
	v, err := strconv.ParseInt(digit.String(), 10, 64)
	if err != nil {
		return 0, false
	}
	return v, true
}

// tulisOutputSoal2 mencetak sesuai format soal (bahasa Indonesia).
func tulisOutputSoal2(h HasilKembalian) {
	if h.KurangBayar {
		fmt.Println("False, kurang bayar")
		return
	}

	fmt.Printf("Kembalian yang harus diberikan kasir: %s,\n", formatTitikRibu(h.KembalianAsli))
	fmt.Printf("dibulatkan menjadi %s\n", formatTitikRibu(h.KembalianDibulatkan))
	fmt.Println("Pecahan uang:")
	for _, it := range h.JumlahPerNominal {
		fmt.Printf("%d %s %s\n", it.Jumlah, satuannya(it.Nominal), formatTitikRibu(it.Nominal))
	}
}

// RunSoal2 stdin: dua baris pertama yang bisa diparse menjadi nominal (angka saja atau teks berisi Rp …).
func RunSoal2() {
	sc := bufio.NewScanner(os.Stdin)
	var baris []string
	for sc.Scan() {
		t := strings.TrimSpace(sc.Text())
		if t != "" {
			baris = append(baris, t)
		}
	}
	if len(baris) < 2 {
		return
	}
	total, ok1 := ambilAngkaDariBaris(baris[0])
	bayar, ok2 := ambilAngkaDariBaris(baris[1])
	if !ok1 || !ok2 {
		return
	}
	h := HitungKembalian(total, bayar)
	tulisOutputSoal2(h)
}
