package main

import (
	"fmt"
	"log"
)

type Calc struct {
}

func (c Calc) calcBMI(p *Person) (bmi float64, err error) {
	if p.Weight < 0 {
		err = fmt.Errorf("weight cannot be negative")
		return
	}
	if p.Weight == 0 {
		err = fmt.Errorf("weight cannot be zero")
		return
	}
	if p.Tall < 0 {
		err = fmt.Errorf("height cannot be negative")
		return
	}
	if p.Tall == 0 {
		err = fmt.Errorf("height cannot be zero")
		return
	}
	bmi = p.Weight / (p.Tall * p.Tall)
	return
}

func (c Calc) CalcFatRate(p *Person) error {
	bmi, err := c.calcBMI(p)
	if err != nil {
		log.Println("计算BMI出错: ", err)
		return err
	}
	sexWeight := 0
	if p.Sex == "男" {
		sexWeight = 1
	} else {
		sexWeight = 0
	}
	p.FatRate = (1.2*bmi + float64(p.Age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return nil
}
