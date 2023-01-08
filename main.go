package main

import (
	"bitly-clone/model"
	"bitly-clone/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
