package main

import (
	//"time"
	"fmt"
	//"log"
	//"syn"
	//"time"
	"math/rand"
	"strconv"
	"time"
)

//声明成游戏
type Payload struct {
	name string
}

// 打游戏
func (p *Payload) Play() {
	fmt.Printf("%s 打LOL游戏...当前任务完成", p.name)
}

//任务
type Job struct {
	Payload Payload
}

//任务队列
var JobQueue chan Job

//工人
type Worker struct {
	name       string        //工人的名字
	WorkerPool chan chan  Job //对象池
	JobChannel chan Job      //通道里面的任务
	quit       chan bool     //是否退出
}

// 新建一个工人
func NewWorker(workerPool chan chan Job, name string) Worker {
	fmt.Printf("创建一个工人,他的名字是：%s", name)
	return Worker{
		name:       name,
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

//工人开始工作
func (w *Worker) Start() {
	go func() {

		for {
			//注册到对象池中
			w.WorkerPool <- w.JobChannel
			fmt.Printf("[%s]把自己注册到对象池中 \n", w.name)
			select {
			case job := <-w.JobChannel:
				fmt.Printf("[%s] 工人接收到了任务 当前任务长度是[%d] \n", w.name, len(w.WorkerPool))
				job.Payload.Play()
				time.Sleep(time.Duration(rand.Int31n(1000)) * time.Microsecond)
			//接收到任务
			case <-w.quit:
				return
			}
		}
	}()
}

//停止工人工作
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	name       string        //调度的名字
	maxWorkers int           //获取工人人数
	WorkerPool chan chan Job //注册和工人一样的通道
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,       //将工人放到一个池中，可以理解成一个部门中
		name:       "调度者",      //调度者的名字,
		maxWorkers: maxWorkers, //这个调度者有好多工人
	}
}

func (d *Dispatcher) Run() {

	//开始运行
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool, fmt.Sprintf("work-%s", strconv.Itoa(i)))
		//开始工作
		worker.Start()
	}
	//监控
	go d.dispatch()

}

func (d *Dispatcher) dispatch() {

	for {
		select {
		case job := <-JobQueue:
			fmt.Println("调度者,接收到一个工作任务")
			time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
		// 调度者接收到一个工作任务
			go func(job Job) {
				//从现有的对象池中拿出一个
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		default:
		//fmt.Println("ok!!")
		}

	}
}

func initialize() {

	maxWorkers := 2
	maxQueue := 4
	//初始化一个调试者,并指定它可以操作的 工人个数
	dispatch := NewDispatcher(maxWorkers)
	JobQueue = make(chan Job, maxQueue)
	//并让它一直运行
	dispatch.Run()
}
func main() {
	//初始化对象池
	initialize()
	for i := 0; i < 5; i++ {
		p := Payload{
			fmt.Sprintf("玩家-[%s]", strconv.Itoa(i)),
		}
		JobQueue <- Job{
			Payload: p,
		}
		time.Sleep(time.Second)
	}
	close(JobQueue)
}
