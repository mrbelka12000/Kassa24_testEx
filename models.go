package main

import (
	"sync"
)

//Broadcast ..
type Broadcast struct {
	Al              *AlfaChat
	Br              *BravoChat
	Ch              *CharlieChat
	T               Txt
	CurrentSender   string
	CurrentReceiver string
	Messages        chan string
	Enemy           chan string
	m               *sync.Mutex
}

type Txt struct {
	Alfa    []string
	Bravo   []string
	Charlie []string
}

//AlfaChat ..
type AlfaChat struct {
	WhatMsg  int
	Capacity int
	Alfa     chan string
}

//BravoChat ..
type BravoChat struct {
	WhatMsg  int
	Capacity int
	Bravo    chan string
}

//CharlieChat ..
type CharlieChat struct {
	WhatMsg  int
	Capacity int
	Charlie  chan string
}
