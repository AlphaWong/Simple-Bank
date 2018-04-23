package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"] = append(beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Send",
			Router: `/v1/accounts/:from/send/:to`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"] = append(beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Balance",
			Router: `/v1/accounts/:id/balance`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"] = append(beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Close",
			Router: `/v1/accounts/:id/close`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"] = append(beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Deposit",
			Router: `/v1/accounts/:id/deposit`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"] = append(beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Withdraw",
			Router: `/v1/accounts/:id/withdraw`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"] = append(beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/v1/accounts/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"] = append(beego.GlobalControllerRouter["gitlab.com/Simple-Bank/controllers:AccountController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/v1/customers/:id/accounts/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
