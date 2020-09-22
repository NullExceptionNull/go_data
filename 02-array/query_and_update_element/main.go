package main

import (
	array2 "data_structrue/02-array/query_and_update_element/array"
	"fmt"
)

func main() {
	array := array2.New(20)

	for i := 0; i < 10; i++ {
		array.AddLast(i)
	}

	array.Add(1, 100)
	fmt.Println(array.String())

	array.AddFirst(-1)
	fmt.Println(array.String())
}
