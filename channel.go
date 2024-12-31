package main

import (
	"fmt"
	"time"
)

// 接收数据
func receiveObly(ch <-chan int) {
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
}

// 发送数据
func sendObly(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	close(ch)
}

func main() {
	fmt.Println("hello world")
	// 创建一个带缓冲的channel
	ch := make(chan int, 3)
	go sendObly(ch)
	go receiveObly(ch)
	// 使用select进行多路复用
	timeout := time.After(time.Second * 2)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("Channel已关闭")
				return
			}
			fmt.Printf("主goroutine接收到: %d\n", v)
		case <-timeout:
			fmt.Println("操作超时")
			return
		default:
			fmt.Println("没有数据，等待中...")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
