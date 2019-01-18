package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

type Ika struct {
	Result struct {
		Regular []struct {
			Maps   []string `json:"maps"`
			StartT int64    `json:"start_t"`
			EndT   int64    `json:"end_t"`
		} `json:"regular"`
		Gachi []struct {
			Rule string   `json:"rule"`
			Maps []string `json:"maps"`
		} `json:"gachi"`
		League []struct {
			Rule string   `json:"rule"`
			Maps []string `json:"maps"`
		} `json:"league"`
	} `json:"result"`
}

type Salmon struct {
	Result []struct {
		StartT int64 `json:"start_t"`
		EndT   int64 `json:"end_t"`
		Stage  struct {
			Name string `json:"name"`
		} `json:"stage"`
		Weapons []struct {
			Name string `json:"name"`
		} `json:"weapons"`
	} `json:"result"`
}

func GetSalmons(context *cli.Context) {
	resp, err := http.Get("https://spla2.yuu26.com/coop/schedule")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var salmon Salmon
	if err := json.NewDecoder(resp.Body).Decode(&salmon); err != nil {
		panic(err)
	}

	var term = 0
	if context.Bool("next") {
		term = 1
	}

	var datetimeLayout = "2006/01/02 15:04"
	var startAt = time.Unix(salmon.Result[term].StartT, 0).Format(datetimeLayout)
	var endAt = time.Unix(salmon.Result[term].EndT, 0).Format(datetimeLayout)
	var stage = salmon.Result[term].Stage.Name
	var weapons = []string{}
	for _, w := range salmon.Result[term].Weapons {
		weapons = append(weapons, w.Name)
	}
	var opneingText = ""
	if salmon.Result[term].StartT < time.Now().Unix() && time.Now().Unix() < salmon.Result[term].EndT {
		opneingText = " 現在開催中!"
	}
	fmt.Println("サーモンラン")
	fmt.Println(startAt + " ~ " + endAt + opneingText)
	fmt.Println(strings.Join([]string{"ステージ:", stage}, ""))
	fmt.Println(strings.Join([]string{"ブキ: ", strings.Join(weapons, ", ")}, ""))
}

func GetBattles(context *cli.Context) {
	resp, err := http.Get("https://spla2.yuu26.com/schedule")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	var ika Ika
	if err := json.NewDecoder(resp.Body).Decode(&ika); err != nil {
		panic(err)
	}

	var term = 0
	if context.Bool("next") {
		term = 1
	}

	var datetimeLayout = "2006/01/02 15:04"
	var startAt = time.Unix(ika.Result.Regular[term].StartT, 0).Format(datetimeLayout)
	var endAt = time.Unix(ika.Result.Regular[term].EndT, 0).Format(datetimeLayout)
	var regularMaps = strings.Join(ika.Result.Regular[term].Maps, " ")
	var rankedBattoleRule = ika.Result.Gachi[term].Rule
	var rankedBattoleMaps = strings.Join(ika.Result.Gachi[term].Maps, " ")
	var leagueBattleRule = ika.Result.League[term].Rule
	var leagueBattleMaps = strings.Join(ika.Result.League[term].Maps, " ")
	fmt.Println(startAt + " ~ " + endAt)
	fmt.Println(strings.Join([]string{"ナワバリバトル", ", ステージ:", regularMaps}, ""))
	fmt.Println(strings.Join([]string{"ガチマッチ:", rankedBattoleRule, ", ステージ:", rankedBattoleMaps}, ""))
	fmt.Println(strings.Join([]string{"リーグマッチ:", leagueBattleRule, ", ステージ:", leagueBattleMaps}, ""))
}

func main() {
	var mode string
	app := cli.NewApp()

	app.Name = "ika2cli"
	app.Usage = "Splatoon2のステージ情報を出力するよ"
	app.Version = "0.2.0"

	app.Action = func(context *cli.Context) error {
		if mode == "salmon" {
			GetSalmons(context)
		} else {
			GetBattles(context)
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "mode, m",
			Value:       "regular",
			Usage:       "Regular mode or Salmon-run mode",
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
