package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile | log.Ldate)
}

type BenefitStatus struct {
	BatchNo string `json:"batch_no"`
	Status  int    `json:"status"`
}

func main() {

	log.Println("bengin", time.Now().Format("2006-01-02 15:04:05"))
	o := "0000000000"
	context := make([]BenefitStatus, 0)
	for i := 0; i < 10000; i++ {
		count := fmt.Sprintf("%d", i)
		l := len(count)
		b := BenefitStatus{
			BatchNo: o[:8-l] + count,
			Status:  1,
		}
		context = append(context, b)
	}
	by, err := json.Marshal(context)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("文件大小是 %dk ", binary.Size(by)/1024)
	err = ioutil.WriteFile("./ttt.json", by, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
	log.Println("bengin", time.Now().Format("2006-01-02 15:04:05"))
	new(BenefitStatus).Unmarshal()

}

func (bi *BenefitStatus) Unmarshal() {
	log.Println("开始读文件解析", time.Now().Format("2006-01-02 15:04:05"))
	context := make([]BenefitStatus, 0)
	b, err := ioutil.ReadFile("./ttt.json")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("读取的文件大小是 %dk ", binary.Size(b)/1024)
	err = json.Unmarshal(b, &context)
	if err != nil {
		log.Println(err)
	}
	log.Printf("总条目数%d", len(context))
	log.Println("读文件解析接受", time.Now().Format("2006-01-02 15:04:05"))
}
