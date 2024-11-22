package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&klubB)

	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d - Masukkan skor %s: ", i, klubA)
		fmt.Scan(&skorA)
		fmt.Printf("Pertandingan %d - Masukkan skor %s: ", i, klubB)
		fmt.Scan(&skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Skor tidak valid. Pertandingan selesai.")
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
			fmt.Printf("Hasil %d: %s\n", i, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
			fmt.Printf("Hasil %d: %s\n", i, klubB)
		} else {
			fmt.Printf("Hasil %d: Draw\n", i)
		}
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for _, klub := range pemenang {
		fmt.Println(klub)
	}
}
