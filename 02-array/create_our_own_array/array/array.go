package array

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
