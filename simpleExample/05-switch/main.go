package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("before noon")
	default:
		fmt.Println("after noon")
	}
}
