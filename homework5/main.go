package main

import "log"

func main() {
	var cli Client = &Person{}

	if err := cli.Register(); err != nil {
		log.Fatal("客户信息注册失败: ", err)
	}

	cli.UpdateData()

	if err := cli.GetRank(); err != nil {
		log.Fatal("客户获取排名失败: ", err)
	}

}
