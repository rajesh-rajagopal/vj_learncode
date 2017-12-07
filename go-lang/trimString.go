package main
import (
 "fmt"
 "regexp"
 "strings"
 //"github.com/pivotal-golang/bytefmt"
)


func main(){
        hdd := "└─vg--one--104-logical_vol5"
	reg, err := regexp.Compile("[^-A-Za-z0-9]+")
	if err != nil {
		fmt.Println(err)
	}
        lvm := strings.TrimSpace(reg.ReplaceAllString(hdd, ""))
	fmt.Println(lvm)
}
