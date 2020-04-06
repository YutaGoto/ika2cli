package salmon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

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
	var hours = time.Unix(salmon.Result[term].EndT, 0).Sub(time.Unix(salmon.Result[term].StartT, 0))
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
	fmt.Println(startAt + " ~ " + endAt + " (" + fmt.Sprintf("%.f", hours.Hours()) + "h)"  + opneingText)
	fmt.Println(strings.Join([]string{"ステージ:", stage}, ""))
	fmt.Println(strings.Join([]string{"ブキ: ", strings.Join(weapons, ", ")}, ""))
}
