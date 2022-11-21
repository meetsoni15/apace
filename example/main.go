package main

import (
	"fmt"
	"log"

	"github.com/meetsoni15/apace"
)

func main() {
	fmt.Println(" ")
	fmt.Println(" //String generator with default character set  ")
	fmt.Println(" ")

	// String generator with default character set
	for i := 0; i < 5; i++ {
		str, err := apace.RandomString(15)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	}

	fmt.Println(" ")
	fmt.Println(" //String generator with custom character set  ")
	fmt.Println(" ")

	// String generator with custom character set
	for i := 0; i < 5; i++ {
		str, err := apace.RandomString(15, "ABCDEFGHIJKLMNOPQRSTVUWXYZ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	}

	fmt.Println(" ")
	fmt.Println(" //Int generator  ")
	fmt.Println(" ")

	// Int generator
	for i := 0; i < 5; i++ {
		num := apace.RandomInt(5, 100000)
		fmt.Println(num)
	}

	fmt.Println(" ")
	fmt.Println(" //Int Array generator  ")
	fmt.Println(" ")

	// Int Array generator
	for i := 0; i < 5; i++ {
		// Here we have provided argument
		// Length = 5 and digits in range upto 9999
		arr := apace.RandomArray(5, 9999)
		fmt.Println(arr)
	}

	fmt.Println(" ")
	fmt.Println(" //String Array generator  ")
	fmt.Println(" ")

	//String Array generator
	for i := 0; i < 5; i++ {
		// Here we have provided argument
		// Length = 5 and digits in range upto 9999
		arr := apace.RandomArray(5, "AbcdDMoPP")
		fmt.Println(arr)
	}

}
