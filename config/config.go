package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type tomlConfig struct {
	Server Server
	Viewer Viewer
	System SystemConfig
}

type Server struct {
	Ip   string
	Port string
}

type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Github      string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

func init() {
	// 程序启动时，就会执行init方法
	Cfg = new(tomlConfig)

	Cfg.System.AppName = "zsq-go-blog"
	Cfg.System.Version = 1.0
	currentDir, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	Cfg.System.CurrentDir = currentDir
	_, err = toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		log.Panicln(err)
	}
}
