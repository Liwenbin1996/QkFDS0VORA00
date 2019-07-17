package api

import (
	"github.com/kataras/iris"
)

// Eg ...
func Eg(ctx iris.Context) {
	ctx.JSON("eg")
}

// Eg1 ...
func Eg1(ctx iris.Context) {
	ctx.JSON("eg1")
}

// Eg2 ...
func Eg2(ctx iris.Context) {
	ctx.JSON("eg2")
}
