package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var IntegerMedia = MediaType("application/vnd.integer+json", func() {
	// 説明
	Description("example Response")
	// どのような値があるか
	Attributes(func() {
		Attribute("id", Integer, "id", func() {
			// レスポンスの例
			Example(1)
		})
		Required("id")
	})
	// 返すレスポンスのフォーマット
	View("default", func() {
		Attribute("id")
	})
})

var MessageMedia = MediaType("application/vnd.message+json", func() {
	Description("example")
	Attributes(func() {
		Attribute("message", String, "メッセージ", func() {
			Example("ok")
		})
		Required("message")
	})
	View("default", func() {
		Attribute("message")
	})
})
