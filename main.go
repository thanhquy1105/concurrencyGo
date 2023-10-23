package main

import (
	complexraceconditionsmutex "github.com/thanhquy1105/concurrencyGo/complexraceconditionsmutex"
	goroutinewaitgroup "github.com/thanhquy1105/concurrencyGo/goroutinewaitgroup"
	raceconditionsmutex "github.com/thanhquy1105/concurrencyGo/raceconditionsmutex"
)

func main() {
	goroutinewaitgroup.Run()
	raceconditionsmutex.Run()
	complexraceconditionsmutex.Run()
}
