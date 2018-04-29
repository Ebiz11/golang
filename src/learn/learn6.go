package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "mangga"
	names[1] = "jambu"
	names[2] = "anggur"
	names[3] = "jeruk"
	
	fmt.Println(names[0], names[1], names[2], names[3])

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
}