# ika2cli

[![CircleCI](https://circleci.com/gh/YutaGoto/ika2cli.svg?style=svg)](https://circleci.com/gh/YutaGoto/ika2cli)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/d24185811973486b9015548a8ca9fc7f)](https://www.codacy.com/app/YutaGoto/ika2cli?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=YutaGoto/ika2cli&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/ace630bf6b66a592179f/maintainability)](https://codeclimate.com/github/YutaGoto/ika2cli/maintainability)

## Usage

The following command available.

To Check splatoon2 battle stages and rules or Salmon-Run stage and weapons.

```sh
ika2cli
```

### Options

#### `-n`, `--next`

To show next terms informations.

#### `--mode value`, `-m value`

If value=salmon, to show Salmon-Run stage and weapons

#### `-s` `--search`

To search when the specified rule will start next.

##### `-g` `--gachi`

Ranked match

##### `-l` `--league`

League match

## Examples

```sh
$ ./ika2cli
2019/03/10 15:00 ~ 2019/03/10 17:00
ナワバリバトル, ステージ:ハコフグ倉庫 マンタマリア号
ガチマッチ:ガチアサリ, ステージ:アロワナモール ショッツル鉱山
リーグマッチ:ガチエリア, ステージ:Ｂバスパーク ホッケふ頭
$ ./ika2cli -n
2019/03/10 17:00 ~ 2019/03/10 19:00
ナワバリバトル, ステージ:モズク農園 アロワナモール
ガチマッチ:ガチエリア, ステージ:チョウザメ造船 ハコフグ倉庫
リーグマッチ:ガチアサリ, ステージ:アジフライスタジアム ホテルニューオートロ
$ ./ika2cli -m salmon
サーモンラン
2019/03/09 09:00 ~ 2019/03/10 21:00 現在開催中!
ステージ:シェケナダム
ブキ: ？, ？, ？, ？
$ ./ika2cli -m salmon -n
サーモンラン
2019/03/11 09:00 ~ 2019/03/12 15:00
ステージ:海上集落シャケト場
ブキ: スプラローラー, パブロ, .96ガロン, バレルスピナー
$ ./ika2cli -s -g zone
2019/03/10 17:00 ~ 2019/03/10 19:00
ガチマッチ:ガチエリア, ステージ:チョウザメ造船 ハコフグ倉庫
$ ./ika2cli -s -l tower
2019/03/10 19:00 ~ 2019/03/10 21:00
リーグマッチ:ガチヤグラ, ステージ:ガンガゼ野外音楽堂 海女美術大学
```

This DateTime is JST.

![example.giff](https://raw.githubusercontent.com/YutaGoto/ika2cli/master/example.gif)

## Note

This command is NOT related to Nintendo Co., Ltd.

It uses an unofficial API.

[https://spla2.yuu26.com/](https://spla2.yuu26.com/)
