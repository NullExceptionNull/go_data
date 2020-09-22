package main

import (
	array2 "data_structrue/02-array/query_and_update_element/array"
	"fmt"
)

func main() {
	a := array2.New(5)
	a.AddLast(1)
	a.AddLast(2)
	a.AddLast(3)
	a.AddLast(5)

	a.Add(3, 4)

	fmt.Println(a)

}
