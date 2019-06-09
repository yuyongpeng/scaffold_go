package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)
var v = viper.New()

func main() {
	v.SetConfigName("viper")     	// name of config file (without extension)
	v.AddConfigPath("conf")  			// path to look for the config file in
	v.AddConfigPath("$HOME/.appname") 	// call multiple times to add many search paths
	v.AddConfigPath(".")              	// optionally look for config in the working director
	v.AddConfigPath("$HOME/")
	v.SetConfigType("yaml") 				// yaml
	err := v.ReadInConfig()           		// Find and read the config file
	if err != nil { 							// Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(v.Get("name"))
	v.SetDefault("xxx", "aaa")
	v.Get("Hacker")
	fmt.Println(v.Get("xxx"))

	// 环境变量自动前缀  FOO_BAR=xx go run main/testViper.go
	v.AutomaticEnv()
	v.SetEnvPrefix("foo")				// 自动转为大写
	v.BindEnv("bar")					// ENV是区分大小写的，这里只放一个参数
	v.SetDefault("bar","arg1")
	fmt.Println(v.Get("bar"))
	//
	v.AutomaticEnv()
	v.BindEnv("id", "PARAM_ARG")		// 指定绑定的环境变量。区分大小写
	v.SetDefault("id","argx")
	fmt.Println(v.Get("id"))
	//
	v.AutomaticEnv()
	os.Setenv("REFRESH_INTERVAL", "30s")
	replacer := strings.NewReplacer("-", "_")
	v.SetEnvKeyReplacer(replacer)
	fmt.Print(v.Get("refresh-interval")) 	// 将变量的"-"调换为"_"
}
