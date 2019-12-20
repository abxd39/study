package main

import (
	"fmt"
	"github.com/mikemintang/go-curl"
)

//func main() {
//	//req := curl.New("https://kernel.org/pub/linux/kernel/v4.x/linux-4.0.4.tar.xz")
//	req := curl.New("C:/Users/Public/Pictures/Sample Pictures/first3.jpg")
//
//	req.Method("POST") // can be "PUT"/"POST"/"DELETE" ...
//
//	req.Header("MyHeader", "Value") // Custom header
//	req.Headers = http.Header{      // Custom all headers
//		"User-Agent": {"mycurl/1.0"},
//	}
//
//	ctrl := req.ControlDownload() // Download control
//	go func() {
//		// control functions are thread safe
//		ctrl.Stop()   // Stop download
//		ctrl.Pause()  // Pause download
//		ctrl.Resume() // Resume download
//	}()
//
//	req.DialTimeout(time.Second * 10) // TCP Connection Timeout
//	req.Timeout(time.Second * 30)     // Download Timeout
//
//	// Print progress status per one second
//	req.Progress(func(p curl.ProgressStatus) {
//		log.Println(
//			"Stat", p.Stat, // one of curl.Connecting / curl.Downloading / curl.Closed
//			"speed", curl.PrettySpeedString(p.Speed),
//			"len", curl.PrettySizeString(p.ContentLength),
//			"got", curl.PrettySizeString(p.Size),
//			"percent", p.Percent,
//			"paused", p.Paused,
//		)
//	}, time.Second)
//	/*
//		2015/05/20 15:34:15 Stat 2 speed 0.0B/s len 78.5M got 0.0B percent 0 paused true
//		2015/05/20 15:34:16 Stat 2 speed 0.0B/s len 78.5M got 0.0B percent 0 paused true
//		2015/05/20 15:34:16 Stat 2 speed 394.1K/s len 78.5M got 197.5K percent 0.0024564497 paused false
//		2015/05/20 15:34:17 Stat 2 speed 87.8K/s len 78.5M got 241.5K percent 0.0030038392 paused false
//		2015/05/20 15:34:17 Stat 2 speed 79.8K/s len 78.5M got 281.5K percent 0.003501466 paused false
//		2015/05/20 15:34:18 Stat 2 speed 63.9K/s len 78.5M got 313.5K percent 0.0038995675 paused false
//	*/
//
//	res, err := req.Do()
//	if err!=nil{
//		log.Println(err)
//	}
//	//ob:=res.(http.Response)                             // related *http.Response struct
//	log.Println(res.Body)                        // Body in string
//	log.Println(res.StatusCode)                  // HTTP Status Code: 200,404,302 etc
//	//log.Println(res.Hearders)                    // Reponse headers
//	log.Println(res.DownloadStatus.AverageSpeed) // Average speed
//}


func main(){
	//url := "http://php.dev/api.php"
	url := "http://61.129.70.111/webedit/api/wiki/TestTool.aspx?name=api/material/images/add"

	headers := map[string]string{
		"User-Agent":    "Sublime",
		"Authorization": "Bearer access_token",
		"Content-Type":  "application/json",
	}

	//cookies := map[string]string{
	//	"userId":    "12",
	//	"loginTime": "15045682199",
	//}
	//
	//queries := map[string]string{
	//	"page": "2",
	//	"act":  "update",
	//}

	postData := map[string]interface{}{
		"image":      "",
		"key":       "4b99ec1244f5475ab781356aec822463",
		"username":"Ebanswers",
		//"interests": []string{"basketball", "reading", "coding"},
		//"isAdmin":   true,
	}

	// 链式操作
	req := curl.NewRequest()
	resp, err := req.
		SetUrl(url).
		SetHeaders(headers).
		//SetCookies(cookies).
		//SetQueries(queries).
		SetPostData(postData).
		Post()

	if err != nil {
		fmt.Println(err)
	} else {
		if resp.IsOk() {
			fmt.Println(resp.Body)
		} else {
			fmt.Println(resp.Raw)
		}
	}
}