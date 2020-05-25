package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
}

type ImportBenefit struct {
	SchoolName        string  `json:"school_name"`
	GradeName         string  `json:"grade_name"`
	ClassName         string  `json:"class_name"`
	StudentName       string  `json:"student_name"`
	IdType            string  `json:"id_type"`
	CardNo            string  `json:"card_no"`
	Amount            float64 `json:"amount"`
	Guardian1Name     string  `json:"guardian1_name"`
	Guardian1CardNo   string  `json:"guardian1_card_no"`
	Guardian1CardType string  `json:"guardian1_card_type"`
	Guardian2Name     string  `json:"guardian2_name"`
	Guardian2CardNo   string  `json:"guardian2_card_no"`
	Guardian2CardType string  `json:"guardian2_card_type" `
}
type Cell struct {
	Col int `form:"col" json:"col"` //列
	Raw int `json:"raw" form:"raw"` //行
}

//sheet表头
//sheet 内容

var importHead = map[string]string{
	"*学校名称":      "SchoolName",
	"*年级":        "GradeName",
	"*班级":        "ClassName",
	"*姓名":        "StudentName",
	"*证件号类型":     "IdType",
	"*证件号":       "CardNo",
	"*补助金额（元）":   "Amount",
	"*监护人1姓名":    "Guardian1Name",
	"*监护人1证件号类型": "Guardian1CardType",
	"*监护人1证件号":   "Guardian1CardNo",
	"监护人2姓名":     "Guardian2Name",
	"监护人2证件号类型":  "Guardian2CardType",
	"监护人2证件号":    "Guardian2CardNo",
}

func main() {
	f, err := xlsx.OpenFile(`/Users/wangyingwen/Documents/health/import-test.xlsx`)
	if err != nil {
		log.Println(err)
		return
	}

	for si, sheet := range f.Sheets {
		headMap := map[string]Cell{} //headMap["学校名称"]
		excelContent := make(map[int]ImportBenefit, 0)
		log.Println("sheet", si)
		if si > 0 {
			continue
		}
		for r, row := range sheet.Rows {
			if row == nil {
				continue
			}

			p := ImportBenefit{}
			for c, cell := range row.Cells {
				str, err := cell.FormattedValue()
				if err != nil {
					if numErr, ok := err.(*strconv.NumError); ok && numErr.Num == "" {
						str = ""
					} else {
						log.Println("game over")
						break
					}
				}
				//log.Println(k, v, str)
				str = strings.Trim(str, " ")

				if r == 1 { //抓取头部字段内容
					for k, v := range importHead {
						if strings.Compare(str, k) == 0 {
							headMap[v] = Cell{
								Col: c,
								Raw: r,
							}
						}
					}
					//log.Println("headMap ", headMap)
				} else {
					for k, v := range headMap {
						if v.Col == c {
							structValue := reflect.ValueOf(&p).Elem()
							structFieldValue := structValue.FieldByName(k)
							if !structFieldValue.IsValid() {
								//log.Printf("No such field: %s in obj", k)
								continue
							}

							if !structFieldValue.CanSet() {
								//log.Printf("Cannot set %s field value", k)
								continue
							}
							structFieldType := structFieldValue.Type() //结构体的类型

							val := reflect.ValueOf(str) //map值的反射值

							var err error
							if structFieldType != val.Type() && str != "" {
								val, err = TypeConversion(fmt.Sprintf("%v", str), structFieldValue.Type().Name()) //类型转换
								if err != nil {
									log.Println(err)
								}
							}
							//log.Println(val)
							structFieldValue.Set(val)
						}
					}

				}
				//log.Printf("row=%d index=%d value=%s", r, c, str)

			}
			//log.Printf(" row =%d %+v", r, p)
			excelContent[r] = p
		}
		log.Printf("%+v", excelContent)
	}
}

//类型转换
func TypeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}

	//else if .......增加其他一些类型的转换

	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}
