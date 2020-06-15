package main

import (
	"github.com/v39lfy/go-module"
)

func main()  {
	hub := module.NewModuleHub()
	hub.RunLoop()
}