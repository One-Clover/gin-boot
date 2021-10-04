package boot

import (
	"github.com/go-yaml/yaml"
	"log"
)

type ServerConfig struct {
	Port int32
	Name string
	Html string
}

type SysConfig struct {
	Server *ServerConfig
	Config CustomizeConfig
}

func (this *SysConfig) Name() string {
	return "SysConfig"
}

// CustomizeConfig 自定义config
type CustomizeConfig map[interface{}]interface{}

func NewSysConfig() *SysConfig {
	return &SysConfig{Server: &ServerConfig{Port: 8080, Name: "gin-boot"}}
}

func InitConfig() *SysConfig {
	config := NewSysConfig()
	if b := LocalConfigFile(); b != nil {
		err := yaml.Unmarshal(b, config)
		if err != nil {
			log.Fatal(err)
		}
	}
	return config
}

func GetConfigValue(m CustomizeConfig, prefix []string, index int) interface{} {
	key := prefix[index]
	if v, ok := m[key]; ok {
		if index == len(prefix)-1 {
			return v
		} else {
			index = index + 1
			if mv, ok := v.(CustomizeConfig); ok {
				return GetConfigValue(mv, prefix, index)
			} else {
				return nil
			}
		}
	}
	return nil
}
