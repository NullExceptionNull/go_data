package array

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
		size: 0,
	}
}

//获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

//获取数组元素的个数
func (a *Array) GetSize() int {
	return a.size
}

//返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

//在第Index 个位置插入一个新元素
func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("add failed,index out of range")
	}
	if a.size == len(a.data) {
		a.reSize(2 * len(a.data))
	}
	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}
	a.data[index] = e
	a.size++
}

func (a *Array) reSize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

func (a *Array) GetFirst() interface{} {
	return a.Get(0)
}

func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("get failed ,index out of range")
	}
	return a.data[index]
}

func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("set failed ,index out of range")
	}
	a.data[index] = e
}

func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			return true
		}
	}
	return false
}

func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			return i
		}
	}
	return -1
}

func (a *Array) FindAll(e interface{}) (indies []int) {
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			indies = append(indies, i)
		}
	}
	return indies
}

func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("remove failed, index out of range")
	}
	//找到需要被移出的元素
	e := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil
	//在缩小的的时候先缩小1/4 防止震荡
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.reSize(len(a.data) / 2)
	}
	return e
}

func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveElement(e interface{}) bool {
	index := a.Find(e)
	if index == -1 {
		return false
	}
	a.Remove(index)
	return true
}

func (a *Array) RemoveAllElement(e interface{}) bool {
	index := a.Find(e)
	if index == -1 {
		return false
	}
	for i := 0; i < a.size; i++ {
		if utils.Compare(a.data[i], e) == 0 {
			a.Remove(i)
		}
	}
	return true
}

// 重写 Array 的 string 方法
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
