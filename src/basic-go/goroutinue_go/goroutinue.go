package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var mutex sync.Mutex
var rwMutex sync.RWMutex

func testgo() {
	fmt.Println("hello goroutinue")
	for i := 0; i < 10; i++ {
		fmt.Println("test goroutine: ", i)
	}
}

func testRuntime() {
	fmt.Println("GOROOT: ", runtime.GOROOT())
	fmt.Println("OS Version: ", runtime.GOOS)
	fmt.Println("logic cpu number: ", runtime.NumCPU())

	gomaxprocs := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(gomaxprocs)

	go func() {
		fmt.Println("go begin")
		runtime.Goexit() //强制结束当前函数(协程)
		fmt.Println("go end")
	}()

	time.Sleep(3000 * time.Millisecond)

}

func Relief1() {
	fmt.Println("relief--1~~~")
	waitGroup.Done()
}

func Relief2() {
	fmt.Println("relief--2~~~")
	waitGroup.Done()
}

func Relief3() {
	fmt.Println("relief--3~~~")
	waitGroup.Done()
}

func testChannel() {
	var channel chan int
	fmt.Printf("通道的数据类型:%T,通道的值:%v\n", channel, channel)
	if channel == nil {
		channel = make(chan int)
		fmt.Printf("通道的数据类型:%T,通道的值:%v\n", channel, channel)
	}

}
func main() {
	//go testgo()
	//for i := 0; i < 10; i++ {
	//	fmt.Println("main", i)
	//}
	////让main函数(main函数就相当于是一个goroutine(主协程))
	//// 并且主协程结束之后并不会等待子协程执行结束,而是直接结束整个程序

	//testRuntime()
	//time.Sleep(3000 * time.Millisecond)
	//fmt.Println("main end")
	//mutex.Lock()
	//mutex.Unlock()
	//mutex.TryLock()
	//waitGroup.Add(3) //var waitGroup sync.WaitGroup
	//go Relief1()
	//go Relief2()
	//go Relief3()
	//
	//waitGroup.Wait() //等待3个协程都执行完毕后主协程才能继续向下执行
	//fmt.Println("main end~~")
	//testChannel()
	ch := make(chan int)
	go func() {
		fmt.Println("子协程执行~~~~~")
		data := <-ch
		fmt.Println("从通道中读取到的数据是:", data)
	}()
	ch <- 28
	fmt.Println("main end")
}
