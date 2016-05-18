package main

import (
	"flag"
	"fmt"

	"github.com/k0kubun/pp"
)

var version string
var show_version = flag.Bool("v", false, "show version")

func main() {
	flag.Parse()
	if *show_version {
		fmt.Printf("version: %s\n", version)
		return
	}
	pp.Println("hello world")
}
