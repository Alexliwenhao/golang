package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var p1 *int
	i := 1
	p1 = &i
	fmt.Println(*p1 == i)
	fmt.Println(i)
	*p1 = 2
	fmt.Println(i)
	fmt.Println(*p1 == i)
	modify_pointer_value()
	unsafe_Pointer()
}

func modify_pointer_value() {
	a := 2
	var p *int
	p = &a
	fmt.Println(p, &a)

	var pp **int
	pp = &p
	fmt.Println(pp, &p)
	**pp = 3
	fmt.Println(pp, *pp, p)
	fmt.Println(**pp, *p)
	fmt.Println(a, &a)
}

func unsafe_Pointer() {
	a := "Hello, world!"
	var upA uintptr
	upA = uintptr(unsafe.Pointer(&a))
	upA += 1
	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)
}
