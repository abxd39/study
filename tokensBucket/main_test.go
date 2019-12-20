package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/trace"
	"testing"
)

func TestBegin(t *testing.T) {
	t.Log("这又是怎么回事呢！！！")
	for i:=0;i<10;i++{
		err :=Begin()
		if err!=nil{
			t.Log("---> \u2718")
		}else {
			t.Log("---> \u2714")
		}
	}

	t.Log("begin end !!")
}

func TestMiddleware(t *testing.T) {
	t.Log("second")
}

func TestMain(t *testing.M) {
	file,err:=os.Create("E:/WorkSpace/src/github.com/abxd39/myproject/tokensBucket/trace.out")
	if err!=nil{
		return
	}
	trace.Start(file)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.Println("first \u2714")
	t.Run()
	log.Println("\u2718")
	file.Close()
	trace.Stop()
}
