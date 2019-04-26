package main

import (
	"fmt"
	"sync"
	"time"

	L "./Lotto649"
	E "./LottoExample"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	messages := make(chan string)

	var wg sync.WaitGroup

	// This is only here to play with execution times.
	for i := 0; i < 30; i++ {
		// add one to the sync group, to mark we should wait for one more
		wg.Add(2)
		go L.Lotto649(messages, &wg)
		go E.LottoExample(messages, &wg)
	}
	go func() {
		wg.Wait()
		close(messages)
	}()
	for i := range messages {
		fmt.Println(i)
	}

}


// Home System ~100 - 120 ms
