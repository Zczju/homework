package gobmi

import "fmt"

func CalcFatRate(bmi float64, sex string, age int) (fatRate float64, err error) {
	sexWeight := 0
	if sex == "男" {
		sexWeight = 1
	} else if sex == "女" {
		sexWeight = 0
	} else {
		err = fmt.Errorf("Warning: Illegal sex input. ")
		return
	}

	if bmi <= 0 {
		err = fmt.Errorf("Warning: Illegal bmi input. ")
		return
	}

	if age <= 0 || age >= 150 {
		err = fmt.Errorf("Warning: Illegal age input . ")
		return
	}

	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100

	if fatRate <= 0 {
		err = fmt.Errorf("Warning: Illegal fatRate output . ")
		return
	}

	return
}
