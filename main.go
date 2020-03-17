package main

import (
	"os"
	"time"
	"moss/config"

        "github.com/jgarff/rpi_ws281x/golang/ws2811"
	"github.com/wangbokun/go/log"
)

const (
        pin = 12
        count = 85
        brightness = 255
	r =  uint32(0x002000)
	g =  uint32(0x200000)
	b =  uint32(0x000020)

	b1 =  uint32(0x100800)
	b2 =  uint32(0x100008)
	b3 =  uint32(0x081000)
	b4 =  uint32(0x080010)
	b5 =  uint32(0x001008)
	b6 =  uint32(0x000810)
)

var t = []string{"t_0","t_1", "t_2", "t_3", "t_4", "t_5", "t_6", "t_7", "t_8","t_9"}


func colorWipe(num int,color uint32) error{
	ws2811.SetLed(num, color)
	err := ws2811.Render()

	if err!= nil {
		ws2811.Clear()
		return err
	}
	return nil
}


func led(t time.Duration, n []int, color uint32) {
	for _,v := range n {
		err  := colorWipe(v,color)
		if err != nil{
			log.Error("%s",err)
		}
	        time.Sleep(t)
	}
}

func main(){


     err := ws2811.Init(pin, count, brightness)

     if err != nil{
	     log.Error("%s",err)
	     os.Exit(1)
	}

	for  _,v1 := range t {
		n1 := config.ClockFaces(1,v1)
		if n1 !=  nil{
			led(30 * time.Millisecond,n1,b1)
			ws2811.Clear()
		}

		for _,v2 := range t{
			n2 := config.ClockFaces(2,v2)
			log.Info("n1======:",n1)
			log.Info("n2>>>>>>>:",n2)
			if n2 != nil{
				led(5 * time.Millisecond,n1,b2)
				led(30 * time.Millisecond,n2,b3)
				ws2811.Clear()
			}
		}
	}
}
