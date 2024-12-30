package main // 声明 main 包，表明当前是一个可执行程序

import (
	"fmt"

	_ "github.com/test/init_project/pkg1"
) // 导入内置 fmt 包

const mainName string = "main"

var mainVar string = getMainVar()

func init() {
	fmt.Println("main init method invoked")
}

func main() {
	fmt.Println("main method invoked!")
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}

// func main() { // main函数，是程序执行的入口
// 	fmt.Println("VSCode Let's GO!") // 在终端打印
// 	var s string = "Go语言"
// 	var bytes []byte = []byte(s)
// 	var runes []rune = []rune(s)
// 	fmt.Println("String length:", len(s))
// 	fmt.Println("bytes length：", len(bytes))
// 	fmt.Println("rune length:", len(runes))
// 	fmt.Println("string sub", s[0:7])
// 	fmt.Println("bytes sub", string(bytes[0:7]))
// 	fmt.Println("rune sub", string(runes[0:3]))

// 	var p1 *int
// 	var p2 *string
// 	i := 1
// 	s2 := "Hello"
// 	p1 = &i
// 	p2 = &s2
// 	p3 := &p2
// 	fmt.Println(p1)
// 	fmt.Println(p2)
// 	fmt.Println(p3)

// }

// func variables() {
// 	var s1 string = "Hello"
// 	var zero int
// 	var b1 = true

// 	var (
// 		i  int = 123
// 		b2 bool
// 		s2 = "test"
// 	)
// }

// func method1() {
// 	a := 1
// 	var b int = 2
// 	var c int
// 	fmt.Println(a, b, c)
// }

// func method2() (a int, b string) {
// 	return 1, "123"
// }

// func method3() (a int, b string) {
// 	a = 1
// 	b = "test"
// 	return
// }
// func method4() (a int, b string) {
// 	return
// }
