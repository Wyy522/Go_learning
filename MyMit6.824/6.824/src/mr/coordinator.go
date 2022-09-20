package mr

import (
	"fmt"
	"log"
	"sync"
)
import "net"
import "os"
import "net/rpc"
import "net/http"

var mu sync.Mutex

type Coordinator struct {
	// Your definitions here.
	DontWaitTaskChannel   bool
	DontWaitReduceChannel bool
	TaskChannelMap        chan *Task
	TaskChannelReduce     chan int
	NumMap                int
	NumReduce             int
	MapState              int //0 wait  1 start  2 finsh
	ReduceState           int //0 wait  1 start  2 finsh
}

type Task struct {
	FileName  string
	TaskType  int // 任务类型判断到底是map还是reduce 0 map 1 reduce
	TaskId    int // 任务的id
	NumReduce int // 传入的reducer的数量，用于hash
	ReduceId  int
}

func (c *Coordinator) GetMapTask(args *RequestArgs, reply *ResponseReply) error {
	// 分发任务应该上锁，防止多个worker竞争，并用defer回退解锁
	mu.Lock()
	defer mu.Unlock()
	if c.DontWaitTaskChannel {
		task := <-c.TaskChannelMap
		if len(c.TaskChannelMap)-1 < 0 {
			c.DontWaitTaskChannel = false
		}
		reply.MapState = 1
		reply.Task = task
		fmt.Println("ok : get map task ")
	} else {
		reply.ReduceState = 1
		fmt.Println("reduce !")
	}
	return nil
}

func (c *Coordinator) GetReduceTask(args *RequestArgs, reply *ResponseReply) error {
	mu.Lock()
	defer mu.Unlock()
	reply.ReduceId, _ = <-c.TaskChannelReduce
	if c.DontWaitReduceChannel {
		if len(c.TaskChannelReduce) >= 0 {
			fmt.Println("ok :make a reduce task ", reply.ReduceId)
		} else if len(c.TaskChannelReduce) == 0 {
			reply.ReduceState = 2
			c.DontWaitReduceChannel = false
			log.Fatal("err :no such ReduceChannel to make")
		} else {
			fmt.Println("errrrrrr")
		}
	}
	return nil
}

//
// start a thread that listens for RPCs from worker.go
//
func (c *Coordinator) server() {
	rpc.Register(c)
	rpc.HandleHTTP()
	//l, e := net.Listen("tcp", ":1234")
	sockname := coordinatorSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//
// main/mrcoordinator.go calls Done() periodically to find out
// if the entire job has finished.
//
func (c *Coordinator) Done() bool {
	ret := false
	// Your code here.
	return ret
}

//
// create a Coordinator.
// main/mrcoordinator.go calls this function.
// nReduce is the number of reduce tasks to use.
//
func MakeCoordinator(files []string, nReduce int) *Coordinator {
	c := Coordinator{
		DontWaitTaskChannel:   true,
		DontWaitReduceChannel: true,
		TaskChannelMap:        make(chan *Task, len(files)),
		TaskChannelReduce:     make(chan int, nReduce),
		NumMap:                len(files),
		NumReduce:             nReduce,
		MapState:              0,
		ReduceState:           0,
	}
	c.makeMapTask(files, nReduce)
	c.server()
	return &c
}

func (c *Coordinator) makeMapTask(files []string, nReduce int) {
	for i, file := range files {
		task := Task{
			FileName:  file,
			TaskType:  0,
			TaskId:    i,
			NumReduce: c.NumReduce,
		}
		fmt.Println("make a map task :", &task)
		c.TaskChannelMap <- &task
	}

	for i := 0; i < nReduce; i++ {
		c.TaskChannelReduce <- i
	}
}
