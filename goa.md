# Goa入れてみる
ここら辺参考
https://tikasan.hatenablog.com/entry/2017/05/05/212501
## What's Goa
DSLでAPIの仕様を書くと必要なものを作成してくれるツール  
* swaggerできるよ(仕様書二重管理からの解放)
* app,client,tool,swaggerは変更する必要なし
## goaの使用流れ
1. designで雛形作る
2. `goagen bootstrap -d`でdesignからファイル作る
3. エンドポイントでの処理を書く

## designの書き方
* 慣習でdesignパッケージ以下に書いてく
* APIの設計はわけた方がよさげ
   * 基本情報
   * レスポンスタイプ
   * リソースへの操作?