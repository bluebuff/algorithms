package array

import "errors"

type DynamicArray struct {
	array []int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		array: make([]int, 0, 10),
	}
}

func (d *DynamicArray) AddLast(e int) {
	d.array = append(d.array, e)
}

func (d *DynamicArray) Add(index int, e int) error {
	if index < 0 || index > len(d.array) {
		return errors.New("index param error")
	}
	d.array = append(d.array, 0)
	copy(d.array[index+1:], d.array[index:])
	d.array[index] = e
	return nil
}

func (d *DynamicArray) Remove(index int) error {
	if index < 0 || index > len(d.array)-1 {
		return errors.New("index param error")
	}
	copy(d.array[index:], d.array[index+1:])
	// d.array[len(d.array)-1] = 0
	d.array = d.array[:len(d.array)-1]
	return nil
}

func (d *DynamicArray) Each(f func(v int)) {
	for _, v := range d.array {
		f(v)
	}
}
