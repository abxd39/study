package main

import "fmt"

func main(){
	str:="http://qrcode.ibcmall.vip/wechat/qr/108?device_id=0&code=1333HHHHlGIltKjGklsglolssiqtkjilj"
	fmt.Println(len(str))
	if len(str)==90{
		newstr:=str[57:]
		fmt.Println(newstr)
		fmt.Println(len(newstr))
	}
}
