package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"./salmon"
  "./battles"
)

func main() {
	var mode string
	app := cli.NewApp()

	app.Name = "ika2cli"
	app.Usage = "Splatoon2のステージ情報を出力します"
	app.Version = "0.3.0"

	app.Action = func(context *cli.Context) error {
		if mode == "salmon" {
			salmon.GetSalmons(context)
		} else {
			battles.GetBattles(context)
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "mode, m",
			Value:       "regular",
			Usage:       "Battles mode or Salmon-run mode",
			Destination: &mode,
		},
		cli.BoolFlag{
			Name:  "next, n",
			Usage: "to show next term",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
