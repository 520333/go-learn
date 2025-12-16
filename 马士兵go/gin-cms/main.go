package main

import (
	"ginCms/handlers"
)

func main() {
	r := handlers.InitEngine()
	r.Run()
}
