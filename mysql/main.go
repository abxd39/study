package main

import (
	"fmt"
	"github.com/abxd39/myproject/mysql/model"
	"os"
	"time"
)

type GooagooShop struct {
	Id            int       `xorm:"not null pk autoincr INT(11)"`
	ShopId        int       `xorm:"not null comment('商户id') INT(11)"`
	GooagooShopId string    `xorm:"not null comment('gooagoo 提供的商家id') VARCHAR(255)"`
	GooagooMallId string    `xorm:"not null comment('gooagoo mallid') VARCHAR(255)"`
	ShopName      string    `xorm:"not null comment('gooagoo 提供的商家名称') VARCHAR(255)"`
	Created       time.Time `xorm:"created"`
	Updated       time.Time `xorm:"updated"`
	PullTime      time.Time `xorm:"not null default '0000-00-00 00:00:00' comment('三秒一次订单最后抓取的时间 服务重启需要根据此时间继续抓取') TIMESTAMP"`
}

func main() {
	fmt.Println(os.Args[0])
	model.ConfigPath="e:/workSpace/src/github.com/abxd39/myproject/mysql/config/config.json"
	err:=model.NewSession()
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer model.TestDb.Close()
	list := make([]GooagooShop, 0)
	model.TestDb.Table("gooagoo_shop_repair").Where("bool_pull=1").Limit(1, 2).Find(&list)
}


func(g*GooagooShop)GetList(){
	list := make([]GooagooShop, 0)
	err:=model.TestDb.Table("gooagoo_shop_repair").Where("bool_pull=1").Limit(1, 2).Find(&list)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}