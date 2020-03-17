package main

import (
	"os"
	"moss/config"
	"github.com/wangbokun/go/log"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

func checkError(err error){
        if err != nil {
		log.Debug("%#v",err)
                panic(err)
        }
}

func main() {
	err  := config.Init("./etc/moss.yaml")
	if err != nil{
		log.Error("init config error %s ,exit",err)
		os.Exit(1)
	}
	log.Debug("%#v",config.Conf)

	opt := ws2811.DefaultOptions
	opt.Channels[0].Brightness = config.Conf.Led.Brightness
	opt.Channels[0].LedCount = config.Conf.Led.Count

	dev, err := ws2811.MakeWS2811(&opt)
	checkError(err)

	cw := &config.ColorWipe{
		Ws: dev,
	}
	checkError(cw.Setup())
	defer dev.Fini()

	cw.Display(uint32(0x0000ff))
	cw.Display(uint32(0x00ff00))
	cw.Display(uint32(0xff0000))
	cw.Display(uint32(0x000000))
}
