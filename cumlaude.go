package main

import "fmt"

func main() {
	var semester, eprt int
	var isCumlaude bool

	fmt.Print("Berapa semester? : ")
	fmt.Scanln(&semester) 

	fmt.Print("Berapa skor eprt? : ")
	fmt.Scanln(&eprt)


	isCumlaude = semester <= 8 && eprt >= 500

	fmt.Println(isCumlaude)
}