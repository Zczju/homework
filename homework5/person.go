package main

import (
	"crypto/rand"
	"log"
	"math/big"
)

type Client interface {
	Register() error
	UpdateData()
	GetRank() error
}

type Person struct {
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
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
	// 随机年龄
	randomCryptoAge, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Println("随机年龄时出错: ", err)
	}
	p.Age = randomCryptoAge.Int64()

	// 随机身高
	randomCryptoTall, err := rand.Int(rand.Reader, big.NewInt(30))
	if err != nil {
		log.Println("随机身高时出错: ", err)
	}
	p.Tall = float64(randomCryptoTall.Int64()) * 0.1

	// 随机体重
	randomCryptoWeight, err := rand.Int(rand.Reader, big.NewInt(10000))
	if err != nil {
		log.Println("随机体重时出错: ", err)
	}
	p.Weight = float64(randomCryptoWeight.Int64()) * 0.01

}

func (p *Person) GetRank() error {
	// 注册进排行榜，并进行排名，由GetRank打印出来
	r.Register(p)
	r.GetRank(p)
	return nil
}
