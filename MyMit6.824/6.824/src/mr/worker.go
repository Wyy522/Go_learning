package mr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"
)
import "log"
import "net/rpc"
import "hash/fnv"

type ByKey []KeyValue

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

//
// Map functions return a slice of KeyValue.
//
type KeyValue struct {
	Key   string
	Value string
}

//
// use ihash(key) % NReduce to choose the reduce
// task number for each KeyValue emitted by Map.
//
func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}

//
// main/mrworker.go calls this function.
//
func Worker(mapf func(string, string) []KeyValue, reducef func(string, []string) string) {
	// declare an argument structure.
	for true {
		args := RequestArgs{}
		// declare a reply structure.
		reply := ResponseReply{}
		CallGetMapTask(&args, &reply)
		if reply.MapState == 1 && reply.Task != nil {
			filename := reply.Task.FileName
			file, err := os.Open(filename)
			if err != nil {
				log.Fatalf("cannot open %v", filename)
			}
			content, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatalf("cannot read %v", filename)
			}
			file.Close()
			kva := mapf(filename, string(content))
			numReduce := reply.Task.NumReduce
			taskId := reply.Task.TaskId
			bucket := make([][]KeyValue, numReduce)
			for _, kv := range kva {
				num := ihash(kv.Key) % numReduce
				bucket[num] = append(bucket[num], kv)
			}
			for i := 0; i < numReduce; i++ {
				tempFile, err := ioutil.TempFile("", "mr-map-*")
				if err != nil {
					log.Fatal("err :cannot open tmp_file")
				}
				enc := json.NewEncoder(tempFile)
				err = enc.Encode(bucket[i])
				if err != nil {
					log.Fatal("err :encode bucket error")
				}
				tempFile.Close()
				out_file := "mr-" + strconv.Itoa(taskId) + "-" + strconv.Itoa(i)
				os.Rename(tempFile.Name(), out_file)
				fmt.Printf("ok :creat map file  %v\n", out_file)
			}
			time.Sleep(time.Second)
			//CallTaskFin(&args, &reply)
		} else if reply.ReduceState == 1 {
			CallGetReduceTask(&args, &reply)
			//id := strconv.Itoa(reply.Task.ReduceId)
			num_map := 8
			reduceId := reply.ReduceId
			intermediate := []KeyValue{}
			for i := 0; i < num_map; i++ {
				map_filename := "mr-" + strconv.Itoa(i) + "-" + strconv.Itoa(reduceId)
				inputfile, err := os.OpenFile(map_filename, os.O_RDONLY, 0777)
				if err != nil {
					log.Fatalf("err :cannot open reduceTask %v", map_filename)
				}
				dec := json.NewDecoder(inputfile)
				for {
					var kv []KeyValue
					if err := dec.Decode(&kv); err != nil {
						break
					}
					intermediate = append(intermediate, kv...)
				}
			}
			sort.Sort(ByKey(intermediate))
			out_file := "mr-out-" + strconv.Itoa(reduceId)
			ofile, err := os.Create(out_file)
			//temp_file, err := ioutil.TempFile("", "mr-reduce-*")
			if err != nil {
				log.Fatal("err :cannot open tmp_file")
			}
			i := 0
			for i < len(intermediate) {
				j := i + 1
				for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
					j++
				}
				values := []string{}
				for k := i; k < j; k++ {
					values = append(values, intermediate[k].Value)
				}

				output := reducef(intermediate[i].Key, values)

				// this is the correct format for each line of Reduce output.
				//fmt.Fprintf(temp_file, "%v %v\n", intermediate[i].Key, output)
				fmt.Fprintf(ofile, "%v %v\n", intermediate[i].Key, output)

				i = j
			}
			//temp_file.Close()
			ofile.Close()
			//os.Rename(temp_file.Name(), out_file)
			fmt.Printf("ok :creat reduce file  %v\n", out_file)
		}
	}
}

//
// example function to show how to make an RPC call to the coordinator.
//
// the RPC argument and reply types are defined in rpc.go.
//
func CallGetMapTask(args *RequestArgs, reply *ResponseReply) {

	ok := call("Coordinator.GetMapTask", &args, &reply)
	if ok {
		// reply.Y should be 100.
		fmt.Printf("ok :get map task %v\n", reply.Task)
	} else {
		fmt.Printf("call failed!\n")
	}
}

func CallGetReduceTask(args *RequestArgs, reply *ResponseReply) {

	ok := call("Coordinator.GetReduceTask", &args, &reply)
	if ok {
		// reply.Y should be 100.
		fmt.Printf("ok :put reduce task %v\n", reply.Task)
	} else {
		fmt.Printf("call failed!\n")
	}
}

//func CallTaskFin(args *RequestArgs, reply *ResponseReply) {
//
//	ok := call("Coordinator.MakeReduceTask", &args, &reply)
//	if ok {
//		// reply.Y should be 100.
//		fmt.Printf("ok :put to reduce %v\n", reply.Task)
//	} else {
//		fmt.Printf("call failed!\n")
//	}
//}

//
// send an RPC request to the coordinator, wait for the response.
// usually returns true.
// returns false if something goes wrong.
//
func call(rpcname string, args interface{}, reply interface{}) bool {
	// c, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	sockname := coordinatorSock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer c.Close()

	err = c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
