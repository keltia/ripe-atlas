/*
This package is just a collection of test cases
 */
package main

import (
    "fmt"
    "os"
    "ripe-atlas"
)

func main() {
	p, err := atlas.GetProbe(14037)
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("p: %#v\n", p)

	q, err := atlas.GetProbes()
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("q: %#v\n", q)

}
