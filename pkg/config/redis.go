package config

import (
	"cn.sockstack/gin_demo/pkg/helper"
	"gopkg.in/ini.v1"
)

var Redis *redis

type redis struct {
	Host string `ini:"host"`
	Port int `ini:"port"`

	source *ini.File
}

func (s *redis) Load(path string) *redis {
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

func (s *redis)Init() *redis {
	//判断配置是否加载成功
	if s.source == nil {
		return s
	}

	err := s.source.Section("redis").MapTo(s)
	if err != nil {
		panic(err)
	}
	return s
}
