package main

import "time"

func main() {
	//ch := make(chan bool)
	//go func() {
	//	defer close(ch) //do some thing heavy
	//	ch <- true      //sendresult
	//}()
	////waituntiltimeout
	//select {
	//case r := <-ch:
	//	println("Theresultis:", r)
	//case <-time.After(time.Second):
	//}
	solution2()
}
func solution2() {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do something heavy
		select {
		case ch <- true: // send result
		default:
			// do nothing if no one's expecting it
		}
	}()
	// wait until timeout
	select {
	case r := <-ch:
		println("The result is:", r)
	case <-time.After(time.Second):
	}
}