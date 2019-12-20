package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

func main(){
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)
	cmdo:=exec.Command("echo","-n","my first command comes from golang")
	if err:=cmdo.Start();err!=nil{
		log.Printf("err: The commend No.o can not be start up %s\n",err)
		return
	}
	stdOut,err:=cmdo.StdoutPipe()
	if err!=nil{
		log.Printf("Error: Couldn't obtiain the stdout pipe for commend No.0 :%s\n",err)
	}
	var outputBuffer bytes.Buffer
	for {
		output0 := make([]byte, 30)
		n, err := stdOut.Read(output0)
		if err != nil {
			if err==io.EOF{
				break
			}else {
				log.Printf("Error:Couldn't read data from  the pipe :%s\n", err)
				return
			}
		}
		if n>0{
			outputBuffer.Write(output0[:n])

		}
		log.Printf("%s\n", output0[:n])
	}

	io.Pipe()
	os.Pipe()
}