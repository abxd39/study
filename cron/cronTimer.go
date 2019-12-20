package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/signal"
	"time"
)

const LAYOUT ="2006-01-02 15:04:05"

const OneSecond = 5*time.Second

func FuncPanicRecovery() {
	cron := cron.New()
	cron.Start()
	defer cron.Stop()
	cron.AddFunc("* * * * * ?", Cron)

	select {
	case <-time.After(OneSecond):
		fmt.Printf("五秒中输出一次%v",time.Now().Format("2006-01-02 15:04:05"))
		return
	}
}

func Cron(){
	now := time.Now().Truncate(time.Minute)
	stime := now
	//etime := now.Add(-time.Minute * 5)
	//fmt.Println("etime",etime)
	fmt.Println("stime",stime)
	fmt.Printf("五分钟输出一次%v\r\n",time.Now().Format("2006-01-02 15:04:05"))
}

var l *log.Logger
var couponCount int32 = 6100001
func main()  {
	//tm:=OneMomentClock("2019-01-21 14:32:08",LAYOUT)
	//tm=time.Now()
	//fmt.Println(tm.UnixNano())
	//str:=tm.Format("2006010215")
	//fmt.Println(str[2:])
	//Count:=atomic.AddInt32(&couponCount,1)
	//fmt.Println(Count)
	//fmt.Println(couponCount)

	/*************/
	//unix:=1550109560000
	//tm:=UnixUano2Time(1550109560)
	//fmt.Println(tm.Format("2006-01-02 15:04:05"))
	/*************/
	t:=time.Time{}
	fmt.Println(t.IsZero())
	t= time.Now()
	fmt.Println(t.IsZero())
	/**************/
	//GetDate()
	/**************/
	//cor:=cron.New()
	//cor.Start()
	//defer cor.Stop()
	//cor.AddFunc("0 0 2 * * ?", func() {fmt.Printf("凌晨两点[%v]",time.Now().Format("2006-01-02 15:04:05"))})
	//cor.AddFunc("@midnight", func() {fmt.Printf("@midnight[%v]",time.Now().Format("2006-01-02 15:04:05"))})

	//Wait for interrupt signal to gracefully shutdown the server with
	//a timeout of 30 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

// 日期转时间戳
func Date2Unix(date string, layout string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(layout, date, loc)
	return theTime.Unix()
}

// 时间戳转日期
func Unix2Date(unix int64, layout string) string {
	return time.Unix(unix, 0).Format(layout)
}

func OneMomentClock(date,layout string) time.Time {
	fmt.Println(date)
	loc, _ := time.LoadLocation("Local")
	now, _:= time.ParseInLocation(layout, date, loc)
	timestamp := now.Unix() - int64(now.Second()) - int64(60 * (now.Minute()%15))
	fmt.Println(timestamp, time.Unix(timestamp, 0).Format(layout), now.Unix())
	return time.Unix(timestamp,0)
}


func dateNew(){
	fmt.Println(time.Now().Format("20060102"))
	fmt.Println(time.Now().Format("20060102150405"))
}

func UnixUano2Time(unixNano int64)time.Time{
	return time.Unix(unixNano,0)
}


//获取昨天的日期
func GetDate(){
	nTime := time.Now()
	yesTime := nTime.AddDate(0,0,-1)
	logDay := yesTime.Format("20060102")
	fmt.Println(logDay)
}