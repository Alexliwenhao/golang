package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Male: %v\n", Male)
	test := string(Male)
	fmt.Printf("test: %v\n", test)
}

type ConnState int

const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateHijacked
	StateClosed
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

func (g *Gender) string() string {
	switch *g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

func (g *Gender) isMale() bool {
	return *g == Male
}
