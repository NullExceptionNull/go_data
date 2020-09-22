package array

import (
	"bytes"
	"data_structrue/utils"
	"fmt"
	"strconv"
)

type Array struct {
	data []int
	size int
}

func New(capacity int) *Array {
	return &Array{
		data: make([]int, capacity),
		size: 0,
	}
}

func (a *Array) GetSize() int {
	return a.size
}
func (a *Array) GetCapacity() int {
	return len(a.data)
}
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//向所有元素的最后添加一个元素
//func (a *Array) AddLast(e int) {
//	if a.size == len(a.data) {
//		panic("add failed,02-array is full")
//	}
//	a.data[a.size] = e
//	a.size++
//}

func (a *Array) AddLast(e int) {
	a.Add(a.size, e)
}

func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}

//向数组中指定位置添加元素

func (a *Array) Add(index int, e int) {
	if a.size == len(a.data) {
		panic("add failed,02-array is full")
	}
	if index < 0 || index > a.size {
		panic("add failed, index out of range")
	}
	//从最后一个元素一直到index那一个 一直往后挪
	for i := a.size - 1; i >= index; i-- {
		//后一个元素的位置给前一个让位
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}

func (a *Array) Get(index int) int {
	if index < 0 || index >= a.size {
		panic("get failed, index out of range")
	}
	return a.data[index]
}
func (a *Array) Set(index int, e int) {
	if index < 0 || index >= a.size {
		panic("set failed, index out of range")
	}
	a.data[index] = e
}

func (a *Array) Contains(e int) bool {
	for _, datum := range a.data {
		if utils.Compare(datum, e) == 0 {
			return true
		}
	}
	return false
}

//查找数组中元素 e 所在的索引
func (a *Array) Find(e int) int {
	for i, datum := range a.data {
		if utils.Compare(datum, e) == 0 {
			return i
		}
	}
	return -1
}

func (a *Array) FindAll(e int) (indies []int) {
	for i, datum := range a.data {
		if utils.Compare(datum, e) == 0 {
			indies = append(indies, i)
		}
	}
	return indies
}

func (a *Array) Remove(index int) int {
	if index < 0 || index >= a.size {
		panic("remove failed, index out of range")
	}
	e := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	return e
}

func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

func (a *Array) RemoveLast() int {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveElement(e int) bool {
	index := a.Find(e)
	if index == -1 {
		return false
	}
	a.Remove(index)
	return true
}

func (a *Array) RemoveAllElement(e int) bool {
	index := a.Find(e)
	if index == -1 {
		return false
	}
	for i, datum := range a.data {
		if datum == e {
			a.Remove(i)
		}
	}
	return true
}

func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array :Size = %d, capaticy = %d \n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(strconv.Itoa(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()

}
