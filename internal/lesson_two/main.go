package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var ap *int = &a

	fmt.Printf("addr of `a`: %x\n", &a)
	fmt.Printf("`ap` value: %x\n", ap)
	fmt.Printf("value of `*ap`: %d\n", *ap)

	var b = [...]int{1, 2, 3}
	var bp = &b[1]

	fmt.Printf("%x\n", bp)
	fmt.Printf("%x\n", *&bp)

}
