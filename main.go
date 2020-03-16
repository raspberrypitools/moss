package main

import (
        "fmt"
	"time"
	"moss/config"

        "github.com/jgarff/rpi_ws281x/golang/ws2811"
)

const (
        pin = 12
        count = 85
        brightness = 255
	r =  uint32(0x002000)
	g =  uint32(0x200000)
	b =  uint32(0x000020)
)


//var h0 = []int{1, 2, 3, 18, 20, 35, 37, 52, 54, 69, 70, 71}
var t_1_0 = []int{0, 1, 2, 33, 31, 34, 36, 67, 65, 68, 69, 70}
var t_2_0 = []int{4, 5, 6, 29, 27, 38, 40, 63, 61, 72, 73, 74}

var d = []int{25, 59}

var t_3_2 = []int{10, 11, 12, 21, 44, 45, 46, 57, 78, 79, 80}
var t_4_2 = []int{14, 15, 16, 17, 48, 49, 50, 53, 82, 83, 84}


func colorWipe(num int,color uint32) error{
	ws2811.SetLed(num, color)
	err := ws2811.Render()

	if err!= nil {
		ws2811.Clear()
		return err
	}
	return nil
}


func led(n []int,color uint32) {
	for _,v := range n {
		fmt.Print("----",v)
		err  := colorWipe(v,color)
		if err != nil{
			fmt.Println(err)
		}
	        time.Sleep(30 * time.Millisecond)
	}
}

func main(){

     defer ws2811.Fini()
     err := ws2811.Init(pin, count, brightness)

     if err != nil{
	fmt.Println(err)	
     }	else {

        t_1_0 = append(t_1_0, t_2_0...)
        t_1_0 = append(t_1_0, d...)
        t_1_0 = append(t_1_0, t_3_2...)
        t_1_0 = append(t_1_0, t_4_2...)

	led(t_1_0,r)
	time.Sleep(50 * time.Millisecond)
	led(t_1_0,g)
	time.Sleep(50 * time.Millisecond)
	led(t_1_0,b)

	var ti = []string{"t_1_0","t_1_1", "t_1_2", "t_1_3", "t_1_4", "t_1_5", "t_1_6", "t_1_7", "t_1_8","t_1_9"}
	for  _,v := range ti {
		t := config.ClockFace(v)
		led(t,r)
		time.Sleep(50 * time.Millisecond)
	}

     }
}
