package main

import (
	B "github.com/tpdns90321/BeagleBone-GPIO-Go/GPIO"
	"fmt"
	"time"
	"os"
)

func main(){
	test := B.BB_GPIO_Start()
	defer test.Close()
	if err := test.PinMode(test.Pin(9,13),B.OUTPUT);err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	for i:=0;i<10;i++{
		test.DigitalWrite(test.Pin(9,13),B.LOW)
		time.Sleep(time.Second)
		test.DigitalWrite(test.Pin(9,13),B.HIGH)
		time.Sleep(time.Second)
	}
}
