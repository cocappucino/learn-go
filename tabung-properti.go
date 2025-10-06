package main

import "fmt"

const PI float64 = 3.14

type cylinder struct {
	height  int
	radius  int
	area    float64
	volume  float64
	baseArea float64
	wallArea float64
}

func main() {
	var t cylinder

	fmt.Print("Masukan radius: ")
	fmt.Scanln(&t.radius)

	fmt.Print("Masukan tinggi: ")
	fmt.Scanln(&t.height)
	
	if t.radius == 0 || t.height == 0 {
		fmt.Println("\nUsing default values: Radius=7, Height=10")
		t.radius = 7
		t.height = 10
	}
	
	fmt.Printf("\nRadius: %d, Height: %d\n", t.radius, t.height)

	t.baseArea = PI * float64(t.radius) * float64(t.radius)

	t.wallArea = float64(t.height) * (2 * PI * float64(t.radius))

	t.area = 2*t.baseArea + t.wallArea

	t.volume = t.baseArea * float64(t.height)

	fmt.Println("--- Hasilnya ---")
	fmt.Printf("Luas: %.2f\n", t.area)
	fmt.Printf("Volume: %.2f\n", t.volume)
}