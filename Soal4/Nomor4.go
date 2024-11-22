package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {

	fmt.Println("Masukkan karakter (akhiri dengan '.'): ")
	var char rune
	for *n = 0; *n < NMAX; *n++ {
		fmt.Scanf("%c", &char)
		if char == '.' {
			break
		}
		(*t)[*n] = char
	}
}

func cetakArray(t tabel, n int) {

	fmt.Println("Isi array:")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func cekPalindrome(t tabel, n int) bool {

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if t[i] != t[j] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	cetakArray(tab, m)

	balikanArray(&tab, m)
	fmt.Println("Array setelah dibalik:")
	cetakArray(tab, m)

	if cekPalindrome(tab, m) {
		fmt.Println("Array adalah palindrom.")
	} else {
		fmt.Println("Array bukan palindrom.")
	}
}
