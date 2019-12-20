package main

import (
	"github.com/abxd39/myproject/mysql/model"
	"log"
	"testing"
)

func TestMain(m *testing.M)  {
	log.Println("begin")
	model.ConfigPath="e:/workSpace/src/github.com/abxd39/myproject/mysql/config/config.json"
	model.NewSession()
	defer model.TestDb.Close()
	log.Println("end")
}


func TestGooagooShop_GetList(t *testing.T) {
	new(GooagooShop).GetList()
}

func BenchmarkGooagooShop_GetList(b *testing.B) {
	for i:=0;i<b.N;i++ {
		new(GooagooShop).GetList()
	}
}