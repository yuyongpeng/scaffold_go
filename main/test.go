//package main
//
//import (
//	"fmt"
//	home "github.com/mitchellh/go-homedir"
//	"log"
//)
//
//func main() {
//	home1, _ := home.Dir()	// "/Users/yuyongpeng"
//	fmt.Println(home1)
//	home2, _ := home.Expand("~/go1114") 	// "/Users/yuyongpeng/go1114"
//	fmt.Println(home2)
//}
package main

import (
	"regexp"
	"strings"
)
import "fmt"
import "reflect"

func main(){
	re := regexp.MustCompile("[0-9]*([a-z]+)[0-9]*([a-z]+)")
	ret := re.FindAllStringSubmatch("09hello0809hello", -1)
	fmt.Println(ret)

	reg := regexp.MustCompile(`(?P<aa>\w)(\w)+`)
	rt := reg.FindAllStringSubmatch("Hello World!", -1)
	fmt.Printf("%q", rt)  // [["Hello" "H" "o"] ["World" "W" "d"]]

	value := 32 << 20
	fmt.Println(reflect.TypeOf(value))

	fmt.Println(strings.ToLower("aaaAAAA"))

}



