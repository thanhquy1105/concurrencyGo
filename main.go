package main

import (
	sleepingbarber "github.com/thanhquy1105/concurrencyGo/sleepingbarber/content"
)

func main() {
	// goroutinewaitgroup.Run()
	// raceconditionsmutex.Run()
	// complexraceconditionsmutex.Run()
	// producerconsumer.Run()
	// diningphilosophers.Run()
	sleepingbarber.Run()
}
