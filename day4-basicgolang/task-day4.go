package main

import "fmt"

func RectnTri(x int) string {
	res := ""

	for i := 0; i < x; i++ {

		if x%2 == 0 {

			for j := 0; j <= x; j++ {
				res += "*"
			}

			res += "\n"
		}

		if x%2 == 1 {

			for j := 0; j <= i; j++ {
				res += "*"
			}

			res += "\n"
		}

	}
	return res
}

func main() {
	var angka int

	fmt.Print("Masukkan Angka : ")
	fmt.Scanln(&angka)

	fmt.Print(RectnTri(angka))
}
