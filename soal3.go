// Soal 3 — Validasi string hanya karakter <>{}[] tanpa regex.
// Stack: penutup harus cocok dengan pembuka terakhir; tidak silang / salah jenis ([>] invalid).
// Panjang 1–4096.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Validasi mengembalikan true jika kurung seimbang dan setiap penutup cocok dengan pembuka terakhir.
func Validasi(s string) bool {
	n := len(s)
	if n < 1 || n > 4096 {
		return false
	}

	tumpukan := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		c := s[i]
		switch c {
		case '<', '{', '[':
			tumpukan = append(tumpukan, c)
		case '>':
			if len(tumpukan) == 0 || tumpukan[len(tumpukan)-1] != '<' {
				return false
			}
			tumpukan = tumpukan[:len(tumpukan)-1]
		case '}':
			if len(tumpukan) == 0 || tumpukan[len(tumpukan)-1] != '{' {
				return false
			}
			tumpukan = tumpukan[:len(tumpukan)-1]
		case ']':
			if len(tumpukan) == 0 || tumpukan[len(tumpukan)-1] != '[' {
				return false
			}
			tumpukan = tumpukan[:len(tumpukan)-1]
		default:
			return false
		}
	}
	return len(tumpukan) == 0
}

// RunSoal3 membaca satu baris string dari stdin, cetak TRUE atau FALSE.
func RunSoal3() {
	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		return
	}
	s := strings.TrimSpace(sc.Text())
	if Validasi(s) {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}
