package routers

import (
	httpContext "context"
	"log"
	"net/http/httputil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/Simple-Bank/utils"
)

func InitMiddleware() {
	AddFootPrintMiddleware()
}

func AddFootPrintMiddleware() {
	var FoodPrintMiddleware = func(ctx *context.Context) {
		footPrint := uuid.NewV4().String()
		dump, _ := httputil.DumpRequest(ctx.Request, true)
		log.Printf("footPrint: %s request: %v", footPrint, string(dump))
		nativeCtx := httpContext.WithValue(ctx.Request.Context(), utils.ContextFootPrintKey, footPrint)
		ctx.Request = ctx.Request.WithContext(nativeCtx)
	}

	beego.InsertFilter("/v1/*", beego.BeforeRouter, FoodPrintMiddleware)
}
