package main

import (
	"./getBattles"
	"./getSalmons"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	mode = kingpin.Flag("mode", "Set verbose mode").Short('m').String()
	next = kingpin.Flag("next", "next terms").Short('n').Bool()
)

func main() {
	// var mode string
	kingpin.Version("0.3.0")
	// kingpin.Usage("Splatoon2のステージ情報を出力するよ")
	kingpin.Parse()

	if *mode == "salmon" {
		getSalmons.GetSalmons(*next)
	} else {
		getBattles.GetBattles(*next)
	}
}
