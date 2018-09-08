package main

import "fmt"

func main() {
	//prints decimal value
	fmt.Println(42)
	//prints binary value
	fmt.Printf("%b \n", 42)
	//prints hex value
	fmt.Printf("%x \n", 42)
	//prints all 3 values
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	//same as above but spacing and capitalised
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
	//prints numbers 1 - 100 in the three systems
	for i := 0; i < 100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
	//prints first 100 unicode characters
	for u := 0; u < 100; u++ {
		fmt.Printf("%q ", u)
	}
}
