package BeagleBone-GPIO-Go

/*import (
	"fmt"
	"os"
)*/

type BB_GPIO struct{
	pin [][]int //pin name ex.P8_13 pin[8][13]
	pin_state [][]byte //pin state
}

var pin_map = map[int][]int{
	8:[]int{ 0,0,0,38,39,34,35,66,67,69,68,45,44,23,26,47,46,27,65,22,63,62,37,36,33,32,61,86,88,87,89,10,11,9,81,8,80,78,79,76,77,74,75,72,73,70,71},
	9:[]int{ 0,0,0,0,0,0,0,0,0,0,0,30,60,31,50,48,51,5,4,0,0,3,2,49,15,117,14,115,123,121,122,120,0,0,0,0,0,0,0,0,020,7,0,0,0,0},
}
func BB_GPIO_Start() (gpio *BB_GPIO){
	gpio = new(BB_GPIO)
	gpio.pin_state = make([][]byte,10)
	gpio.pin = make([][]int,10)
	for i:=0;i<10;i++{
		if i<8{
			gpio.pin_state[i] = make([]byte,0,47)
			gpio.pin[i] = make([]int,0,47)
		}else{
			gpio.pin_state[i] = make([]byte,47)
			gpio.pin[i] = make([]int,47)
		}
	}
	for v,i:=range pin_map[8]{
		gpio.pin[8][v] = i
	}
	for v,i:=range pin_map[9]{
		gpio.pin[9][v] = i
	}
	return
}
