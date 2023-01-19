package main

import "github.com/dupreehkuda/gin-test-task/internal/server"

func main() {
	srv := server.NewByConfig()
	srv.Run()
}
