package main

import (
	"fmt"
	"os"
)

// stderrInteractiveHint menjelaskan ke stderr jika stdin adalah terminal (karakter device),
// supaya program tidak terlihat "macet" saat menunggu input dari keyboard.
func stderrInteractiveHint(nomorSoal string) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return
	}
	if fi.Mode()&os.ModeCharDevice == 0 {
		return
	}
	fmt.Fprintf(os.Stderr,
		"[stdin interaktif] Soal %s membaca input dari keyboard. "+
			"Ketik sesuai README lalu Enter; akhiri dengan Ctrl+D (EOF). "+
			"Alternatif tanpa mengetik: printf '…' | go run . %s\n",
		nomorSoal, nomorSoal)
}
