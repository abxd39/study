package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

type INT interface {
	Name(string2 string) error
}

func main() {
	//bl := strings.Contains("wangyingwen", " ")
	//
	//fmt.Printf("%v\n", bl)
	//
	//bl=strings.HasPrefix("//","/")
	//fmt.Printf("%v\n",bl)
	//fmt.Printf("%+q\n",`√`)
	//fmt.Printf("%d\n",'a')
	//fmt.Printf("%d\n",`//`)
	//fmt.Printf("%d\n",'/')
	//GetName(strings.NewReader("wyw"))
	//fmt.Println(strings.ToUpper(time.Now().Weekday().String()))
	//
	//var err error
	//loginKey :="mEc9bxtcDsC8%20B33PyFq7A2lnmteJ23OJo%20B3bzyd2Wl4F%20BrcPozGFuh8xsUjBzfz3"
	//loginKey ="mEc9bxtcDsC8+33PyFq7A2lnmteJ23OJo+3bzyd2Wl4F+rcPozGFuh8xsUjBzfz3"
	//fmt.Printf("%+v\r\n",loginKey)
	//loginKey,err= url.QueryUnescape(loginKey)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(loginKey)
	//loginKey=strings.Replace(loginKey,"","+",-1)
	//fmt.Println(loginKey)
	//
	//b, err := base64.StdEncoding.DecodeString(loginKey)
	//if err != nil {
	//	fmt.Println("111",err)
	//	return
	//}
	//fmt.Println(string(b))

	//name1 := "王欧米伽味美"
	//fmt.Println(fmt.Sprintf(` s.name like '%%%s%%'`, name1))
	//return
	//data := "2019-03-07 00:00:00"
	//fmt.Println(data[:10])
	//Code := "G346HHHHqqIIli,Hsklsglok,isoqglotgs,"
	//fmt.Println(len(Code))
	//newCode := strings.Replace(Code, ",", "", -1)
	//
	//fmt.Println(newCode)
	//newCode = strings.TrimRight(Code, ",")
	//fmt.Println(newCode)
	//fmt.Println(len(newCode))
	//str := "prepay_id=wx17173709364960faadc3d8093207779401"
	//str = strings.Replace(str, "prepay_id=", "", 1)
	//fmt.Println(str)

	//arrayAll := []rune{'a', 'b', 'c', 'd', 'e'}
	//var p = []int{0: 9}
	//x := p[arrayAll[0]-'a']
	//fmt.Println(p)
	//fmt.Println(x)
	//fmt.Println(arrayAll)
	//
	//fmt.Println(strconv.FormatInt(2, 2))
	//f := 1.23
	//
	//fmt.Println(reflect.TypeOf(f).Name())
	//fmt.Println(strconv.FormatFloat(1.23456, 'f', 3, 32))
	//shopName := ",斯凯奇儿童,"
	//Name1 := strings.TrimLeft(shopName, ",")
	//Name2 := strings.TrimRight(Name1, ",")
	//fmt.Println(Name2)
	//str = "78,48,62,73,49,84,86,89,20,21,27,74,76,77,88,54,55,64,30,72,17,22,29,57,59,26,28,50,69,75,19,31,41,23,46,53"
	//list1 := strings.Split(str, ",")
	//sort.Strings(list1)
	//fmt.Println(list1)
	//AlphabetToX("1592008315")
	//str := "78"
	//strlist := strings.Split(str, ",")
	//fmt.Println(strlist)
	fmtttt()
	//GetName(&name{"wyw"})
}

type strL struct {
	list []string
}

func (s *strL) Len() int {
	return len(s.list)
}

func (s *strL) Less(i, j int) bool {
	if s.list[i] > s.list[j] {
		return false
	}
	return true
}

func (s *strL) Swap(i, j int) {
	s.list[i], s.list[j] = s.list[j], s.list[i]
}

type name struct {
	Na string
}

func (n *name) Name(na string) error {
	fmt.Println("function Name")
	return nil
}
func (na *name) Read(p []byte) (n int, err error) {
	fmt.Println("function Read")
	return 0, nil
}

func GetName(in io.Reader) error {
	val, bl := in.(INT)

	fmt.Printf("%v\n", bl)
	fmt.Printf("%+v", val)
	return nil
}

func AlphabetToX(beferText string) string {
	if len(beferText) < 3 {
		return beferText
	}
	replaceAlphabet := "**************************************************"
	plength := 3
	slength := 4
	if len(beferText) < 8 {
		plength = 1
		slength = 1
	}
	suffix := ""
	prefix := ""
	index := len(beferText) - slength
	suffix = beferText[index:]
	prefix = beferText[:plength]
	middleLength := len(beferText[plength:index])
	middleAlphabet := replaceAlphabet[:middleLength]
	afterText := prefix + middleAlphabet + suffix
	log.Println(afterText)
	return afterText
}


func fmtttt(){

	flatAmount, err := strconv.ParseFloat("-111111111111", 64)
	if err != nil {
		log.Println(err)

	}
	log.Println( int(flatAmount * 100))
}