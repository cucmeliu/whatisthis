package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:MainController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:PlantController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:PlantController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:RecLogsController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/cucmeliu/whatisthis/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
