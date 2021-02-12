package main

import (
	"2fa/config"
	"2fa/web"
)

func main() {
	conf := config.ReadConfig()
	web.Init(conf)
}
