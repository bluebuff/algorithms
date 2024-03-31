package array

type DynamicArray struct {
	size  int
	cap   int
	array interface{}
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		size:  0,
		cap:   8,
		array: [8]int{},
	}
}

func (d *DynamicArray) AddLast(e int) {
	d.array.([]int)[d.size] = e
}
