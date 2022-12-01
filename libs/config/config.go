// Package config 配置加载
package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"transformer_mz/libs/errors"
)

var (
	cPath string
	env   string
	data  []byte
)

// 初始化读取配置文件
func init() {
	var err error
	flag.StringVar(&cPath, "config", "", "transformer_mz -config /etc/transformer_mz.yaml")
	flag.Parse()
	if cPath == "" {
		env = os.Getenv("ENV")
		switch env {
		case "dev", "uat", "pro":
		default:
			env = "dev"
		}
		cPath = "config/" + env + ".yaml"
	}
	data, err = ioutil.ReadFile(cPath)
	if err != nil {
		log.Fatalln(errors.New(err))
	}
}

// Load 加载配置内容
func Load(v interface{}) error {
	return errors.New(yaml.Unmarshal(data, v))
}
