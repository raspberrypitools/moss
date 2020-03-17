package config


import (
	"github.com/wangbokun/go/log"
	"github.com/wangbokun/go/types"
)


type Config struct{
	Color   `yaml:"color"`
	Led	`yaml:"led"`
	Pi	`yaml:"pi"`
}


type Color struct{
	Red	uint32	`yaml:"red"`
	Green	uint32	`yaml:"green"`
	Blue	uint32	`yaml:"blue"`
}




type Led struct{
	Count	int	`yaml:"count"`
	Brightness int	`yaml:"brightness"`
}


type Pi struct{
	Pin	int	`yaml:"pin"`
}



var Conf *Config = &Config{}

func Init(filename string) error {
	log.Debug("%#v",Conf)
	return types.ParseConfigFile(Conf, filename)
}
