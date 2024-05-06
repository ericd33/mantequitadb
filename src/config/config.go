package config

import "fmt"

var Debug bool = true

func DebugIt(msg string) {
	if Debug {
		fmt.Println(msg)
	}

}
