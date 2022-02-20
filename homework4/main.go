package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	r := &FatRateRank{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			p := &Person{
				name:        fmt.Sprint(i),
				BaseFatRate: rand.Float64() * 0.4,
			}

			// 无限循环（不停地去更新自己的体脂信息）
			for {
				if err := p.changeFatRate(); err == nil {
					rank := r.updateRecord(p.name, p.CurrentFatRate)
					fmt.Println(p.name, p.CurrentFatRate, rank)
				} else {
					fmt.Println(p.name, err)
				}
				time.Sleep(time.Second)
			}
		}()
	}

	wg.Wait()
}

type Person struct {
	name string

	BaseFatRate    float64
	CurrentFatRate float64
}

func (p *Person) changeFatRate() error {
	min := -0.2
	max := 0.2
	delta := rand.Float64()*(max-min) + min
	result := p.BaseFatRate + delta

	if result < 0 || result > 0.4 {
		// 超出范围返回错误
		return fmt.Errorf("invalid fat rate: %f", result)
	} else {
		p.CurrentFatRate = result
		return nil
	}
}
