package main

import (
	"errors"
	"fmt"
)

var a []int

func f(b []int) []int {
	a = b[:2]
	return a
}

func main() {
	// b := []int{1, 2, 3, 4}
	// c := f(b)
	// fmt.Println(".....c......", c, len(c), cap(c))
	// fmt.Println(".....a.....", a, len(a), cap(a))
	// fmt.Println(".....b.....", b, len(b), cap(b))
	// b = []int{1, 2, 3, 4, 5}
	// fmt.Println(".....c......", c, len(c), cap(c))
	// fmt.Println(".....a.....", a, len(a), cap(a))
	// fmt.Println(".....b.....", b, len(b), cap(b))

	var e = cc()
	a := fmt.Sprintf("%+v", e)
	b := fmt.Sprintf("%v", e)
	fmt.Println(a)
	fmt.Println(b)

}

func cc() error {
	return errors.New("xxx")
}
