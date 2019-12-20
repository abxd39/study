package main

import (
	"reflect"

	"github.com/lunny/log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

type Ainmal interface {
	Fly() error
	Eat() error
}

func main() {
	pig := &Pig{
		Class: "宠物猪",
	}
	bird := &Bird{
		Class: "鹦鹉",
		Area:  "非洲",
	}
	global(pig)
	global(bird)
	log.Printf("pig class -->%v", pig.Class)
	log.Printf("bird class -->%v,area -->%v", bird.Class, bird.Area)

}

func global(face interface{}) {
	//tp := reflect.TypeOf(face)
	tv := reflect.ValueOf(face)
	st := tv.Type().Elem()

	//if nameField, found := st.FieldByName("Class"); found == false || nameField.Type.Kind() != reflect.String {
	//	return
	//}
	for i := 0; i < st.NumField(); i++ {
		//log.Println(st.Name()) //对象名
		field := st.Field(i)
		log.Println(field.Name)
		//log.Println(field.Tag)
		//log.Println(tv.String())

	}

	//if st.Name() == "Bird" {
	//	for i:=0;i<st.NumField();i++{
	//		field:=st.Field(i)
	//		if field.Name=="Class"{
	//			//log.Println(st.FieldByName(field.Name))//.SetString("麻雀")
	//			log.Println(tv.String())
	//			//tv.SetString("麻雀！！")
	//		}
	//	}
	//}

}

type Pig struct {
	Class string
}

func (p *Pig) Fly() error {
	log.Printf("%v 在飞奔！！", p.Class)
	return nil
}

func (p *Pig) Eat() error {
	log.Printf("%v 吃草！！ ", p.Class)
	return nil
}

type Bird struct {
	Class string
	Area  string
}

func (b *Bird) Fly() error {
	log.Printf("%v在翱翔！！ 栖息在%v", b.Class, b.Area)
	return nil
}

func (b *Bird) Eat() error {
	log.Printf("%v 吃虫子", b.Class)
	return nil
}
