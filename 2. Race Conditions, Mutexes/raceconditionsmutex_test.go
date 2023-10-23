package raceconditionsmutex

import "testing"

// go test -race .
// go test .

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("x!")
	go updateMessage("Goodbye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Errorf("incorrect value in msg")
	}
}
