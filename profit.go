package main

import "fmt"

func main() {
	var modalAwal float64
	var profit float64
	var jumlahProfit float64
	var hargaJual float64

	fmt.Print("Masukkan Modal Awal Barang: ")
	fmt.Scanln(&modalAwal)

	fmt.Print("Masukkan Persentase Profit yang Diinginkan: ")
	fmt.Scanln(&profit)

	jumlahProfit = modalAwal * (profit / 100)

	hargaJual = modalAwal + jumlahProfit

	fmt.Printf("Harga Jual Barang: Rp %.2f\n", hargaJual)
}
