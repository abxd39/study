package main

import (
	"MyProject/config/mq"
	"fmt"
	"os"
	"os/signal"
)

func main()  {
	mq.LoadConfig()
	//fmt.Printf("%v\r\n",mq.Config)
	fmt.Println(mq.Config.MaxQueueSize,mq.Config.MaxWork)
	mq.LoadCfg()
	//fmt.Printf("%v\r\n",mq.Cfg)
	//Wait for interrupt signal to gracefully shutdown the server with
	//a timeout of 30 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
