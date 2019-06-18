/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-09 16:50:24
LastEditors:
LastEditTime: 2019-06-09 16:50:24
Description:  对配置文件的解析
*/
package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

/**
通过viper变量的方式来获得数据
import "scaffold_go/conf"
vp := conf.getViper()
name := vp.Get("name")
*/
var Vp *viper.Viper

func init() {
	Vp = viper.New()

	Vp.SetEnvPrefix("") // will be uppercased automatically
	Vp.BindEnv("environment")
	//fmt.Println(Vp.GetString("environment"))

	Vp.SetConfigName(CONF_FILE) // 设定配置文件的名称（不包括后缀）
	for _, path := range CONF_SEARCH_PATH {
		//fmt.Println(path)
		Vp.AddConfigPath(path)
	}
	Vp.SetConfigType(CONF_FILE_TYPE) // 设定配置文件的格式： yaml
	// 设置监听文件的变更
	if CONF_WATCHING {
		Vp.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Printf("Config file changed: %s", e.Name)
		})
	}
	err := Vp.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Vp.Unmarshal(&Cfg)
	Cfg.Environment = Vp.GetString("environment")
}

func getViper() (vp *viper.Viper) {
	Vp.SetConfigName(CONF_FILE) // 设定配置文件的名称（不包括后缀）
	for _, path := range CONF_SEARCH_PATH {
		fmt.Println(path)
		Vp.AddConfigPath(path)
	}
	Vp.SetConfigType(CONF_FILE_TYPE) // 设定配置文件的格式： yaml
	// 设置监听文件的变更
	if CONF_WATCHING {
		Vp.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("Config file changed:", e.Name)
		})
	}
	err := Vp.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return nil
	} else {
		return Vp
	}
}

/**
通过反序列化的方式直接变为一个结构体的对象
import "scaffold_go/conf"
conf.NewVp()
fmt.Println(conf.Cfg.Name)
*/
var Cfg CfgStruct = CfgStruct{}

// 将配置文件转换为一个结构体的变量
func NewVp() {
	Vp.Unmarshal(&Cfg)
	//fmt.Println(Cfg)
	Cfg.Environment = Vp.GetString("environment")
}

// 从配置文件获得Iris的配置信息
func InitIrisConfiguration() (configuration iris.Configuration) {
	environment := Cfg.Environment
	cfgIris := Cfg.Iris[environment]
	defaultConfiguration := iris.DefaultConfiguration()
	defaultConfiguration.DisableStartupLog = cfgIris.DisableStartupLog
	defaultConfiguration.DisableInterruptHandler = cfgIris.DisableInterruptHandler
	defaultConfiguration.DisablePathCorrection = cfgIris.DisablePathCorrection
	defaultConfiguration.EnablePathEscape = cfgIris.EnablePathEscape
	defaultConfiguration.FireMethodNotAllowed = cfgIris.FireMethodNotAllowed
	defaultConfiguration.DisableBodyConsumptionOnUnmarshal = cfgIris.DisableBodyConsumptionOnUnmarshal
	defaultConfiguration.DisableAutoFireStatusCode = cfgIris.DisableAutoFireStatusCode
	defaultConfiguration.TimeFormat = cfgIris.TimeFormat
	defaultConfiguration.Charset = cfgIris.Charset
	defaultConfiguration.PostMaxMemory = cfgIris.PostMaxMemory
	defaultConfiguration.TranslateFunctionContextKey = cfgIris.TranslateFunctionContextKey
	defaultConfiguration.ViewLayoutContextKey = cfgIris.ViewLayoutContextKey
	defaultConfiguration.ViewDataContextKey = cfgIris.ViewDataContextKey
	defaultConfiguration.EnableOptimizations = cfgIris.EnableOptimizations
	return defaultConfiguration
}
