package ThreadGo

import "fmt"

//管道流
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

//单向channel
var ch1 chan int       //能都能写
var ch2 chan<- float64 //只写
var ch3 <-chan int     //只读

//初始化
//ch4 := make(chan int)
//ch5 := <-chan int(ch4)
//ch6 := chan<- int(ch4)

func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parse value", value)
	}
}
