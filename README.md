# ika2cli

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/897b37819dc7423297152315d6e4f021)](https://app.codacy.com/app/YutaGoto/ika2cli?utm_source=github.com&utm_medium=referral&utm_content=YutaGoto/ika2cli&utm_campaign=Badge_Grade_Dashboard)

## Usage

The following command available.

To Check splatoon2 battle stages and rules.

```sh
ika2cli
```

### Options

#### `-n`, `--next`

To show terms stages and rules

## Examples

```sh
$ ika2cli
2019/01/12 13:00 ~ 2019/01/12 15:00
ナワバリバトル, ステージ:アジフライスタジアム スメーシーワールド
ガチマッチ:ガチエリア, ステージ:ショッツル鉱山 コンブトラック
リーグマッチ:ガチホコバトル, ステージ:ホッケふ頭 ムツゴ楼
$ ika2cli -n
2019/01/12 15:00 ~ 2019/01/12 17:00
ナワバリバトル, ステージ:ホッケふ頭 ショッツル鉱山
ガチマッチ:ガチヤグラ, ステージ:ガンガゼ野外音楽堂 アジフライスタジアム
リーグマッチ:ガチエリア, ステージ:ハコフグ倉庫 アロワナモール
```

This DateTime is JST.

![example.giff](https://raw.githubusercontent.com/YutaGoto/ika2cli/master/example.gif)

## Note

This command is NOT related to Nintendo Co., Ltd.

It uses an unofficial API.

[https://spla2.yuu26.com/](https://spla2.yuu26.com/)
