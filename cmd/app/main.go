package main

import (
	"fmt"
	"task-balancing/internal/config"
)

func main() {
	fmt.Println("app started")
	cfg := config.MustLoad()
	fmt.Println(cfg.Env)
	fmt.Println(cfg.HTTPServer.Address)
}
