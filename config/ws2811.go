package config

import (
	"time"
//	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)



const (
	sleepTime  = 50
)

type wsEngine interface {
	Init() error
	Render() error
	Wait() error
	Fini()
	Leds(channel int) []uint32
}


type ColorWipe struct {
	Ws wsEngine
}


func (cw *ColorWipe) Setup() error {
	return cw.Ws.Init()
}


func (cw *ColorWipe) Display(color uint32) error {
	for i := 0; i < len(cw.Ws.Leds(0)); i++ {
		cw.Ws.Leds(0)[i] = color
		if err := cw.Ws.Render(); err != nil {
			return err
		}
		time.Sleep(sleepTime * time.Millisecond)
	}
	return nil
}
