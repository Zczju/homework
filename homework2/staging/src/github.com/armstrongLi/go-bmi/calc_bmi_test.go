package gobmi

import "testing"

func TestBMI(t *testing.T) {
	{
		inputHeight, inputWeight := 1.5, 72.0
		expectingOutput := 32.0
		actualOutput, _ := BMI(inputHeight, inputWeight)
		if expectingOutput != actualOutput {
			t.Fatalf("期望bmi为32.0，结果却为 %f ", actualOutput)
		}
	}
	{
		inputHeight, inputWeight := 0.0, 81.0
		_, err := BMI(inputHeight, inputWeight)
		if err == nil {
			t.Fatalf("期望出现身高为零错误，结果却没有")
		}
	}
	{
		inputHeight, inputWeight := 1.81, 0.0
		_, err := BMI(inputHeight, inputWeight)
		if err == nil {
			t.Fatalf("期望出现体重为零错误，结果却没有")
		}
	}
	{
		inputHeight, inputWeight := -1.0, 81.0
		_, err := BMI(inputHeight, inputWeight)
		if err == nil {
			t.Fatalf("期望出现身高为负错误，结果却没有")
		}
	}
	{
		inputHeight, inputWeight := 1.81, -3.0
		_, err := BMI(inputHeight, inputWeight)
		if err == nil {
			t.Fatalf("期望出现体重为负错误，结果却没有")
		}
	}
}
