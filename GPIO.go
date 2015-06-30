import (
	"fmt"
	"os"
)

const pin_mapP8={}
const pin_mapP9={}

type BB_GPIO struct{
	var pin map[int][]int //pin name
	var pin_state []byte //pin state
}

func BB_GPIO_Start() (gpio *BB_GPIO){
	gpio = new(BB_GPIO)
}
	
