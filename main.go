package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Ika is Battle struct
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
			StartT int64    `json:"start_t"`
			EndT   int64    `json:"end_t"`
		} `json:"gachi"`
		League []struct {
			Rule string   `json:"rule"`
			Maps []string `json:"maps"`
			StartT int64    `json:"start_t"`
			EndT   int64    `json:"end_t"`
		} `json:"league"`
	} `json:"result"`
}

// Salmon is SalmonRun struct
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

var (
	mode = kingpin.Flag("mode", "Set verbose mode").Short('m').String()
	next = kingpin.Flag("next", "next terms").Short('n').Bool()
	search = kingpin.Flag("search", "search mode").Short('s').Bool()
	league = kingpin.Flag("league", "search next league").Short('l').String()
	gachi = kingpin.Flag("gachi", "search next gachi").Short('g').String()
)

func ConvertRuleEn2Ja(name string) string {
	switch name {
	case "zone", "Zone":
		return "ガチエリア"
	case "rain", "Rain":
		return "ガチホコバトル"
	case "clam", "Clam":
		return "ガチアサリ"
	case "tower", "Tower":
		return "ガチヤグラ"
	default:
		return name
	}
}

// SearchBattles can search next league or gachi type of battle
func SearchBattles(league string, gachi string) {
	resp, err := http.Get("https://spla2.yuu26.com/schedule")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	var datetimeLayout = "2006/01/02 15:04"

	var ika Ika
	if err := json.NewDecoder(resp.Body).Decode(&ika); err != nil {
		panic(err)
	}

	if league != "" {
		var leagues = ika.Result.League
		var leagueRuleName = ConvertRuleEn2Ja(league)
		FOR_LEAGUE_RULE_LABEL:
			for _, l := range leagues {
				if leagueRuleName == l.Rule {
					var ruleName = l.Rule
					var mapName = strings.Join(l.Maps, " ")
					var startAt = time.Unix(l.StartT, 0).Format(datetimeLayout)
					var endAt = time.Unix(l.EndT, 0).Format(datetimeLayout)
					fmt.Println(startAt + " ~ " + endAt)
					fmt.Println(strings.Join([]string{"リーグマッチ:", ruleName, ", ステージ:", mapName }, ""))
					break FOR_LEAGUE_RULE_LABEL
				}
			}
	} else if gachi != "" {
		var gachis = ika.Result.Gachi
		var gachiRuleName = ConvertRuleEn2Ja(gachi)
		FOR_GACHI_RULE_LABEL:
			for _, g := range gachis {
				if gachiRuleName == g.Rule {
					var ruleName = g.Rule
					var mapName = strings.Join(g.Maps, " ")
					var startAt = time.Unix(g.StartT, 0).Format(datetimeLayout)
					var endAt = time.Unix(g.EndT, 0).Format(datetimeLayout)
					fmt.Println(startAt + " ~ " + endAt)
					fmt.Println(strings.Join([]string{"ガチマッチ:", ruleName, ", ステージ:", mapName }, ""))
					break FOR_GACHI_RULE_LABEL
				}
			}
	} else {
		fmt.Println("ガチかリーグを選択してください")
		return
	}
}

// GetSalmons can get Salmon-Run informations
func GetSalmons(next bool) {
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
	if next {
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

// GetBattles can get Splatoon2 Battles informations
func GetBattles(next bool) {
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
	if next {
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

func run(args []string) {
	kingpin.Version("0.3.0")
	_, err := kingpin.CommandLine.Parse(args)
	if err != nil {
		fmt.Println(err)
	}

	if *mode == "salmon" {
		GetSalmons(*next)
	} else if *search {
		SearchBattles(*league, *gachi)
	} else {
		GetBattles(*next)
	}

}

func main() {
	run(os.Args[1:])
}
