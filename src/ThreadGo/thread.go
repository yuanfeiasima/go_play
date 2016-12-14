package main

import (
	"fmt"
	"runtime"
)

var counter int = 0

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func main() {
	//chs := make([]chan int, 10)
	//for i := 0; i < 10; i++ {
	//	chs[i] = make(chan int)
	//	go Count(chs[i])
	//}
	//
	//for _, ch := range chs {
	//	i := <-ch
	//	fmt.Println(i)
	//}
	//ThreadGo.GetCpuNUm()
	multiThread(2)
}

//设置GOMAXPROCS 充分利用多核的优势 真正的并发执行
func multiThread(cpuNum int) {
	runtime.GOMAXPROCS(cpuNum)
	for {
		fmt.Print(1)
		go fmt.Print(0)
	}
}
