package main

import (
	"fmt"
	"strings"
)

//Start ..
func (b *Broadcast) Start() {
	for i := 0; i < 3; i++ {
		b.m.Lock()
		fmt.Print(AlfaMsgs[i])
		b.Current(AlfaMsgs[i])
		b.Al.WhatMsg++
		b.m.Unlock()
	}
}

func (b *Broadcast) Current(msg string) {
	sender, receiver := "", ""

	if strings.HasPrefix(msg, "-Говорит") {
		sender, receiver = checkMsg(msg)
		b.CurrentSender = sender
		b.CurrentReceiver = receiver
	}

	if msg == "-Конец связи\n" {
		b.CurrentSender = "Альфа"
		b.Messages <- msg
		b.Enemy <- msg
		return
	}
	switch b.CurrentReceiver {
	case "Альфа":
		b.Al.Alfa <- msg
	case "Браво":
		b.Br.Bravo <- msg
	case "Чарли":
		b.Ch.Charlie <- msg
	}
	b.Enemy <- msg
}

func checkMsg(msg string) (string, string) {

	message := strings.Split(msg, " ")

	sender, receiver := message[1], message[3]

	return sender[:len(sender)-1], receiver[:len(receiver)-1]

}

//Chating ..
func (b *Broadcast) Chating() {

	go func() {
		for {
			select {
			case msg := <-b.Al.Alfa:
				if msg == "error" {
					fmt.Println("Cant create file Alfa.txt")
					return
				}
				if strings.HasSuffix(msg, "Прием\n") {
					b.m.Lock()
					for i := b.Al.WhatMsg; i < b.Al.Capacity; i++ {
						fmt.Print(AlfaMsgs[i])
						b.Current(AlfaMsgs[i])
						b.Al.WhatMsg++
						if strings.HasSuffix(AlfaMsgs[i], "Прием\n") {
							break
						}
					}
					b.m.Unlock()
				}
			case <-b.Messages:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case msg := <-b.Br.Bravo:
				if msg == "error" {
					fmt.Println("Cant create file Bravo.txt")
					return
				}
				if strings.HasSuffix(msg, "Прием\n") {
					b.m.Lock()
					for i := b.Br.WhatMsg; i < b.Br.Capacity; i++ {
						fmt.Print(BravoMsgs[i])
						b.Current(BravoMsgs[i])
						b.Br.WhatMsg++
						if strings.HasSuffix(BravoMsgs[i], "Прием\n") {
							break
						}
					}
					b.m.Unlock()
				}
			case <-b.Messages:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case msg := <-b.Ch.Charlie:
				if msg == "error" {
					fmt.Println("Cant create file Charlie.txt")
					return
				}
				if strings.HasSuffix(msg, "Прием\n") {
					b.m.Lock()
					for i := b.Ch.WhatMsg; i < b.Ch.Capacity; i++ {
						fmt.Print(CharlieMsgs[i])
						b.Current(CharlieMsgs[i])
						b.Ch.WhatMsg++
						if strings.HasSuffix(CharlieMsgs[i], "Прием\n") {
							break
						}
					}
					b.m.Unlock()
				}
			case <-b.Messages:
				return
			}
		}
	}()
}
