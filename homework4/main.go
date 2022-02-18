package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := &FatRateRank{}
	for i := 0; i < 999; i++ {
		go func() {
			p := &Person{
				name: fmt.Sprint(i),
			}

			p.getBase()

			p.frChange()
			r.inputRecord(p.name, p.ChangedFR)
			rank, _ := r.getRank(p.name)
			fmt.Println(p.name, p.ChangedFR, rank)
		}()
	}
	time.Sleep(1 * time.Second)
}

type Person struct {
	name      string
	FRBase    float64
	ChangedFR float64
}

func (p *Person) getBase() {
	p.FRBase = rand.Float64() * 0.4
}

func (p *Person) frChange() {
	min := -0.2
	max := 0.2
	change := rand.Float64()*(max-min) + min
	if p.FRBase+change < 0 {
		p.ChangedFR = p.FRBase - change
	}
	p.ChangedFR = p.FRBase + change
}
