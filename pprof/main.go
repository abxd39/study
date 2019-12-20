package main

import (
	"log"
	"net/http"
	_"net/http/pprof"
)

func init(){
	log.SetFlags(log.Lshortfile|log.Ldate|log.Ltime)
}
func main(){
	go func() {
		log.Println("wyw study pprof")
	}()
	http.ListenAndServe("0.0.0.0:6060",nil)
}
