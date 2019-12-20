package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Temp struct {

}

func main()  {

	new(Temp).ResponseImage("http://61.129.70.111/Project/UserImages/201906111120572309.jpg")
}

func (t *Temp) ResponseImage(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	rsp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rsp.Body.Close()
	if len(body) == 0 {
		err = fmt.Errorf("图片大小为零 picturePath=%v", url)
		return nil, err
	}
	return body, nil

}

