package main

import (
	"log"
)

type Client interface {
	Register() error
	UpdateData()
	GetRank() error
}

type Person struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Sex    string  `json:"sex"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`

	FatRate float64 `json:"fat_rate"`
}

func (p *Person) Register() error {
	// 获取个人信息
	input := &inputFromStd{}
	if err := input.GetInput(p); err != nil {
		log.Println("录入个人信息失败: ", err)
		return err
	}

	// 计算体脂率
	c := &Calc{}
	if err := c.CalcFatRate(p); err != nil {
		log.Println("计算FatRate出错: ", err)
		return err
	}

	// 以json格式保存个人信息到文件
	r := &Record{
		filePath: "C:\\Users\\Administrator\\Desktop\\records.txt",
	}
	if err := r.saveRecord(p); err != nil {
		log.Println("保存文件失败: ", err)
		return err
	}

	return nil
}

func (p *Person) UpdateData() {

}

func (p *Person) GetRank() error {
	// 注册进排行榜，并进行排名，由GetRank打印出来
	var r Rank = &FatRateRank{
		clients: map[float64]*Person{},
	}
	r.Register(p)
	r.GetRank(p)
	return nil
}
