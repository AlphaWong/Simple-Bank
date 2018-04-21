// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gitlab.com/Simple-Bank/controllers"
	"gitlab.com/Simple-Bank/utils"

	"github.com/astaxie/beego"
)

func init() {
	initDB()
	beego.Include(&controllers.AccountController{})
}

func initDB() {
	utils.InitDB()
}
