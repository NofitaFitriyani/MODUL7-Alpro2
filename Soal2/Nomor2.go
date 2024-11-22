package main
import (
	"fmt"
	"math"
)

func tampilSemua(array []int) {
	fmt.Println("Isi array:", array)
}

func tampilGanjil(array []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilGenap(array []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(array); i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func tampilKelipatan(array []int, x int) {
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := x; i < len(array); i += x {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}

func hapusIndeks(array []int, indeks int) []int {
	fmt.Printf("Menghapus elemen pada indeks %d\n", indeks)
	return append(array[:indeks], array[indeks+1:]...)
}

func rataRata(array []int) float64 {
	total := 0
	for _, val := range array {
		total += val
	}
	return float64(total) / float64(len(array))
}

func standarDeviasi(array []int) float64 {
	mean := rataRata(array)
	var sum float64
	for _, val := range array {
		sum += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(sum / float64(len(array)))
}

func frekuensi(array []int, nilai int) int {
	count := 0
	for _, val := range array {
		if val == nilai {
			count++
		}
	}
	return count
}

func main() {
	var N, x, hapusIdx, cariFrekuensi int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&N)

	array := make([]int, N)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	tampilSemua(array)
	tampilGanjil(array)
	tampilGenap(array)

	fmt.Print("Masukkan nilai x untuk menampilkan elemen dengan indeks kelipatan x: ")
	fmt.Scan(&x)
	tampilKelipatan(array, x)

	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&hapusIdx)
	if hapusIdx >= 0 && hapusIdx < len(array) {
		array = hapusIndeks(array, hapusIdx)
		tampilSemua(array)
	} else {
		fmt.Println("Indeks yang dimasukkan tidak valid!")
	}

	fmt.Printf("Rata-rata elemen array: %.2f\n", rataRata(array))

	fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi(array))

	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cariFrekuensi)
	fmt.Printf("Frekuensi %d di dalam array: %d\n", cariFrekuensi, frekuensi(array, cariFrekuensi))
}