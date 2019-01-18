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

## Examples

```sh
$ ./ika2cli
2019/01/18 21:00 ~ 2019/01/18 23:00
ナワバリバトル, ステージ:バッテラストリート モンガラキャンプ場
ガチマッチ:ガチエリア, ステージ:ザトウマーケット エンガワ河川敷
リーグマッチ:ガチアサリ, ステージ:ホテルニューオートロ ショッツル鉱山
$ ./ika2cli -n
2019/01/18 23:00 ~ 2019/01/19 01:00
ナワバリバトル, ステージ:ザトウマーケット デボン海洋博物館
ガチマッチ:ガチホコバトル, ステージ:アンチョビットゲームズ チョウザメ造船
リーグマッチ:ガチエリア, ステージ:タチウオパーキング スメーシーワールド
$ ./ika2cli -m salmon
サーモンラン
2019/01/17 15:00 ~ 2019/01/18 21:00 開催中
ステージ:海上集落シャケト場
ブキ: わかばシューター, ダイナモローラー, ケルビン525, ソイチューバー
$ ./ika2cli -m salmon -n
サーモンラン
2019/01/19 09:00 ~ 2019/01/20 21:00
ステージ:難破船ドン・ブラコ
ブキ: ？, ？, ？, ？
```

This DateTime is JST.

![example.giff](https://raw.githubusercontent.com/YutaGoto/ika2cli/master/example.gif)

## Note

This command is NOT related to Nintendo Co., Ltd.

It uses an unofficial API.

[https://spla2.yuu26.com/](https://spla2.yuu26.com/)
