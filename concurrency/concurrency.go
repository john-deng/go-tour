package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)


func TestSimpleGoRoutine() {

	fmt.Println(">>> Test Simple Go Routine")

	c := make(chan bool)

	go func() {
		fmt.Println("Simple go, go, go ...")
		c <- true
	}()

	<- c

}

func Go(c chan bool, index int)  {
	a := 1
	for i := 0; i <= 1000000; i++ {
		a += i
	}

	fmt.Println(index, a)

	c <- true
}

func TestSimpleGoRoutineWithCache(numOfTasks int) {

	fmt.Println(">>> Test Go Routine with Cache")

	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, numOfTasks)
	for i := 0; i < numOfTasks; i++ {
		go Go(c, i)
	}

	for i := 0; i < numOfTasks; i++ {
		<- c
	}
}



func GoWg(wg *sync.WaitGroup, index int)  {
	a := 1
	for i := 0; i <= 1000000; i++ {
		a += i
	}

	fmt.Println(index, a)

	wg.Done()
}

func TestSimpleGoRoutineWithWaitGroup(numOfTasks int) {

	fmt.Println(">>> Test Go Routine with sync.WaitGroup")

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(numOfTasks)
	for i := 0; i < numOfTasks; i++ {
		go GoWg(&wg, i)
	}

	wg.Wait()
}


func TestGoRoutine() {

	fmt.Println(">>> Test Go Routine")

	c := make(chan bool)

	go func() {
		fmt.Println("Go, go, go ...")
		c <- true
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}

func TestGoRoutineWithMultiChannel()  {
	fmt.Println(">>> Test go routine with multiple channel")
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("c1")
					o <- true
					break
				}
				fmt.Println("c1", v)

			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2")
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 9
	c2 <- "go"
	c1 <- 8
	c2 <- "golang"

	close(c1)
	close(c2) // can be removed

	for i := 0; i < 2; i++ {
		<- o
	}
}

func Run()  {
	TestSimpleGoRoutine()

	TestGoRoutine()

	TestSimpleGoRoutineWithCache(15)

	TestSimpleGoRoutineWithWaitGroup(20)

	TestGoRoutineWithMultiChannel()
}