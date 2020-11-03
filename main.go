package main

import (
	"os"

	"xmbsmdsj.com/logflog/flogger"
)

func main() {
	flogger.FLog(os.Args[1:])
}
