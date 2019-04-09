package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa simple sample", func() {
	// APIのタイトル
	Title("sample")
	// APIの説明
	Description("適当に作るよ")
	// 作者へのコンタクト情報
	Contact(func() {
		Name("hoge")
		Email("hoge@hogehoge.com")
		URL("https://sample")
	})
	// APIのライセンス
	License(func() {
		Name("MIT")
		URL("https://hogehoge/license")
	})
	// APIのドキュメント
	Docs(func() {
		Description("wiki")
		URL("https://hoge/wiki")
	})
	// ホストの設定
	Host("localhost;8080")
	// 対応しているプロトコル定義
	Scheme("http", "https")
	// 全てのエンドポイントへのベースパス
	BasePath("/api/v1")

	// CORSポリシーの定義
	Origin("http://localhost:8080/swagger", func() {
		// クライアントに公開された1つ以上のヘッダー
		Expose("X-Time")
		// 許可するHTTPメソッド
		Methods("GET", "POST", "PUT", "DELETE")
		// プリフライト要求応答をキャッシュする時間
		MaxAge(600)
		// Access-Control-Allow-Credentialsヘッダーの設定
		Credentials()
	})
})
