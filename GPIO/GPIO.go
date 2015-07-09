package BeagleBone_GPIO

import (
        "fmt"
        "os"
)

type BB_GPIO struct {
        pin       [][]int  //Pin name ex.P8_13 Pin(9,12)
        pin_state [][]byte //Pin state
        check     int      //Error
}

//pin location store
type pin_data struct {
        num_pin    int
        beagle_pin []int
}

//const allocate
const (
        OUTPUT = 1
        INPUT  = 2
        NONE   = 0
        HIGH   = 1
        LOW    = 0
)

//mode set
var mode = map[int]string{
        1: "out",
        2: "in",
}

//pin allocate
var pin_map = map[int][]int{
        8: []int{0, 0, 0, 38, 39, 34, 35, 66, 67, 69, 68, 45, 44, 23, 26, 47, 46, 27, 65, 22, 63, 62, 37, 36, 33, 32, 61, 86, 88, 87, 89, 10, 11, 9, 81, 8, 80, 78, 79, 76, 77, 74, 75, 72, 73, 70, 71},
        9: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 30, 60, 31, 50, 48, 51, 5, 4, 0, 0, 3, 2, 49, 15, 117, 14, 115, 123, 121, 122, 120, 0, 0, 0, 0, 0, 0, 0, 0, 20, 7, 0, 0, 0, 0},
}

//pin numbers initialize
func BB_GPIO_Start() (gpio *BB_GPIO) {
        gpio = new(BB_GPIO)
        gpio.pin = make([][]int, 10)
        gpio.pin_state = make([][]byte, 10)
        for i := 0; i < 10; i++ {
                if i < 8 {
                        gpio.pin[i] = make([]int, 0, 47)
                        gpio.pin_state[i] = make([]byte, 0, 47)
                } else {
                        gpio.pin[i] = make([]int, 47)
                        for t, v := range pin_map[i] {
                                gpio.pin[i][t] = v
                        }
                        gpio.pin_state[i] = make([]byte, 47)
                }
        }
        return
}

//pin_data return function, and it's error check
func (gpio *BB_GPIO) Pin(fir, sec int) (data *pin_data) {
        data = new(pin_data)
        data.num_pin = gpio.pin[fir][sec]
        if data.num_pin == 0 || data.num_pin > 123 {
                gpio.check = 1
                return nil
        }
        data.beagle_pin = make([]int, 2)
        t := []int{fir, sec}
        for i, v := range t {
                data.beagle_pin[i] = v
        }
        gpio.check = 0
        return
}

//setting GPIO pin
func (gpio *BB_GPIO) PinMode(data *pin_data, mode_pin int) error {
        if data == nil {
                return gpio
        }

        export, err := os.OpenFile("/sys/class/gpio/export",os.O_WRONLY | os.O_APPEND,0200)
        if err != nil {
                return err
        }

        defer export.Close()
        fmt.Fprintf(export, "%d",data.num_pin)

        direction,err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/direction",data.num_pin),os.O_WRONLY,0311)
        if err != nil {
                gpio.check=1
                return err
        }
        defer direction.Close()

        fmt.Fprintf(direction,"%s",mode[mode_pin])

        gpio.pin_state[data.beagle_pin[0]][data.beagle_pin[1]] = byte(mode_pin)

        return nil
}

var signal = map[int]string{
        1:"high",
        0:"low",
}

//GPIO on/off
func (gpio *BB_GPIO) DigitalWrite(data *pin_data, on int) error {
        if data == nil {
                return gpio
        }else if gpio.pin_state[data.beagle_pin[0]][data.beagle_pin[1]]!=1{
                gpio.check = 1
                return gpio
        }

        value,err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/direction", data.num_pin),os.O_RDWR,0311)
        if err!=nil {
                return err
        }
        defer value.Close()

        fmt.Fprintf(value,signal[on])

        return nil
}

//read GPIO pin's bit
func (gpio *BB_GPIO) DigitalRead(pin *pin_data) (data int,err error){
        if pin == nil{
                return -1,gpio
        }else if gpio.pin_state[pin.beagle_pin[0]][pin.beagle_pin[1]] != 2{
                gpio.check = 1
                return -1,gpio
        }

        value,err := os.OpenFile(fmt.Sprintf("/sys/class/gpio/gpio%d/value",data.num_pin),os.O_RDONLY,0311)
        if err != nil{
                return -1,err
        }
        read := make([]byte,1,2)
        value.Read(read)
        data = int(read[0]-48)
        err = nil
        return
}

//for error return
func (gpio *BB_GPIO) Error() string {
        switch gpio.check {
        case 1:
                return "Pin Error"
        }
        return ""
}

//cleanup pin settings
func (gpio *BB_GPIO) Close() error{
        unexport,err := os.OpenFile("/sys/class/gpio/unexport",os.O_WRONLY | os.O_APPEND,0200)
        if err!=nil{
                return err
        }
        defer unexport.Close()
        for i:=8;i<10;i++{
                for v := range gpio.pin_state[i]{
                        if v != 0{
                                fmt.Fprintf(unexport,"%d\n",v)
                        }
                }
        }
        return nil
}