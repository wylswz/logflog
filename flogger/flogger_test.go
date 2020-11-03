package flogger

import (
	"log"
	"os"
	"testing"
)

func TestInitPos(t *testing.T) {
	f, _ := os.OpenFile("../test.log", os.O_RDONLY, 0666)
	log.Println(initPosition(f))
}
