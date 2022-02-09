package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	broadcast := CreateBreadcast()
	go broadcast.CollectMessages()
	go broadcast.Chating()
	broadcast.Start()

	<-broadcast.Messages
	broadcast.WriteFile("Альфа")
	broadcast.WriteFile("Чарли")
	broadcast.WriteFile("Браво")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done

	broadcast.m.Lock()
	close(broadcast.Messages)
	close(broadcast.Al.Alfa)
	close(broadcast.Br.Bravo)
	close(broadcast.Ch.Charlie)
	close(broadcast.Enemy)
	broadcast.m.Unlock()

	os.Remove("Альфа.txt")
	os.Remove("Браво.txt")
	os.Remove("Чарли.txt")
	log.Println("Broadcast stopped")
}
