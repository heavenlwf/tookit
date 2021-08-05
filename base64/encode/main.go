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
	encodeStr := base64.StdEncoding.EncodeToString([]byte(originString))
	if isSetClipboard {
		if err := utils.SetClipboard(encodeStr); err != nil {
			panic(err)
		}
	}
	fmt.Println(encodeStr)
}
