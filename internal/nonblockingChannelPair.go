package internal

// returns a write, read channel pair. The write channel accepts non-blocking writes.
func MakeNonblockingChanPair[T any]() (chan<- T, <-chan T) {
	write := make(chan T)
	read := make(chan T)

	go func() {
		var inQueue []T

		readChan := func() chan T {
			if len(inQueue) == 0 {
				return nil
			}
			return read
		}

		readVal := func() T {
			if len(inQueue) == 0 {
				var zero T
				return zero
			}
			return inQueue[0]
		}

		for len(inQueue) > 0 || write != nil {
			select {
			case v, ok := <-write:
				if !ok { // channel is closed
					write = nil
				} else {
					inQueue = append(inQueue, v)
				}
			// Block write with a nil channel when inQueue is empty
			case readChan() <- readVal():
				inQueue = inQueue[1:]
			}
		}
		close(read)
	}()

	return write, read
}
