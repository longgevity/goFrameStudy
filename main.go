package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	//: 命名匹配规则 ；* 模糊匹配规则 ；{field} 字段匹配规则
	s := g.Server()
	s.BindHandler("GET:/:", luyou)
	s.BindHandler("/:name/update", luyou)
	s.BindHandler("/:name/:action", luyou)
	s.BindHandler("POST:/:/*", luyou)
	s.BindHandler("/user/list/{field}.html", luyou)
	s.SetPort(80)
	s.Run()
}

func luyou(r *ghttp.Request) {
	r.Response.WriteJson(r.Get("name"))
}
