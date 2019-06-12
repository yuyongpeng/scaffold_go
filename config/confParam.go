/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-10 11:20:04
LastEditors:
LastEditTime: 2019-06-10 11:20:04
Description:
*/
package config

// 项目的名称
var APP_NAME = "SCAFFOLD"

// 配置文件的文件名，不包括后缀
var CONF_FILE = "scaffold"
// 配置文件的类型
var CONF_FILE_TYPE = "yaml"
// 配置文件的搜索路径
var CONF_SEARCH_PATH = []string{"conf", ".", "$HOME/", "../conf"}
// 是否监听配置文件的变更
var CONF_WATCHING = false

// 配置文件的结构体
// 配置文件和这个结构体是一一对应的，配置文件有修改，只需要对应的变更这个结构体即可
type CfgStruct struct {
	Hacker      bool             `mapstructure:"Hacker"`
	Name        string           `mapstructure:"name"`
	Hhobbies    []string         `mapstructure:"hobbies"`
	Clothing    interface{}      `mapstructure:"clothing"`
	Age         int              `mapstructure:"age"`
	Eyes        string           `mapstructure:"eyes"`
	Beard       bool             `mapstructure:"beard"`
	//////////////////////////////////////////////////////////
	Environment string           `mapstructure:"environment"`
	Iris        Iris             `mapstructure:"iris"`
	Database    map[string]Mysql `mapstructure:"test"`
}
type Mysql struct {
	Username        string `mapstructure:"userame"`
	Password        string `mapstructure:"password"`
	Ip              string `mapstructure:"ip"`
	Port            string `mapstructure:"port"`
	Database        string `mapstructure:"database"`
	Param           string `mapstructure:"param"`
	Maxidleconns    int    `mapstructure:"maxidleconns"`
	Maxopenconns    int    `mapstructure:"maxopenconns"`
	Connmaxlifetime int64  `mapstructure:"connmaxlifetime"`
}

type Iris struct {
	Html                              string `mapstructure:"html"`
	Upload                            string `mapstructure:"upload"`
	DisableStartupLog                 bool   `mapstructure:"DisableStartupLog"`
	DisableInterruptHandler           bool   `mapstructure:"DisableInterruptHandler"`
	DisablePathCorrection             bool   `mapstructure:"DisablePathCorrection"`
	EnablePathEscape                  bool   `mapstructure:"EnablePathEscape"`
	FireMethodNotAllowed              bool   `mapstructure:"FireMethodNotAllowed"`
	DisableBodyConsumptionOnUnmarshal bool   `mapstructure:"DisableBodyConsumptionOnUnmarshal"`
	DisableAutoFireStatusCode         bool   `mapstructure:"DisableAutoFireStatusCode"`
	TimeFormat                        string `mapstructure:"TimeFormat"`
	Charset                           string `mapstructure:"Charset"`
	PostMaxMemory                     int64  `mapstructure:"PostMaxMemory"`
	TranslateFunctionContextKey       string `mapstructure:"TranslateFunctionContextKey"`
	ViewLayoutContextKey              string `mapstructure:"ViewLayoutContextKey"`
	ViewDataContextKey                string `mapstructure:"ViewDataContextKey"`
	EnableOptimizations               bool   `mapstructure:"EnableOptimizations"`
}

// 日志的格式化输出 "json" 和 "text", 填写错误,默认使用json
var LOG_FORMATER = "text"
// 日志的输出 "console(os.Stderr)" "file", 填写错误,默认使用console
var LOG_OUTPUT = "console"
// 文件日志的路径
var LOG_OUTPUT_FILE = "/tmp/" + APP_NAME + ".log"
// 设置日志的级别 "trace" "debug" "info" "warn" "error" "fatal" "panic", 填写错误,默认使用info
var LOG_LEVEL = "debug"
