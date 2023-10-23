package main

import (
	complexraceconditionsmutex "github.com/thanhquy1105/concurrencyGo/complexraceconditionsmutex/content"
	goroutinewaitgroup "github.com/thanhquy1105/concurrencyGo/goroutinewaitgroup/content"
	raceconditionsmutex "github.com/thanhquy1105/concurrencyGo/raceconditionsmutex/content"
)

func main() {
	goroutinewaitgroup.Run()
	raceconditionsmutex.Run()
	complexraceconditionsmutex.Run()
}
