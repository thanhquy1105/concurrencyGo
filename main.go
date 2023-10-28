package main

import (
	diningphilosophers "github.com/thanhquy1105/concurrencyGo/diningphilosophers/content"
)

func main() {
	// goroutinewaitgroup.Run()
	// raceconditionsmutex.Run()
	// complexraceconditionsmutex.Run()
	// producerconsumer.Run()
	diningphilosophers.Run()
}
