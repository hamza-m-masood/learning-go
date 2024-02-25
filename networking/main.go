package main

import (
	"HamzaMasood1/learning-go/networking/client"
	"HamzaMasood1/learning-go/networking/server"
	"time"
)

func main(){
	go server.Server()
  time.Sleep(time.Second)
	client.Get()
	client.Post()
	client.Delete()
}
