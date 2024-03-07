package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, " num:", num)
		}
	}
	fmt.Println(sum)

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for _, v := range m {
		fmt.Println(v)
	}
	for k := range m {
		fmt.Println(k)
	}
}
