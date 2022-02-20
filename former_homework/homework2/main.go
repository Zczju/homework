package main

//
//import (
//	"fmt"
//	_ "github.com/armstrongli/go-bmi"
//	gobmi "github.com/armstrongli/go-bmi"
//)
//
//func main() {
//	for {
//		mainFatRateBody()
//
//		if cont := whetherContinue(); !cont {
//			break
//		}
//	}
//
//}
//
//func mainFatRateBody() {
//	defer recoverMainBody()
//	_, age, sex, tall, weight := getInfosInput()
//	bmi, errOfBmi := gobmi.BMI(tall, weight)
//	if errOfBmi != nil {
//		fmt.Printf("Warning:  can't calcualte bmi: %v\n", errOfBmi)
//		return
//	}
//
//	fatRate, errOfFatRate := gobmi.CalcFatRate(bmi, sex, age)
//	if errOfFatRate != nil {
//		fmt.Printf("Warning: can't calcualte fatrate: %v\n", errOfFatRate)
//		return
//	}
//	fmt.Println("体脂率为: ", fatRate)
//
//	var checkHealthinessFunc func(age int, fatRate float64)
//	if sex == "男" {
//		checkHealthinessFunc = getHealthySuggestionsForMale
//	} else {
//		checkHealthinessFunc = getHealthySuggestionsForFemale
//	}
//	getHealthySuggestions(age, fatRate, checkHealthinessFunc)
//}
//
//func recoverMainBody() {
//	if re := recover(); re != nil {
//		fmt.Println("Warning: Catch critical error: ", re)
//	}
//
//}
//
//func getHealthySuggestions(age int, fatRate float64, getSuggestions func(age int, fatRate float64)) {
//	getSuggestions(age, fatRate)
//}
//
//func getHealthySuggestionsForMale(age int, fatRate float64) {
//	if age >= 18 && age <= 39 {
//		if fatRate <= 0.1 {
//			fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
//		} else if fatRate > 0.1 && fatRate <= 0.16 {
//			fmt.Println("目前状况是：标准；请继续保持！")
//		} else if fatRate > 0.16 && fatRate <= 0.21 {
//			fmt.Println("目前状况是：偏重；少吃点长点心吧！")
//		} else if fatRate > 0.21 && fatRate <= 0.26 {
//			fmt.Println("目前状况是：肥胖；不能再吃了哥！")
//		} else {
//			fmt.Println("目前状况是：严重肥胖；该减减肥了！")
//		}
//
//	} else if age >= 40 && age <= 59 {
//		if fatRate <= 0.11 {
//			fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
//		} else if fatRate >= 0.11 && fatRate <= 0.17 {
//			fmt.Println("目前状况是：标准；请继续保持！")
//		} else if fatRate >= 0.17 && fatRate <= 0.22 {
//			fmt.Println("目前状况是：偏重；少吃点长点心吧！")
//		} else if fatRate >= 0.22 && fatRate <= 0.27 {
//			fmt.Println("目前状况是：肥胖；不能再吃了哥！")
//		} else {
//			fmt.Println("目前状况是：严重肥胖；该减减肥了！")
//		}
//
//	} else if age >= 60 {
//		if fatRate <= 0.13 {
//			fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
//		} else if fatRate >= 0.13 && fatRate <= 0.19 {
//			fmt.Println("目前状况是：标准；请继续保持！")
//		} else if fatRate >= 0.19 && fatRate <= 0.24 {
//			fmt.Println("目前状况是：偏重；少吃点长点心吧！")
//		} else if fatRate >= 0.24 && fatRate <= 0.29 {
//			fmt.Println("目前状况是：肥胖；不能再吃了哥！")
//		} else {
//			fmt.Println("目前状况是：严重肥胖；该减减肥了！")
//		}
//	} else {
//		fmt.Println("小孩子没有体脂，请长大点再来吧")
//	}
//}
//
////getHealthySuggestionsForFemale
//func getHealthySuggestionsForFemale(age int, fatRate float64) {
//	if age >= 18 && age <= 39 {
//		if fatRate <= 0.2 {
//			fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
//		} else if fatRate > 0.2 && fatRate <= 0.27 {
//			fmt.Println("目前状况是：标准；请继续保持！")
//		} else if fatRate > 0.27 && fatRate <= 0.34 {
//			fmt.Println("目前状况是：偏重；少吃点长点心吧！")
//		} else if fatRate > 0.34 && fatRate <= 0.39 {
//			fmt.Println("目前状况是：肥胖；不能再吃了姐！")
//		} else {
//			fmt.Println("目前状况是：严重肥胖；该减减肥了！")
//		}
//
//	} else if age >= 40 && age <= 59 {
//		if fatRate <= 0.21 {
//			fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
//		} else if fatRate >= 0.21 && fatRate <= 0.28 {
//			fmt.Println("目前状况是：标准；请继续保持！")
//		} else if fatRate >= 0.28 && fatRate <= 0.35 {
//			fmt.Println("目前状况是：偏重；少吃点长点心吧！")
//		} else if fatRate >= 0.35 && fatRate <= 0.40 {
//			fmt.Println("目前状况是：肥胖；不能再吃了姐！")
//		} else {
//			fmt.Println("目前状况是：严重肥胖；该减减肥了！")
//		}
//
//	} else if age >= 60 {
//		if fatRate <= 0.22 {
//			fmt.Println("目前状况是：偏瘦；要多吃肉哦！")
//		} else if fatRate >= 0.22 && fatRate <= 0.29 {
//			fmt.Println("目前状况是：标准；请继续保持！")
//		} else if fatRate >= 0.29 && fatRate <= 0.36 {
//			fmt.Println("目前状况是：偏重；少吃点长点心吧！")
//		} else if fatRate >= 0.36 && fatRate <= 0.41 {
//			fmt.Println("目前状况是：肥胖；不能再吃了姐！")
//		} else {
//			fmt.Println("目前状况是：严重肥胖；该减减肥了！")
//		}
//	} else {
//		fmt.Println("小孩子没有体脂，请长大点再来吧")
//	}
//}
//
//func getInfosInput() (string, int, string, float64, float64) {
//	var (
//		name   string
//		age    int
//		sex    string
//		tall   float64
//		weight float64
//	)
//	fmt.Print("姓名：")
//	fmt.Scanln(&name)
//	fmt.Print("年龄：")
//	fmt.Scanln(&age)
//	fmt.Print("性别：")
//	fmt.Scanln(&sex)
//	fmt.Print("身高：")
//	fmt.Scanln(&tall)
//	fmt.Print("体重：")
//	fmt.Scanln(&weight)
//	return name, age, sex, tall, weight
//}
//
//func whetherContinue() bool {
//	var whetherContinue string
//	fmt.Print("是否输入下一个?(y/n)")
//	fmt.Scanln(&whetherContinue)
//	if whetherContinue != "y" {
//		return false
//	}
//	return true
//}
