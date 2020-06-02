package config

import (
	"cn.sockstack/gin_demo/pkg/helper"
	"gopkg.in/ini.v1"
)

var App *app

type app struct {
	Mode string `ini:"mode"`

	source *ini.File
}

func (s *app) Load(path string) *app {
	var err error
	//判断配置文件是否存在
	exists, err := helper.PathExists(path)
	if !exists {
		return s
	}
	s.source, err = ini.Load(path)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *app)Init() *app {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}

	err := s.source.MapTo(s)
	if err != nil {
		panic(err)
	}
	return s
}
