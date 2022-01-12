package gobmi

import "testing"

func TestCalcFatRate(t *testing.T) {
	{
		inputBmi, inputAge, inputSex := 20.0, 30, "男"
		expectingOutput := 0.147000
		actualOutput, _ := CalcFatRate(inputBmi, inputSex, inputAge)
		if actualOutput != expectingOutput {
			t.Fatalf("期望体脂率为0.147，结果却为%f", actualOutput)
		}
	}
	{
		inputBmi, inputAge, inputSex := 0.0, 30, "女"
		_, err := CalcFatRate(inputBmi, inputSex, inputAge)
		if err == nil {
			t.Fatalf("期望出现bmi为零错误，结果却没有")
		}
	}
	{
		inputBmi, inputAge, inputSex := -30.0, 30, "女"
		_, err := CalcFatRate(inputBmi, inputSex, inputAge)
		if err == nil {
			t.Fatalf("期望出现bmi为负错误，结果却没有")
		}
	}
	{
		inputBmi, inputAge, inputSex := 20.0, 0, "男"
		_, err := CalcFatRate(inputBmi, inputSex, inputAge)
		if err == nil {
			t.Fatalf("期望出现年龄为零错误，结果却没有")
		}
	}
	{
		inputBmi, inputAge, inputSex := 20.0, -15, "男"
		_, err := CalcFatRate(inputBmi, inputSex, inputAge)
		if err == nil {
			t.Fatalf("期望出现年龄为负错误，结果却没有")
		}
	}
	{
		inputBmi, inputAge, inputSex := 20.0, 200, "男"
		_, err := CalcFatRate(inputBmi, inputSex, inputAge)
		if err == nil {
			t.Fatalf("期望出现年龄超过150错误，结果却没有")
		}
	}
	{
		inputBmi, inputAge, inputSex := 20.0, 30, "wei"
		_, err := CalcFatRate(inputBmi, inputSex, inputAge)
		if err == nil {
			t.Fatalf("期望出现性别录入错误，结果却没有")
		}
	}

}
