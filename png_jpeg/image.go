package main

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/golang/freetype"
	"github.com/nfnt/resize"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	var backgroundImage image.Image
	var backgroundImagem image.Image
	buff, err := ioutil.ReadFile(`/Users/wangyingwen/work/project/golang/src/github.com/abxd39/study/png_jpeg/background.png`)
	if err != nil {
		log.Println(err)
		return
	}
	bufm, err := ioutil.ReadFile(`/Users/wangyingwen/work/project/golang/src/github.com/abxd39/study/png_jpeg/mid.png`)
	if err != nil {
		log.Println(err)
		return
	}

	backgroundImage, err = png.Decode(bytes.NewReader(buff))
	if err != nil {
		log.Println(err)
		return
	}
	backgroundImagem, err = png.Decode(bytes.NewReader(bufm))
	if err != nil {
		log.Println(err)
		return
	}

	img := image.NewRGBA(backgroundImage.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, backgroundImage.At(x, y))
		}
	}

	imgm := image.NewRGBA(backgroundImage.Bounds())
	for y := 0; y < imgm.Bounds().Dy(); y++ {
		for x := 0; x < imgm.Bounds().Dx(); x++ {
			imgm.Set(x, y, backgroundImagem.At(x, y))
		}
	}
	//加文字
	fontBytes, err := ioutil.ReadFile(`/Users/wangyingwen/work/project/golang/src/github.com/abxd39/study/png_jpeg/simsun.ttc`)
	if err != nil {
		log.Println(err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}

	f := freetype.NewContext()
	f.SetDPI(72)
	f.SetFont(font)
	f.SetFontSize(48)
	f.SetClip(backgroundImage.Bounds())
	f.SetDst(img)
	f.SetSrc(image.NewUniform(color.RGBA{R: 0, G: 0, B: 0, A: 255}))

	pt := freetype.Pt(img.Bounds().Dx()/4, 50)
	_, err = f.DrawString("健康未上报名单为", pt)
	pt = freetype.Pt(img.Bounds().Dx()/3, 108)
	f.SetFontSize(12)
	str := time.Now().Format("2006/01/02 15:04:05")
	_, err = f.DrawString(str, pt)

	f.SetClip(backgroundImagem.Bounds())
	f.SetDst(imgm)
	f.SetSrc(image.NewUniform(color.RGBA{R: 0, G: 0, B: 0, A: 255}))
	f.SetFontSize(24)
	//pt := freetype.Pt(img.Bounds().Dx()-20, img.Bounds().Dy()-12)
	pt = freetype.Pt(10, imgm.Bounds().Dy()-12)
	_, err = f.DrawString("西皮 新富 老韩 孝勇 特朗普", pt)

	//生成新的图片
	sub, err := ioutil.ReadFile(`/Users/wangyingwen/work/project/golang/src/github.com/abxd39/study/png_jpeg/1582785013.png`)
	if err != nil {
		log.Println(err)
		return
	}
	// 缩略图的大小
	// 产生缩略图,等比例缩放
	iqr, _, err := image.Decode(bytes.NewReader(sub))
	if err != nil {
		return
		log.Println(err)

	}
	imgNew := resize.Resize(uint(backgroundImage.Bounds().Dx()), 1000, img, resize.Lanczos3)
	//设置二维码在背景图片的右下角
	var offsetm image.Point
	var offsetl image.Point
	//中间

	offsetm = image.Pt(0, backgroundImage.Bounds().Dy())
	offsetl = image.Pt(0, backgroundImage.Bounds().Dy()+backgroundImagem.Bounds().Dy())

	//offset:=image.Pt(10+50,10+50)
	newImage := image.NewRGBA(image.Rect(0, 0, imgNew.Bounds().Dx(), imgNew.Bounds().Dy()))
	for y := 0; y < newImage.Bounds().Dy(); y++ {
		for x := 0; x < newImage.Bounds().Dx(); x++ {
			c := newImage.At(x, y)
			r, g, b, _ := c.RGBA()
			if r == 0 && g == 0 && b == 0 {
				imgm.Set(x, y, color.NRGBA{255, 255, 255, 255})
			} else {
				imgm.Set(x, y, newImage.At(x, y))
			}
		}
	}

	draw.Draw(newImage, backgroundImage.Bounds(), imgNew, image.ZP, draw.Src)
	draw.Draw(newImage, backgroundImage.Bounds().Add(offsetm), imgm, image.ZP, draw.Over)
	draw.Draw(newImage, backgroundImage.Bounds().Add(offsetl), iqr, image.ZP, draw.Over)
	//
	buf := new(bytes.Buffer)

	err = jpeg.Encode(buf, newImage, &jpeg.Options{100})
	if err != nil {
		log.Println(err)
		return
	}
	//写文件
	dir := "/Users/wangyingwen/work/project/golang/src/github.com/abxd39/study/png_jpeg/"
	imgw, err := os.Create(dir + "/wyw-jpg.jpg")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = imgw.Write(buf.Bytes())
	if err != nil {
		log.Println(err)
	}
	_ = imgw.Close()
}
