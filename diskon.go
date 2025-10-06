package main

import "fmt"

func main() {
	var totalBelanja int
	var isMember bool
	var diskon bool

	fmt.Print("Masukkan Total Belanja (Rp): ")
	fmt.Scanln(&totalBelanja)

	fmt.Print("Apakah Pelanggan Membership? (true/false): ")
	fmt.Scanln(&isMember)

	diskon = totalBelanja > 500000 || isMember

	fmt.Println("Pelanggan Dapat Diskon: ", diskon)
}