package main

import (
	B "BeagleBone"
	"fmt"
)

func main(){
	test := B.BB_GPIO_Start()
	fmt.Println(*test)
	fmt.Println(test.PinMode(test.Pin(8,12),1))
}
