package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// actionsの定義をする
var _ = Resource("actions", func() {
	// actionsリソースのベースパス
	BasePath("/actions")
	/*
		リソースに対しての操作
		add: リソース追加
		list リソースをリストで取得
		delete　リソースを削除する
	*/
	Action("ping", func() {
		Description("サーバーとの疎通確認")
		Routing(
			// エンドポイント
			GET("/ping"),
		)
		Response(OK, MessageMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("hello", func() {
		Description("挨拶します")
		Routing(
			GET("/hello"),
		)
		Params(func() {
			Param("name", String, "名前", func() {
				Default("")
			})
			Required("name")
		})
		Response(OK, MessageMedia)
		Response(BadRequest, ErrorMedia)
	})
	Action("ID", func() {
		Description("複数アクション(:id)")
		Routing(
			// エンドポイントにリソース指定可能
			GET("/:id"),
		)
		Params(func() {
			Param("id", Integer, "id")
			// リソース指定されている場合はRequired不要
		})
		Response(OK, MessageMedia)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})
