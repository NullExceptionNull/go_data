package main

import (
	array2 "data_structrue/03-stacks-and-queues/Array-stack/array"
	"fmt"
)

func main() {
	array := array2.New(10)

	for i := 0; i < 10; i++ {
		array.AddLast(i)
	}
	fmt.Println(array)
	array.AddLast("11")
	fmt.Println(array)
	array.RemoveLast()
	array.RemoveLast()
	array.RemoveLast()

	fmt.Println(array)

}
