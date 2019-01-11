package main

import (
    "os"
    "fmt"
    "strings"
    "time"
    "net/http"
    "encoding/json"
    "github.com/urfave/cli"
)

type Ika struct {
  Result struct {
    Regular []struct {
      Rule   string `json:"rule"`
      Maps   []string `json:"maps"`
      StartT   int64       `json:"start_t"`
      EndT     int64       `json:"end_t"`
    } `json:"regular"`
    Gachi []struct {
      Rule   string `json:"rule"`
      Maps   []string `json:"maps"`
    } `json:"gachi"`
    League []struct {
      Rule   string `json:"rule"`
      Maps   []string `json:"maps"`
    } `json:"league"`
  } `json:"result"`
}

func main() {
  app := cli.NewApp()

  app.Name = "ika2"
  app.Usage = "Splatoon2のステージ情報を出力するよ"
  app.Version = "0.0.1"

  app.Action = func (context *cli.Context) error {

    resp, err := http.Get("https://spla2.yuu26.com/schedule")

    var term = 0
    if context.Bool("next") {
      term = 1
    }

    if err != nil {
      fmt.Println(err)
      return nil
    }

    defer resp.Body.Close()
    var ika Ika
    if err := json.NewDecoder(resp.Body).Decode(&ika); err != nil {
        panic(err)
    }

    var datetime_layout = "2006/01/02 15:04"
    fmt.Println(time.Unix(ika.Result.Regular[term].StartT, 0).Format(datetime_layout) + " ~ " + time.Unix(ika.Result.Regular[term].EndT, 0).Format(datetime_layout))
    fmt.Println(strings.Join([]string{"ナワバリバトル", ", ステージ:", strings.Join(ika.Result.Regular[term].Maps, " ")}, ""))
    fmt.Println(strings.Join([]string{"ガチマッチ:", ika.Result.Gachi[term].Rule  , ", ステージ:", strings.Join(ika.Result.Gachi[term].Maps,   " ")}, ""))
    fmt.Println(strings.Join([]string{"リーグマッチ:", ika.Result.League[term].Rule , ", ステージ:", strings.Join(ika.Result.League[term].Maps,  " ")}, ""))
    return nil
  }

  app.Flags = []cli.Flag{
    cli.BoolFlag {
      Name: "next, n",
      Usage: "to show next term",
    },
  }

  app.Run(os.Args)
}
