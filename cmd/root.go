package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var author string
var config string
var cfgFile string
var file string
var license string
//var viper string

var name string
var age string



var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "cobra [command]",
	Long: `cobra [command]2`,
	Args: cobra.MinimumNArgs(1),		// cobra命令 最少必须要有1个参数
	// PreRun 之前运行的内容 (在运行子命令的时候，也会出现)
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	},
	// Run之前运行的内容
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	//这个命令的处理函数，获得参数进行处理
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(strings.Join(args, " "))
		fmt.Println(file)
	},
	// Run 之后运行的内容
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	},
	// PostRun 之后运行的内容 (在运行子命令的时候，也会出现)
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
}
var addCmd = &cobra.Command{
	Use: "add",
	Short: "添加姓名和年龄",
	Long: "用于添加用户的姓名和年龄到内存中",
	Args: cobra.MinimumNArgs(1),		// 允许的最小参数个数
	Run: func(cmd *cobra.Command, args []string){
		//cmd.
	},
}

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize a Cobra Application",
	Long: "Initialize a Cobra Application Long",
	Args: cobra.MaximumNArgs(2),		// 最多只能有2个参数
	Run: func(cmd *cobra.Command, args []string){
		for k,v := range args{
			fmt.Print("%s,%s", k, v)
		}
	},
}

func init() {
	//cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)

	rootCmd.Flags().StringVarP(&author, "author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.Flags().StringVar(&config, "config", "$HOME/.cobra.yaml", "config file ")
	rootCmd.Flags().StringVarP(&license, "license", "l", "", "name of license for the project")

	addCmd.Flags().StringVarP(&name, "name", "n", "no name", "添加用户的姓名")
	addCmd.Flags().StringVarP(&age, "age", "g", "0", "添加用户的年龄")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "apache")
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func Execute() {
	//home, _ := homedir.Dir()
	//fmt.Println(home)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

