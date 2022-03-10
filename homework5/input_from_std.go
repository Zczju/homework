package main

import (
	"fmt"
	"log"
)

type inputFromStd struct {
}

func (inputFromStd) GetInput(p *Person) error {
	var name string
	fmt.Print("姓名: ")
	_, err := fmt.Scanln(&name)

	var age int64
	fmt.Print("年龄: ")
	_, err = fmt.Scanln(&age)

	var sex string
	fmt.Print("性别(男/女): ")
	_, err = fmt.Scanln(&sex)

	var tall float64
	fmt.Print("身高(米): ")
	_, err = fmt.Scanln(&tall)

	var weight float64
	fmt.Print("体重(千克): ")
	_, err = fmt.Scanln(&weight)

	if err != nil {
		log.Fatal("scan输入时出错: ", err)
		return err
	}
	p.Name = name
	p.Age = age
	p.Sex = sex
	p.Tall = tall
	p.Weight = weight

	return nil
}
