package raceconditionsmutex

import (
	"sync"
	"testing"
)

// go test -race .
// go test .

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"
	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Goodbye, cruel world!", &mutex)
	go updateMessage("x!", &mutex)
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Errorf("incorrect value in msg")
	}
}
