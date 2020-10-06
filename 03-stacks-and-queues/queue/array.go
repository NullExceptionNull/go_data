package queue

import (
	"bytes"
	"data_structrue/utils"
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}

func New(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
	}
}

func (a *Array) GetCapacity() int {
	return len(a.data)
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) IsEmpty() bool {
	return len(a.data) == 0
}

func (a *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("add failed, index out of range")
	}
	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}

func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("get failed, index out of range")
	}
	return a.data[index]
}

func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("get failed, index out of range")
	}
	a.data[index] = e
}

func (a *Array) Contains(e interface{}) bool {
	for _, datum := range a.data {
		if utils.Compare(datum, e) == 0 {
			return true
		}
	}
	return false
}

//查找数组中元素 e 所在的索引，不存在则返回 -1
func (a *Array) Find(e interface{}) int {
	for index, datum := range a.data {
		if utils.Compare(datum, e) == 0 {
			return index
		}
	}
	return -1
}

// 查找数组中元素 e 所有的索引组成的切片，不存在则返回 -1
func (a *Array) FindAll(e interface{}) (indexes []int) {
	for i, datum := range a.data {
		if utils.Compare(e, datum) == 0 {
			indexes = append(indexes, i)
		}
	}
	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("remove failed, index out of range")
	}
	e := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil
	if a.size == len(a.data)/2 {
		a.resize(len(a.data) / 2)
	}
	return e
}

// 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

// 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 从数组中删除一个元素 e
func (a *Array) RemoveElement(e interface{}) bool {
	index := a.Find(e)
	if index == -1 {
		return false
	}
	a.Remove(index)
	return true
}

func (a *Array) GetFirst() interface{} {
	return a.Get(0)
}

// 从数组中删除所有元素 e
func (a *Array) RemoveAllElement(e interface{}) bool {
	if a.Find(e) == -1 {
		return false
	}

	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			a.Remove(i)
		}
	}
	return true
}

func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
