package route

import (
	"github.com/kataras/iris"
	"github.com/T25ldGhpbmdRdGh3/QkFDS0VORA00/api"
)

// Routes 路由规则
func Routes(app *iris.Application) {
	// 设置路由策略
	
	// 以eg页面举例 所有handlers都以api包中函数的形式调用
	app.Get("/eg", api.Eg)
	app.Get("/eg/1", api.Eg1)
	app.Get("/eg/2", api.Eg2)
}
