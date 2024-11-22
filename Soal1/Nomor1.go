package main

import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	center Titik
	radius int
}

func jarakKuadrat(p, q Titik) int {
	return (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)
}

func diDalam(c Lingkaran, p Titik) bool {
	return jarakKuadrat(c.center, p) <= c.radius*c.radius
}

func main() {
	var cx1, cy1, r1 int
	fmt.Print("Masukkan 1: ")
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 int
	fmt.Print("Masukkan 2: ")
	fmt.Scan(&cx2, &cy2, &r2)

	var px, py int
	fmt.Print("Masukkan koordinat titik: ")
	fmt.Scan(&px, &py)

	lingkaran1 := Lingkaran{Titik{cx1, cy1}, r1}
	lingkaran2 := Lingkaran{Titik{cx2, cy2}, r2}
	titik := Titik{px, py}

	dalamL1 := diDalam(lingkaran1, titik)
	dalamL2 := diDalam(lingkaran2, titik)

	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
