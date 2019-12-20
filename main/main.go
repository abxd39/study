package main

import (
	manageMq "MyProject/manageMq/temp"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	value := "我也不知道啊!!"
	err := fmt.Errorf("这是怎么回事啊-------》%s", value)
	fmt.Println(err)
	exchange := "exchange"
	log.Printf("got Channel, declaring Exchange (%q)", exchange)
	ExampleLoggerOutput(err.Error())
	manageMq.InitMq()


	err = manageMq.GlobalMq.Ping()
	if err != nil {
		ExampleLoggerOutput(err.Error())
	}
	//manageMq.GlobalMq.Publish("amq.fanout", "{\"phone\":\"15920038315\",\"message\":\"fuck you\"}")
	manageMq.GlobalMq.ReceiveMessage("yf_manage_message")
	//i := 1
	//for {
	//	fmt.Printf("第%d次推送信息", i)
	//	manageMq.GlobalMq.Publish("amq.fanout", "{\"phone\":\"15920038315\",\"message\":\"fuck you\"}")
	//	time.Sleep(5 * time.Second)
	//	i++
	//
	//}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 30 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	manageMq.GlobalMq.Shutdown("yf_sms_consumer")
	ExampleLoggerOutput("程序退出")
	//phone := "15920038315"
	//message := "are you sure??"
	//message = fmt.Sprintf("{\"phone\":\"%q\",\"message\":\"%q\"}", phone, message)
	//fmt.Println(message)
}

func ExampleLoggerOutput(info string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)

		infoMessage = func(info string) {
			logger.Output(2, info)
		}
	)

	infoMessage(info)

	fmt.Print(&buf)
	// Output:
	// INFO: example_test.go:36: Hello world
}
