package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	b := [3]int{1, 2, 3}
	var c [3]int
	d := [...]int{2: 1, 5: 5, 7: 13}
	c[0] = 10
	c[1] = 11

	fmt.Printf("a:%v\n", a)
	fmt.Printf("b:%v\n", b)
	fmt.Printf("c:%v\n", c)
	fmt.Printf("d:%v\n", d)
}
