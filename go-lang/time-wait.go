package main

import (
	"fmt"
	"time"
)

func WaitCondition(timeout time.Duration, cond func() bool) error {
	ok := make(chan struct{})
	exit := make(chan struct{})
	go func() {
		for {
			select {
			case <-exit:
			default:
				if cond() {
					close(ok)
					return
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()
	select {
	case <-ok:
		return nil
	case <-time.After(timeout):
		close(exit)
		return fmt.Errorf("timed out waiting for condition after %s", timeout)
	}
}

func main() {
       var a = 0;
	err := WaitCondition(5*time.Second, func() bool {
               a = a+1
               fmt.Println(a)
		return a == 4
	})

fmt.Println(err)

}
