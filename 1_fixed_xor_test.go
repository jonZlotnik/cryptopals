package main

import "fmt"

func ExampleFixedXOR() {
	a := "1c0111001f010100061a024b53535009181c"
	b := "686974207468652062756c6c277320657965"
	fmt.Println(FixedXOR(a, b))
	// Output:
	// 746865206b696420646f6e277420706c6179
}
