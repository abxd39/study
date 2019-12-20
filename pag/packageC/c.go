package packageC

import (
	"fmt"
	"github.com/abxd39/myproject/pag/packageA"
	"github.com/abxd39/myproject/pag/packageB"
)

type CTemp struct {

}


func (c*CTemp)Name()string{
	return "packageC"
}


func (c*CTemp)GetName(){
	name:=new(packageA.ATemp).Name()
	fmt.Println(name)
	name =new(packageB.BTemp).Name()
	fmt.Println(name)
	name = c.Name()
	fmt.Println(name)
}

func (c*CTemp)GetAName()string{
	return new(packageA.ATemp).Name()
}

func(c*CTemp)GetBName()string{
	return new(packageB.BTemp).Name()
}