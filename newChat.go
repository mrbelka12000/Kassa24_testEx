package main

import (
	"sync"
)

//CreateBreadcast ..
func CreateBreadcast() *Broadcast {
	return &Broadcast{
		Al: &AlfaChat{
			Capacity: 7,
			Alfa:     make(chan string),
		},
		Br: &BravoChat{
			Capacity: 4,
			Bravo:    make(chan string),
		},
		Ch: &CharlieChat{
			Capacity: 3,
			Charlie:  make(chan string),
		},
		Messages: make(chan string),
		Enemy:    make(chan string),
		m:        &sync.Mutex{},
	}
}
