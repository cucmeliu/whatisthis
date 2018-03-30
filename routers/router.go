// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/cucmeliu/whatisthis/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/a",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),

		beego.NSNamespace("/plant",
			beego.NSInclude(
				&controllers.PlantController{},
			),
		),

		beego.NSNamespace("/rec_logs",
			beego.NSInclude(
				&controllers.RecLogsController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
