package main

//
// 在使用了git配置的项目中，引入项目内的包需以github.com/....的形式，如下
// github.com/T25ldGhpbmdRdGh3/QkFDS0VORA00  +  go.mod所在目录下的包路径
//
// 请注意函数及变量的命名规则(代码规范)
//

import (
	"github.com/T25ldGhpbmdRdGh3/QkFDS0VORA00/middleware"
	"github.com/T25ldGhpbmdRdGh3/QkFDS0VORA00/route"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/pprof"
)

func main() {
	app := iris.New()

	// 中间件部分 相关功能在middleware包中
	middleware.Register(app)
	// 路由 在route包中
	// 这部分规则是：
	// route/routes.go 中包含所有的路由选项
	// 路由所调用的所有 handlers 全部放在api包中
	route.Routes(app)
	// 提供性能分析 参考 https://www.jianshu.com/p/4e4ff6be6af9
	app.Any("/debug/pprof/{action:path}", pprof.New())
	// app.Run(iris.Addr("0.0.0.0:" + config.Get().Port))
	app.Run(iris.Addr("0.0.0.0:8080"))
}
