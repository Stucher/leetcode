package main

// 数字反转
var (
	minNum int64 = -(1 << 31)
	maxNum int64 = 1<<31 - 1
)

func reverse(x int) int {
	var isMinus bool = x < 0;
	if (isMinus) {
		x = -x;
	}
	var ret int64 = 0;
	for x > 0 {
		ret = ret*10 + int64(x%10);
		x /= 10;
	}
	if (isMinus) {
		ret = -ret;
	}
	if (ret < minNum || ret > maxNum) {
		ret = 0;
	}
	return int(ret);
}
