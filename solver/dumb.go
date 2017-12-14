package main

type StringList []string
type IntList []int

func (los *StringList) cons(s string) {
	*los = append([]string{s}, *los...)
}

func (loi *IntList) cons(i int) {
	*loi = append([]int{i}, *loi...)
}

type AbstractList []interface{}

func (al *AbstractList) cons(elem interface{}) {
	*al = append([]interface{}{elem}, *al...)
}

func Cons(al interface{}, elem interface{}) {
	al = append([]interface{}{elem}, al)
}