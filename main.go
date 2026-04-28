package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go run . <1|2|3|4>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "1":
		stderrInteractiveHint("1")
		RunSoal1()
	case "2":
		stderrInteractiveHint("2")
		RunSoal2()
	case "3":
		stderrInteractiveHint("3")
		RunSoal3()
	case "4":
		stderrInteractiveHint("4")
		RunSoal4()
	default:
		fmt.Fprintln(os.Stderr, "nomor soal harus 1, 2, 3, atau 4")
		os.Exit(1)
	}
}
