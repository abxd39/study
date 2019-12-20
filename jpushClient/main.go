
package main

import (
"fmt"
"github.com/ylywyn/jpush-api-go-client"
)

const (
	appKey = "you jpush appkey"
	secret = "you jpush secret"
)

func main() {

	//Platform
	var pf jpushclient.Platform
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	pf.Add(jpushclient.WINPHONE)
	//pf.All()

	//Audience
	var ad jpushclient.Audience
	s := []string{"1", "2", "3"}
	ad.SetTag(s)
	ad.SetAlias(s)
	ad.SetID(s)
	//ad.All()

	//Notice
	var notice jpushclient.Notice
	notice.SetAlert("alert_test")
	notice.SetAndroidNotice(&jpushclient.AndroidNotice{Alert: "AndroidNotice"})
	notice.SetIOSNotice(&jpushclient.IOSNotice{Alert: "IOSNotice"})
	notice.SetWinPhoneNotice(&jpushclient.WinPhoneNotice{Alert: "WinPhoneNotice"})

	var msg jpushclient.Message
	msg.Title = "Hello"
	msg.Content = "你是ylywn"

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetMessage(&msg)
	payload.SetNotice(&notice)

	bytes, _ := payload.ToBytes()
	fmt.Printf("%s\r\n", string(bytes))

	//push
	c := jpushclient.NewPushClient(secret, appKey)
	str, err := c.Send(bytes)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	} else {
		fmt.Printf("ok:%s", str)
	}
}
