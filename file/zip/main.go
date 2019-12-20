package main

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func init()  {
	log.SetFlags(log.Ltime|log.Ldate|log.Lshortfile)
}

type fZip struct {

}
func (fz*fZip) zipFile() error {
	originPath:="C:/Users/yingwenwang/Pictures/123.rar"
	readFunc := func()  {
		// Open a zip archive for reading.
		r, err := zip.OpenReader(originPath)
		if err != nil {
			log.Println(err)
			return
		}
		defer r.Close()

		// Iterate through the files in the archive,
		// printing some of their contents.
		for _, f := range r.File {
			fInfo:=f.FileInfo()
			if fInfo.IsDir(){
				continue
			}
			log.Printf("Contents of %v:\n", f.Name)
			count:=strings.Count(f.Name,"/")
			if count<2||count >2{
				log.Println(f.Name)
				continue
			}
			indx:=strings.LastIndex(f.Name,"/")
			sub:=f.Name[:indx]
			indxb:=strings.LastIndex(sub,"/")
			no:=sub[indxb+1:]

			log.Println(no)
			rc, err := f.Open()
			if err != nil {
				log.Println(err)
				rc.Close()
				continue
			}
			log.Println("惊恐")
			_, err= ioutil.ReadAll(rc)
			if err != nil {
				log.Println(err)
				rc.Close()
				continue
			}
			//fmt.Println(base64.StdEncoding.EncodeToString(Body))

		}
		return
	}
	//缺学号
	//缺教师账号
	//如果是家长 则必须要有学生学号
	readFunc()


	return nil
}

func main()  {
	zi:=&fZip{}
	zi.zipFile()
	time.Sleep(time.Minute)
}