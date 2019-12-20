package main

import (
	"context"
	"runtime"
	"testing"
)

func TestPayload1_Fuck(t *testing.T) {

}

func TestWorking(t *testing.T) {
	var slice = []int{
		1,10,15,20,30,40,
	}
	for v:=range slice{
		runtime.GOMAXPROCS(v)
		Working(context.TODO())
	}

}