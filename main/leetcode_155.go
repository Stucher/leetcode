package main

type MinStack struct {
	stackArr []int
	minArr   []int
	size     int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stackArr: make([]int, 0),
		minArr:   make([]int, 0),
		size:     0,
	}
}

func (this *MinStack) Push(x int) {
	// 添加元素 追加或者修改
	if len(this.stackArr) == this.size {
		this.stackArr = append(this.stackArr, x)
	} else {
		this.stackArr[this.size] = x
	}

	// 更新最小值的数组
	if this.size == 0 && len(this.minArr) == 0 {
		this.minArr = append(this.minArr, x)
	} else if this.size == 0 {
		this.minArr[this.size] = x
	} else if this.size == len(this.minArr) {
		this.minArr = append(this.minArr, minInt(x, this.minArr[this.size-1]))
	} else {
		this.minArr[this.size] = minInt(x, this.minArr[this.size-1])
	}

	this.size++
}

func (this *MinStack) Pop() {
	if this.size == 0 {
		return
	}
	this.size--
}

func (this *MinStack) Top() int {
	if this.size == 0 {
		return 0
	}

	return this.stackArr[this.size-1]
}

func (this *MinStack) GetMin() int {
	if this.size == 0 {
		return 0
	}

	return this.minArr[this.size-1]
}

func minInt(intArr ...int) int {
	minValue := intArr[0]
	for i := 1; i < len(intArr); i++ {
		if minValue < intArr[i] {
			continue
		}
		minValue = intArr[i]
	}
	return minValue
}
