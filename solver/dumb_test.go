package main

import (
	"fmt"
	"testing"
)

func TestRegList(t *testing.T) {
	loi := IntList{1, 2, 3, 4, 5}
	fmt.Println(loi)

	loi.cons(1000)
	fmt.Println(loi)

	Cons(loi, 52)
	fmt.Println(loi)
}
