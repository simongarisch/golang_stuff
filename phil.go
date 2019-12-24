package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	name            string
	leftCS, rightCS *ChopStick
}

type Host struct {
	ch chan *Philosopher
}

func NewHost() Host {
	// at most two philosophers can be hosted
	return Host{ch: make(chan *Philosopher, 2)}
}

func (h *Host) PleaseCanIEat(p *Philosopher) {
	h.ch <- p
}

func (p *Philosopher) eat(host *Host, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		host.ch <- p

		// randomly choose the chopsticks to eat with
		sticks := []*ChopStick{p.leftCS, p.rightCS}
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(sticks), func(i, j int) {
			sticks[i], sticks[j] = sticks[j], sticks[i]
		})

		sticks[0].Lock()
		sticks[1].Lock()
		fmt.Println("starting to eat " + p.name)
		sticks[0].Unlock()
		sticks[1].Unlock()
		<-host.ch // move one philosopher from the channel
	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopStick)
	}

	host := NewHost()

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			name:    "p" + strconv.Itoa(i+1),
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%5],
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat(&host, &wg)
	}

	wg.Wait()
}
