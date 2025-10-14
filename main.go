package main

import "fmt"

func main() {
 tes := "tess"

 if tes == "tess" {
	fmt.Printf("benar \n")
 } else {
	fmt.Printf("salah")
 }

 fmt.Printf("%s \n", tes)

 var xs = "12345" // string
for i, v := range xs {
    fmt.Println("Index=", i, "Value=", v)
}

var ys = [5]int{10, 20, 30, 40, 50} // array
for _, v := range ys {
    fmt.Println("Value=", v)
}
}