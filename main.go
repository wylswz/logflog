package main

import (
	"os"

	"github.com/wylswz/logflog/flogger"
)

func main() {
	flogger.FLog(os.Args[1:])
}
