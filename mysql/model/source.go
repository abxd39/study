package model

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	ConfigPath string
)

var (
	TestDb *xorm.Engine
	Con    *Config
)

func NewSession() error {
	if Con == nil {
		err:=getConfig()
		return err
	}
	master, err := xorm.NewEngine("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			Con.User,
			Con.Pass,
			Con.Host,
			Con.Port,
			Con.Db))
	if err != nil {
		return err
	}
	TestDb = master
	return nil
}

func getConfig()error {
	err := LoadYaml(getConfigFilePath(), Con)
	if err != nil {
		return err
	}

	fmt.Printf("config=%v", Con)

	return nil
}

func getConfigFilePath() string {
	if ConfigPath == "" {
		ConfigPath = filepath.Join(GetCurrentDirectory(), "config.yml")
	}
	return ConfigPath
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Panicln(err)
	}
	path := dir
	log.Printf("处理前%v", path)
	path = filepath.FromSlash(dir)
	log.Printf("处理后%v", path)
	return path
}

// LoadYaml is used to load yaml file
func LoadYaml(path string, data *Config)  error {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		return err
	}
	fmt.Printf("path=%v\r\n", path)
	body, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		return  err
	}
	data1:=new(Config)
	err = json.Unmarshal(body, data1)
	if err != nil {
		fmt.Println("3")
		fmt.Println(err)
		return  err
	}

	if data1.User==""{
		err=fmt.Errorf("配置文件加载失败 000")
		return err
	}
	fmt.Println(data1)
	data =data1
	return nil
}

//read yaml config
//注：path为yaml或yml文件的路径
func ReadYamlConfig(path string) error {
	if f, err := os.Open(path); err != nil {
		return err
	} else {
		err = yaml.NewDecoder(f).Decode(Con)
		if err != nil {
			return err
		}
	}
	return nil
}
