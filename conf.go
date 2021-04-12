package cstools

import (
	"github.com/Unknwon/goconfig"
)

const ConfPosition = "conf.ini"

func GetConf(section string, key string) (result string) {
	cfg, err := goconfig.LoadConfigFile(ConfPosition)
	if err != nil {
		return "配置文件读取错误"
	}
	result, err = cfg.GetValue(section, key)
	if err != nil {
		return ""
	} else {
		return result
	}
}
