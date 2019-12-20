package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
)

func main() {
	file1, _ := os.Open("C:/Users/Public/Pictures/Sample Pictures/1080X1920/winding-road-1556177.jpg")
	defer file1.Close()
	var (
		img1 image.Image
		err  error
	)
	if img1, _, err = image.Decode(file1); err != nil {
		log.Fatal(err)
		return
	}
	b1 := img1.Bounds()

	new1W := fixSize(b1.Max.X)
	// 调用resize库进行图片缩放(高度填0，resize.Resize函数中会自动计算缩放图片的宽高比)
	m1 := resize.Resize(uint(new1W), 0, img1, resize.Lanczos3)
	newWith := m1.Bounds().Max.X
	newHeight := m1.Bounds().Max.Y
	newImg := image.NewRGBA(image.Rect(0, 0, newWith, newHeight))
	draw.Draw(newImg, newImg.Bounds(), m1, m1.Bounds().Min, draw.Over)

	buf := new(bytes.Buffer)
	//生成新图片new.jpg,并设置图片质量
	imgw, err := os.Create("C:/Users/Public/Pictures/110.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer imgw.Close()
	err = jpeg.Encode(buf, newImg, &jpeg.Options{100})
	if err != nil {
		fmt.Println(err)
		return
	}
	imgw.Write(buf.Bytes())

}

func Response(url string) ([]byte, error) {
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
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if len(body) == 0 {
		err = fmt.Errorf("图片大小为零 picturePath=%v", url)
		return nil, err
	}
	return body, nil

}

func ProduceQrImage(url string) ([]byte, error) {
	return qrcode.Encode(url, qrcode.Medium, 256)
}

func fixSize(img1W int) (new1W int) {
	var (
		img1With = float64(img1W)
		ratio1   float64
	)
	// 如果最小宽度大于800，那么两张图片都需要进行缩放
	if img1With > 800 {
		ratio1 = 800 / img1With
		return int(img1With * ratio1)
	}
	//如果最小宽度小于800，那么需要将较大的图片缩放，使得两张图片的宽度一致
	return img1W

}
