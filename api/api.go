package api

import "github.com/kataras/iris"

// Error ...
func Error(ctx iris.Context, code int, msg string, data interface{}) {
	ctx.JSON(map[string]interface{}{
		"status": code,
		"msg":    msg,
		"data":   data,
	})
}

// Success ...
func Success(ctx iris.Context, msg string, data interface{}) {
	ctx.JSON(map[string]interface{}{
		"status": 0,
		"msg":    msg,
		"data":   data,
	})
}
