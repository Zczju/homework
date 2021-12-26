package gobmi

import "fmt"

func BMI(heightM, weightKG float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("weight cannot be negative")
		return
	}
	if heightM < 0 {
		err = fmt.Errorf("height cannot be negative")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("height cannot be 0")
		return
	}
	if weightKG == 0 {
		err = fmt.Errorf("weight cannot be 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}
