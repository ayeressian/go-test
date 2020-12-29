package test2

import "fmt"

func Test2() {
	ch := make(chan int)
	go func() {
		// close(ch)
		for i := 0; i < 10; i++ {
			ch <- i	
		}
		close(ch)
	}()

	for val := range ch{
		fmt.Printf("%v", val)
	}
}


func foo(c <- chan int) {
		
}