package internal

import (
	"fmt"
	"sync"
	"testing"
)

func TestMakeNonblockingChanPair(t *testing.T) {
	write, read := MakeNonblockingChanPair[int]()
	last := -1
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for v := range read {
			fmt.Printf("%d <- read\n", v)
			if last+1 != v {
				t.Errorf("Unexpected value received. Want %d, got %d", last+1, v)
			}
			last = v
		}
		wg.Done()
	}()

	for i := 0; i < 100; i++ {
		write <- i
		fmt.Printf("write <- %d\n", i)
	}
	close(write)
	wg.Wait()

	if last != 99 {
		t.Errorf("Last value incorrect. Want 99, got %d", last)
	}
}
