package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(0))
	// 第一次提交时未仔细审题，没有看到
	// 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	var s string
	if x < 0 {
		s = strconv.Itoa(-x)
	} else {
		s = strconv.Itoa(x)
	}
	length := len(s)
	r := make([]rune, length)
	for i, v := range s {
		r[length-1-i] = v
	}
	var reversedX int64
	reversedX, _ = strconv.ParseInt(string(r), 10, 64)
	if x < 0 {
		reversedX = -reversedX
	}
	// 如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。
	limit := int64(1 << 31)
	if reversedX < -limit || reversedX >= limit {
		return 0
	}
	return int(reversedX)
}
