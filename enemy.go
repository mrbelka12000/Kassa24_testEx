package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//CollectMessages ..
func (b *Broadcast) CollectMessages() {
	for {
		msg := <-b.Enemy

		customFormat := "2006-01-02 15:04:05"
		tTime := time.Now()
		now := tTime.Format(customFormat)

		msgToFile := fmt.Sprintf("%v : %v", now, msg)
		// b.m.Lock()
		switch b.CurrentSender {
		case "Альфа":
			b.T.Alfa = append(b.T.Alfa, msgToFile)
		case "Браво":
			b.T.Bravo = append(b.T.Bravo, msgToFile)
		case "Чарли":
			b.T.Charlie = append(b.T.Charlie, msgToFile)
		}
		// b.m.Unlock()
	}
}

//WriteFile ..
func (b *Broadcast) WriteFile(name string) {
	file, err := os.Create(name + ".txt")
	if err != nil {
		log.Fatalln(err)
	}
	switch name {
	case "Альфа":
		for _, v := range b.T.Alfa {
			file.WriteString(v)
		}
	case "Браво":
		for _, v := range b.T.Bravo {
			file.WriteString(v)
		}
	case "Чарли":
		for _, v := range b.T.Charlie {
			file.WriteString(v)
		}
	}
}
