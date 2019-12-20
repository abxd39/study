package mq

import (
	"flag"
	"github.com/koding/multiconfig"
	"os"
)

type MqFlagConfig struct {
	ConfigFile string `json:"config_file" default:"rmq.json"`
}

//连接结构
type Connect struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

//信道结构
type Channel struct {
	Name     string `json:"name"`
	Connect  string `json:"connect"`
	QosCount int    `json:"qos_count"`
	QosSize  int    `json:"qos_size"`
}

//交换机绑定结构
type EBind struct {
	Destination string `json:"destination"`
	Key         string `json:"key"`
	NoWait      bool   `json:"no_wait"`
}

//交换机结构
type Exchange struct {
	Name        string                 `json:"name"`
	Channel     string                 `json:"channel"`
	Type        string                 `json:"type"`
	Durable     bool                   `json:"durable"`
	AutoDeleted bool                   `json:"auto_deleted"`
	Internal    bool                   `json:"internal"`
	NoWait      bool                   `json:"no_wait"`
	Bind        []EBind                `json:"ebind"`
	Args        map[string]interface{} `json:"args"`
}

//队列绑定结构
type QBind struct {
	ExchangeName string `json:"exchange_name"`
	Key          string `json:"key"`
	NoWait       bool   `json:"no_wait"`
}

//队列结构
type Queue struct {
	Name       string                 `json:"name"`
	Channel    string                 `json:"channel"`
	Durable    bool                   `json:"durable"`
	AutoDelete bool                   `json:"auto_delete"`
	Exclusive  bool                   `json:"exclusive"`
	NoWait     bool                   `json:"no_wait"`
	Bind       []QBind                `json:"qbind"`
	Args       map[string]interface{} `json:"args"`
}

//发送者配置
type Pusher struct {
	Name         string `json:"name"`
	Channel      string `json:"channel"`
	Exchange     string `json:"exchange"`
	Key          string `json:"key"`
	Mandatory    bool   `json:"mandatory"`
	Immediate    bool   `json:"immediate"`
	ContentType  string `json:"content_type"`
	DeliveryMode uint8  `json:"delivery_mode"`
}

//接收者配置
type Popup struct {
	Name      string `json:"name"`
	QName     string `json:"q_name"`
	Channel   string `json:"channel"`
	Consumer  string `json:"consumer"`
	AutoACK   bool   `json:"auto_ack"`
	Exclusive bool   `json:"exclusive"`
	NoLocal   bool   `json:"no_local"`
	NoWait    bool   `json:"no_wait"`
}

//配置文件结构
type mqCfg struct {
	Connects  []Connect  `json:"connects"`
	Channels  []Channel  `json:"channels"`
	Exchanges []Exchange `json:"exchanges"`
	Queue     []Queue    `json:"queue"`
	Pusher    []Pusher   `json:"pusher"`
	Popup     []Popup    `json:"popup"`
}

var Cfg *mqCfg = new(mqCfg)

//读取配置文件
func LoadCfg() (err error) {
	if err = Cfg.load(); err != nil {
		return err
	}
	return nil
}

func (c *MqFlagConfig) load() error {
	t := &multiconfig.TagLoader{}
	f := &multiconfig.FlagLoader{}
	m := multiconfig.MultiLoader(t, f)
	if err := m.Load(c); err == flag.ErrHelp {
		os.Exit(0)
	} else if err != nil {
		return err
	}
	return nil
}

func (c *mqCfg) load() error {
	f := &MqFlagConfig{}
	err := f.load()
	if err == flag.ErrHelp {
		os.Exit(0)
	} else if err != nil {
		return err
	}
	t := &multiconfig.TagLoader{}
	j := &multiconfig.JSONLoader{Path: f.ConfigFile}
	m := multiconfig.MultiLoader(t, j)
	err = m.Load(c)
	return err
}