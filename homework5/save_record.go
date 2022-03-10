package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Record struct {
	filePath string
}

func (r *Record) saveRecord(p *Person) error {
	// 将个人信息进行JSON格式序列化
	data, err := json.Marshal(p)
	if err != nil {
		log.Println("Marshal时出错", err)
		return err
	}

	// 将个人数据写入文件
	if err := r.writeFile(data); err != nil {
		log.Println("写入JSON时出错: ", err)
		return err
	}
	return nil
}

func (r *Record) writeFile(data []byte) error {
	// 创建一个可写入文件，并打开
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("无法打开文件: ", r.filePath, "错误信息是: ", err)
		os.Exit(-1)
	}
	defer file.Close()

	// 将Marshal的内容写入文件
	_, err = file.Write(data)
	if err != nil {
		log.Println("写入文件失败: ", err)
		return err
	}
	return nil
}
