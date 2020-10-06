package main

import (
	"data_structrue/02-array/dynamic_array/array"
	"fmt"
)

func main() {
	a := array.New(2)

	for i := 0; i < 2; i++ {
		a.Add(i, i)
	}

	fmt.Println(a.String())

	a.Remove(1)

	fmt.Println(a.String())

}
