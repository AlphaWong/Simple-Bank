package routers

import (
	httpContext "context"
	"log"
	"net/http/httputil"

	"gitlab.com/Simple-Bank/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	uuid "github.com/satori/go.uuid"
)

func InitMiddleware() {
	AddFootPrintMiddleware()
}

func AddFootPrintMiddleware() {
	var foodPrintMiddleware = func(ctx *context.Context) {
		footPrint := uuid.NewV4().String()
		dump, _ := httputil.DumpRequest(ctx.Request, true)
		log.Printf("footPrint: %s request: %v", footPrint, string(dump))
		nativeCtx := httpContext.WithValue(ctx.Request.Context(), utils.ContextFootPrintKey, footPrint)
		ctx.Request = ctx.Request.WithContext(nativeCtx)
	}

	beego.InsertFilter("/v1/*", beego.BeforeRouter, foodPrintMiddleware)
}
