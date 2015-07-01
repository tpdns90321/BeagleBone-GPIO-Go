package BeagleBone

import (
	"os"
	"fmt"
	"io"
)

type BB_GPIO struct{
	pin [][]int //Pin name ex.P8_13 Pin(9,12)
	pin_state [][]byte //Pin state
	check int //Error
}

const (
	OUTPUT = 1
	INPUT = 2
	NONE = 0
	HIGH = 1
	LOW = 0
)

var mode = map[int]string{
	1:"out",
	2:"in",
}

var Pin_map = map[int][]int{
	8:[]int{ 0,0,0,38,39,34,35,66,67,69,68,45,44,23,26,47,46,27,65,22,63,62,37,36,33,32,61,86,88,87,89,10,11,9,81,8,80,78,79,76,77,74,75,72,73,70,71},
	9:[]int{ 0,0,0,0,0,0,0,0,0,0,0,30,60,31,50,48,51,5,4,0,0,3,2,49,15,117,14,115,123,121,122,120,0,0,0,0,0,0,0,0,020,7,0,0,0,0},
}

func BB_GPIO_Start() (gpio *BB_GPIO){
	gpio = new(BB_GPIO)
	gpio.pin = make([][]int,10)
	gpio.pin_state = make([][]byte,10)
	for i:=0;i<10;i++{
		if i<8{
			gpio.pin[i] = make([]int,0,47)
			gpio.pin_state[i] = make([]byte,0,47)
		}else{
			gpio.pin[i] = make([]int,47)
			gpio.pin_state[i] = make([]byte,47)
		}
	}
	return
}

func (gpio *BB_GPIO) Pin(fir int, sec int) (func() (int,map[int]int,error)){
	v := gpio.pin[fir][sec]
	var Pin = map[int]int{
		1:fir,
		2:sec,
	}
	return func() (int,map[int]int,error){
		if v==0 || v >= 123 { //exception exclude
			gpio.check = 1
			return v,Pin,gpio
		}
		return v,Pin,nil
	}
}

func (gpio *BB_GPIO) PinMode(pin func() (int,map[int]int,error),M int) error{
	num_pin,beagle_pin,error := pin()
	if error!=nil{
		return error
	}else if M<3 || M==0{
		gpio.check = 1
		return gpio
	}

	export,error := os.Create("/sys/class/gpio/export")
	if error != nil{
		return error
	}
	defer export.Close()
	io.WriteString(export,fmt.Sprintf("%d",num_pin))
	direction,error := os.Create(fmt.Sprintf("/sys/class/gpio/gpio%d/direction",num_pin))
	if error != nil{
		return error
	}
	defer direction.Close()
	_,error = io.WriteString(direction,mode[M])
	if error!=nil{
		return error
	}
	gpio.pin_state[beagle_pin[1]][beagle_pin[2]] = byte(M)
	return nil
}

func (gpio *BB_GPIO) DigitalWrite(pin func() (int,map[int]int,error),state int) error{
	num_pin,beagle_pin,error := pin()
	if error!=nil{
		return error
	}
	if(gpio.pin_state[beagle_pin[1]][beagle_pin[2]] == 1){
		value,error := os.Create(fmt.Sprintf("/sys/class/gpio/gpio%d/value",num_pin))
		if error!=nil{
			return error
		}
		defer value.Close()
		io.WriteString(value,string(state))
		return nil
	}
	gpio.check = 1
	return gpio
}

func (gpio *BB_GPIO) Error() string{
	switch gpio.check{
	case 1:
		return "Pin Error"
	}
	return "No Error"
}
