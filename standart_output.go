package main

import (
	"fmt"
)

func main() {
	i := 21
	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", i)

	fmt.Printf("%% \n")

	j := true
	fmt.Printf("%t \n", j)
	fmt.Printf("%t \n", !j)

	fmt.Printf("%c atau %c  \n", 'Я', 0x042F)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%X \n", 15)
	fmt.Printf("%U \n", 'Я')

	k := 123.456
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
}