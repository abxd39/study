package source

import (
	"testing"
)


const 	NO ="\u2717"
const 	YES ="\u2714"




func BenchmarkDoDefer(b *testing.B){
	for i:=0;i<=b.N;i++{
		DoDefer("中文","https://segmentfault.com/a/1190000")
	}
}

func BenchmarkNoneDefer(b *testing.B){
	for i:=0;i<b.N;i++  {
		NoneDefer("中文","https://segmentfault.com/a/1190000")
	}
}

func TestMain(t*testing.M){
	//额外的初始化
	t.Run()
}

func TestDoDefer(t *testing.T) {
	tests:=[] string{
		"中国","english","",
	}
	err:=DoDefer(tests[0],tests[1])
	if err!=nil{
		t.Fatal(err)
	}
	err = DoDefer(tests[1],tests[2])
	if err==nil{
		t.Fatal("参数string2为空测试失败！！",NO)
	}
	err = DoDefer(tests[2],tests[0])
	if err==nil{
		t.Fatal("参数string1为空测试失败",NO)
	}
	err =DoDefer(tests[2],tests[2])
	if err==nil{
		t.Fatal("参数string1, string2 都为空测试失败" ,NO)
	}
	t.Log("tesetDodefer",YES)
}