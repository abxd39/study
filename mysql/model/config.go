package model

import (
	"fmt"
)

type Config struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
	Db      string `yaml:"db"`
	Charset string `yaml:"charset"`
}


func (m*Config) String()string{
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",m.User,m.Pass,m.Host,m.Port,m.Db,m.Charset)
}

