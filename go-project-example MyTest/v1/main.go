package main

import (
	"os"
	"v1/Common"
	"v1/Repository"
	"v1/Router"
	"v1/Util"
)

func main() {
	// 数据初始化
	if err := Init(); err != nil {
		os.Exit(Common.CodeError)
	}

	Router.Start()
}

func Init() error {
	if err := Repository.InitIndex(); err != nil {
		return err
	}
	if err := Util.InitLogger(); err != nil {
		return err
	}
	return nil
}
