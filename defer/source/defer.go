package source

import "fmt"

func DoDefer(string1,string2 string) error {
	if string1==""|| string2==""{
		return fmt.Errorf("参数错误")
	}
	defer func(str1,str2 string) {
		_=str1+str2
	}(string1,string2)
	return nil
}

func NoneDefer(string1,string2 string){
	_=string1+string2
}


//func Temp(){
//	fmt.Println("什么情况")
//}