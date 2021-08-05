package main

import (
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/heavenlwf/tookit/utils"
)

var (
	isSetClipboard bool
	originString   string
)

func parseArgs() {
	flag.BoolVar(&isSetClipboard, "c", false, "set clipboard")
	flag.StringVar(&originString, "s", "", "origin string")
	flag.Parse()
}

func main() {
	parseArgs()
	decodeByte, err := base64.StdEncoding.DecodeString(originString)
	if err != nil {
		panic(err)
	}
	if isSetClipboard {
		if err := utils.SetClipboard(string(decodeByte)); err != nil {
			panic(err)
		}
	}
	fmt.Println(string(decodeByte))
}
