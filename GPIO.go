package BeagleBone

import (
	"os"
	"fmt"
)

type BB_GPIO struct{
	Pin [][]int //Pin name ex.P8_13 Pin[8][13]
	Pin_state []byte //Pin state
	check int //Error
	Pin_count uint
}

const (
	OUTPUT = 1
	INPUT = 2
	NONE = 0
)

var Pin_map = map[int][]int{
	8:[]int{ 0,0,0,38,39,34,35,66,67,69,68,45,44,23,26,47,46,27,65,22,63,62,37,36,33,32,61,86,88,87,89,10,11,9,81,8,80,78,79,76,77,74,75,72,73,70,71},
	9:[]int{ 0,0,0,0,0,0,0,0,0,0,0,30,60,31,50,48,51,5,4,0,0,3,2,49,15,117,14,115,123,121,122,120,0,0,0,0,0,0,0,0,020,7,0,0,0,0},
}

func BB_GPIO_Start() (gpio *BB_GPIO){
	gpio = new(BB_GPIO)
	gpio.Pin_state = make([]byte,94)
	gpio.Pin = make([][]int,10)
	for i:=0;i<10;i++{
		if i<8{
			gpio.Pin[i] = make([]int,0,0)
		}else{
			gpio.Pin[i] = make([]int,47)
		}
	}
	for v,i:=range Pin_map[8]{
		gpio.Pin[8][v] = i
	}
	for v,i:=range Pin_map[9]{
		gpio.Pin[9][v] = i
	}
	gpio.Pin_count = 0
	return
}

func (gpio *BB_GPIO) Error() string{
	switch gpio.check{
	case 1:
		return "Pin Error"
	case 2:
		return "Permission Error"
	}
	return "No Error"
}

var mode map[int]string{
	1:"out",
	2:"in",
}

//GPIO num export
func (gpio *BB_GPIO) PinMode(Pin_num int,Pin_mode int) (error){
	if(Pin_num==0 || Pin_num>=123){
		gpio.check = 1
		return gpio
	}
	else if(Pin_mode>0 & Pin_mode<3){
		gpio.check = 1
		return gpio
	}

	export,err := os.Open("/sys/class/gpio/export")
	if err!=nil{
		gpio.check = 2
		return gpio
	}
	defer export.Close()
	fmt.Fprintf(export,"%d",Pin_num)

	direction,err := os.open(fmt.Sprintf("/sys/class/gpio/gpio%d/direction",Pin_num))
	if err != nil{
		gpio.check = 2
		return gpio
	}
	defer direction.CLose()
	fmt.Fprintf(direction,"%s",mode[pin_mode])

	gpio.Pin_state[Pin_count] =  Pin_num
	gpio.Pin_count++

	return nil
}

func (gpio *BB_GPIO) BB_GPIO_Stop{
	unexport,_ := os.Open("/sys/class/gpio/unexport")
	defer unexport.Close()
	for i:=range Pin_state{
		if i!=0{
			fmt.Fprintf(unexport,"%d\n",i)
		}else {
			break
		}
	}
}