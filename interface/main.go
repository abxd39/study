package main

type VerifyOptions struct {
	DNSName string
}



type VerifyOption interface {
	Func(interface{})

}

func ReturnInterface()interface{}  {
	var inter interface{}
	return &inter
}

func What()VerifyOption  {
	return ReturnInterface().(VerifyOption)
}


func main() {
	//opts:=&VerifyOptions{
	//	"gmail.google.com",
	//}
	//hasDNSName := opts != nil && len(opts.DNSName) > 0
	//fmt.Printf("typeName--> %v\n",reflect.TypeOf(hasDNSName))
	//fmt.Println(hasDNSName)
	//vi:=What()
	What()
	//vi.Func("你大爷")
}
