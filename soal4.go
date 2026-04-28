// Soal 4 — Cuti pribadi: kuota = 14 − cuti bersama; tahun pertama prorata dari (join+180)+1 s/d 31 Des;
// tidak boleh cuti pribadi sebelum join+180 hari; maksimal 3 hari berturut.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func potongTanggal(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func hariInklusif(a, b time.Time) int {
	a = potongTanggal(a)
	b = potongTanggal(b)
	if b.Before(a) {
		return 0
	}
	return int(b.Sub(a).Hours()/24) + 1
}

// EvaluasiCuti menerapkan aturan cuti karyawan baru pada tahun kalender bergabung (sesuai contoh soal).
func EvaluasiCuti(cutiBersama int, join, rencanaCuti time.Time, durasiHari int) (boleh bool, alasan string) {
	if durasiHari < 1 || durasiHari > 3 {
		return false, "Cuti pribadi maksimal 3 hari berturut-turut"
	}

	join = potongTanggal(join)
	rc := potongTanggal(rencanaCuti)
	loc := join.Location()

	cutiPribadiTahunan := 14 - cutiBersama
	if cutiPribadiTahunan < 0 {
		cutiPribadiTahunan = 0
	}

	awalBolehCuti := join.AddDate(0, 0, 180)
	akhirTahunBergabung := time.Date(join.Year(), 12, 31, 0, 0, 0, 0, loc)

	if rc.Before(awalBolehCuti) {
		return false, "Karena belum 180 hari sejak tanggal join karyawan"
	}

	selesaiCuti := rc.AddDate(0, 0, durasiHari-1)
	if selesaiCuti.After(akhirTahunBergabung) || rc.After(akhirTahunBergabung) {
		return false, "Rencana cuti harus sepenuhnya dalam tahun kalender tahun bergabung"
	}

	// Hari untuk prorata: dari (join+180)+1 s/d 31 Des tahun bergabung (inklusif), cocok contoh 64 hari untuk Mei 2021.
	mulaiHitungKuota := awalBolehCuti.AddDate(0, 0, 1)
	hari := hariInklusif(mulaiHitungKuota, akhirTahunBergabung)
	if mulaiHitungKuota.After(akhirTahunBergabung) {
		hari = 0
	}

	kuota := (hari * cutiPribadiTahunan) / 365
	if durasiHari > kuota {
		return false, fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti", kuota)
	}

	return true, ""
}

// RunSoal4 stdin: baris 1 cuti_bersama (int), baris 2 tanggal masuk YYYY-MM-DD,
// baris 3 tanggal mulai cuti YYYY-MM-DD, baris 4 durasi (hari).
func RunSoal4() {
	sc := bufio.NewScanner(os.Stdin)
	var lines []string
	for sc.Scan() {
		t := strings.TrimSpace(sc.Text())
		if t != "" {
			lines = append(lines, t)
		}
	}
	if len(lines) < 4 {
		return
	}
	cb, err := strconv.Atoi(lines[0])
	if err != nil {
		return
	}
	join, err := time.ParseInLocation("2006-01-02", lines[1], time.Local)
	if err != nil {
		return
	}
	rc, err := time.ParseInLocation("2006-01-02", lines[2], time.Local)
	if err != nil {
		return
	}
	dur, err := strconv.Atoi(lines[3])
	if err != nil {
		return
	}
	boleh, alasan := EvaluasiCuti(cb, join, rc, dur)
	if boleh {
		fmt.Println("True")
		return
	}
	fmt.Println("False")
	fmt.Println("Alasan: " + alasan)
}
