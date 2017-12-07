//  https://play.golang.org/p/nai7XtTMfr

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func worker(wg *sync.WaitGroup, cs chan string, i int) {
	defer wg.Done()
	cs <- "worker" + strconv.Itoa(i)
}

func monitorWorker(wg *sync.WaitGroup, cs chan string) {
	wg.Wait()
	close(cs)
}
func main() {
	wg := &sync.WaitGroup{}
	cs := make(chan string)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(wg, cs, i)
	}

	go monitorWorker(wg, cs)

	for i := range cs {
		fmt.Println(i)

	}
}

/*

For one of my requirement I have to create N number of worker go routines,
which will be monitored by one monitoring routine.
monitoring routine has to end when all worker routines completes.
My code ending in deadlock, please help.


Your monitorWorker never dies. When all the workers finish, it continues to wait on cs.
This deadlocks because nothing else will ever send on cs and therefore wg will never reach 0.
A possible fix is to have the monitor close the channel when all workers finish. If the for loop is in main,
it will end the loop, return from main, and end the program.


*/



