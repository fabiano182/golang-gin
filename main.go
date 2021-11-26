package main

import "golang-gin/server"

func main() {
	s := server.NewServer()

	s.Run()
}
