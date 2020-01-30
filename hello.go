package main

import "fmt"

func main() {
	i := 10
	if i > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is less than 5")
	}

	a := 3
	for i := 0; i < 10; i++ {
		if i == a {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
}
