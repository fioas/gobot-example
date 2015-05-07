package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	// Firmataプロトコルを使用。ポート指定はArudinoを接続しているホストコンピュータの環境に依存
	firmataAdaptor := firmata.NewFirmataAdaptor("arduino", "/dev/tty.usbmodem1411")
	// Arduinoのデジタル出力ピン番号を指定
	led := gpio.NewLedDriver(firmataAdaptor, "led", "13")

	// 1秒間隔でLEDを明滅
	work := func() {
		gobot.Every(1*time.Second, func() {
			// 指定ピンに対して交互にHIGH/LOWを出力
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
