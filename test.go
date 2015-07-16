package main

import (
	B "github.com/tpdns90321/BeagleBone-GPIO-Go/GPIO"
	"fmt"
	"time"
	"os"
)

func main(){
	gpio := B.BB_GPIO_Start()
	defer gpio.Close()
	led := gpio.Pin(9,13)
	if err := gpio.PinMode(led,B.OUTPUT);err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	for i:=0;i<10;i++{
		gpio.DigitalWrite(led,B.LOW)
		time.Sleep(time.Second)
		gpio.DigitalWrite(led,B.HIGH)
		time.Sleep(time.Second)
	}
}
