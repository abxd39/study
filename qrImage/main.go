//package main
///*
//二维码生成及美化
//*/
//
//import (
//	"bytes"
//	"fmt"
//	"github.com/skip2/go-qrcode"
//	"image"
//	"image/color"
//	"image/draw"
//	"image/png"
//	"os"
//)
//
//func main()  {
//	qrImage,err:=QrImage("www.baidu.com")
//	if err!=nil{
//		fmt.Println(err)
//		return
//	}
//
//	tempImage,_,err:=image.Decode(bytes.NewReader(qrImage))
//	if err!=nil{
//		fmt.Println(err)
//		return
//	}
//
//	// Create an 100 x 50 image
//	img := image.NewRGBA(tempImage.Bounds())
//	//offset:=image.Pt(0,0)
//	draw.Draw(img,tempImage.Bounds(),tempImage,image.ZP,draw.Over)
//	//Draw a red dot at (2, 3)
//	height:=img.Bounds().Dy()
//	width :=img.Bounds().Dx()
//	for i:=0;i<height;i++{
//		for j:=0;j<width;j++{
//			r,g,b,_:=img.At(i,j).RGBA()
//			if r==0&&g==0&&b==0 {
//				img.Set(i, j, color.RGBA{23, 190, 209, 255})
//			}
//		}
//	}
//	// Save to out.png
//	 f, _ := os.OpenFile("C:\\Users\\Public\\Pictures\\Sample Pictures\\qrImage\\out.png", os.O_WRONLY|os.O_CREATE, 0600)
//	 defer f.Close()
//	 png.Encode(f, img)
//
//}
//
//
//
//
//func QrImage(url string) ([]byte,error) {
//	return  qrcode.Encode(url,qrcode.Medium,350)
//}
package main

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	//err = png.Encode(file, img)
	 err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println(file.Name())
}

func main() {
	base64 := "www.baidu.com"
	log.Println("Original data:", base64)
	code, err := qr.Encode(base64, qr.L, qr.Unicode)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data: ", code.Content())

	if base64 != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		log.Fatal(err)
	}

	writePng("C:/Users/Public/Pictures/Sample Pictures/qrImage/test20190622.png", code)
}