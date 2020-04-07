package battle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
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

// ConvertRuleEn2Ja is translating En to Ja
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

// BuildRule is join rule strings
func BuildRule(ruleType string, ruleName string, mapName string) string {
	return strings.Join([]string{ruleType, ":", ruleName, ", ステージ:", mapName }, "")
}

// BuildDateTime is build start and end
func BuildDateTime(startT int64, endT int64) string {
	var datetimeLayout = "2006/01/02 15:04"
	var startAt = time.Unix(startT, 0).Format(datetimeLayout)
	var endAt = time.Unix(endT, 0).Format(datetimeLayout)
	return (startAt + " ~ " + endAt)
}

// SearchBattles can search next league or gachi type of battle
func SearchBattles(league string, gachi string) {
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

	if league != "" {
		var leagues = ika.Result.League
		var leagueRuleName = ConvertRuleEn2Ja(league)
		FOR_LEAGUE_RULE_LABEL:
			for _, l := range leagues {
				if leagueRuleName == l.Rule {
					fmt.Println(BuildDateTime(l.StartT, l.EndT))
					fmt.Println(BuildRule("リーグマッチ", l.Rule, strings.Join(l.Maps, " ")))
					break FOR_LEAGUE_RULE_LABEL
				}
			}
	} else if gachi != "" {
		var gachis = ika.Result.Gachi
		var gachiRuleName = ConvertRuleEn2Ja(gachi)
		FOR_GACHI_RULE_LABEL:
			for _, g := range gachis {
				if gachiRuleName == g.Rule {
					fmt.Println(BuildDateTime(g.StartT, g.EndT))
					fmt.Println(BuildRule("ガチマッチ", g.Rule, strings.Join(g.Maps, " ")))
					break FOR_GACHI_RULE_LABEL
				}
			}
	} else {
		fmt.Println("ガチかリーグを選択してください")
		return
	}
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

	var regularMaps = strings.Join(ika.Result.Regular[term].Maps, " ")
	var rankedBattoleRule = ika.Result.Gachi[term].Rule
	var rankedBattoleMaps = strings.Join(ika.Result.Gachi[term].Maps, " ")
	var leagueBattleRule = ika.Result.League[term].Rule
	var leagueBattleMaps = strings.Join(ika.Result.League[term].Maps, " ")
	fmt.Println(BuildDateTime(ika.Result.Regular[term].StartT, ika.Result.Regular[term].EndT))
	fmt.Println(strings.Join([]string{"ナワバリバトル", ", ステージ:", regularMaps}, ""))
	fmt.Println(BuildRule("ガチマッチ", rankedBattoleRule, rankedBattoleMaps))
	fmt.Println(BuildRule("リーグマッチ", leagueBattleRule, leagueBattleMaps))
}
