package main

import (
	"embed"
	"main/service"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	service.PRODUCTION = true // 生产环境模式 默认开启
	go service.LoadConfig("1.1.6")
	go service.InitSoundFont("./assets/Yamaha-Grand-Lite-v2.0.sf2")
	go service.ListenDevices()
	service.Run(assets)
}
