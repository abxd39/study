package main

import (
	"fmt"
	"time"
)
var (
	DATE             string = "2006-01-02"
	DATETIME         string = "2006-01-02 15:04:05"
	DATETIMET        string = "2006-01-02T15:04:05"
	YEARMONTH        string = "2006-01"
	YEARMONTHCHINESE string = "2006年01月"
	DATETIMESECOND  string="20060102150405"
)

func main(){
 	//GetZeroTimeUnix()
	//t:=StringToTime("2019-07-31 23:59:59")
	//fmt.Println(t.Unix())

	//fmt.Println(time.Unix(int64(1556208000), 0))
	fmt.Println(time.Now().Format(DATETIME))
	sec:=time.Second*24*3600-time.Second
	fmt.Printf("%v\n",sec)
	fmt.Println(time.Now().Add(sec).Format(DATETIME))

}

//获取零点时间
func GetZeroTimeUnix() int64  {
	timeStr := time.Now().Format("2006-01-02")
	t:=StringToDateTime(timeStr)
	return t.Unix()

}

func StringToDateTime(str string) time.Time {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, str, loc)
	return theTime
}

//字符串转time
func StringToTime(str string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, str, loc)
	return theTime
}

