package main

import (
	"fmt"
	"github.com/krishpranav/gostocks/version"
)

func main() {
	version := true

	if version == true {
		fmt.Println(version.Version)		
		return
	}
	
}