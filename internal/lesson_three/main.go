package main

import "fmt"

func main() {
	var c map[string]*int = nil
	fmt.Printf("%x\n", c["1"])

	var ds []int = nil
	fmt.Printf("%d, %d\n", len(ds), cap(ds))
	ds = append(ds, 1)
	fmt.Printf("%d, %d\n", len(ds), cap(ds))

}
