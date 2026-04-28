// Soal 1 — Cocokkan string (case insensitive).
// Kelompok pertama: indeks i terkecil di mana ada j < i dengan string sama (abaikan kapital).
// Tampilkan semua nomor string (1-based) dalam kelompok itu, atau false.
// Tanpa fungsi array search/filter bawaan; for/if manual.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CocokkanString mengembalikan indeks 1-based yang termasuk kelompok cocokan pertama, atau ada=false.
func CocokkanString(barisan []string) (indeks []int, ada bool) {
	if len(barisan) < 2 {
		return nil, false
	}

	norm := make([]string, len(barisan))
	for i := range barisan {
		norm[i] = hurufKecilManual(barisan[i])
	}

	var pola string
	ditemukan := false
	for i := 1; i < len(norm); i++ {
		for j := 0; j < i; j++ {
			if samaStringManual(norm[i], norm[j]) {
				pola = norm[i]
				ditemukan = true
				break
			}
		}
		if ditemukan {
			break
		}
	}
	if !ditemukan {
		return nil, false
	}

	var keluar []int
	for k := 0; k < len(norm); k++ {
		if samaStringManual(norm[k], pola) {
			keluar = append(keluar, k+1)
		}
	}
	if len(keluar) < 2 {
		return nil, false
	}
	urutManual(keluar)
	return keluar, true
}

func hurufKecilManual(s string) string {
	b := make([]byte, len(s))
	copy(b, s)
	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = b[i] - 'A' + 'a'
		}
	}
	return string(b)
}

func samaStringManual(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func urutManual(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

// RunSoal1 stdin: baris pertama N, lalu N baris string (satu string per baris).
func RunSoal1() {
	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		fmt.Println("false")
		return
	}
	n, err := strconv.Atoi(strings.TrimSpace(sc.Text()))
	if err != nil || n <= 0 {
		fmt.Println("false")
		return
	}
	barisan := make([]string, 0, n)
	for len(barisan) < n && sc.Scan() {
		barisan = append(barisan, strings.TrimSpace(sc.Text()))
	}
	if len(barisan) != n {
		fmt.Println("false")
		return
	}
	idx, ok := CocokkanString(barisan)
	if !ok {
		fmt.Println("false")
		return
	}
	for i := 0; i < len(idx); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(idx[i])
	}
	fmt.Println()
}
