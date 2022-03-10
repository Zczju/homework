package main

import (
	"fmt"
	"log"
)

func main() {
	for {
		var cli Client = &Person{}

		if err := cli.Register(); err != nil {
			log.Fatal("客户信息注册失败: ", err)
		}

		cli.UpdateData()

		if err := cli.GetRank(); err != nil {
			log.Fatal("客户获取排名失败: ", err)
		}

		if cont := whetherContinue(); !cont {
			break
		}
	}
}

func whetherContinue() bool {
	var whetherContinue string
	fmt.Println("是否继续输入? (y/n)")
	fmt.Scanln(&whetherContinue)
	if whetherContinue != "y" {
		return false
	}
	return true
}
