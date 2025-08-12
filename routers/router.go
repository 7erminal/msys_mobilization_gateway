// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"msys_api_gateway/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/clients",
	// 		beego.NSInclude(
	// 			&controllers.ClientsController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/services",
	// 		beego.NSInclude(
	// 			&controllers.ServicesController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/client_services",
	// 		beego.NSInclude(
	// 			&controllers.Client_servicesController{},
	// 		),
	// 	),
	// )

	// beego.AddNamespace(ns)

	bs := beego.NewNamespace("/v2",

		beego.NSNamespace("/api",
			beego.NSInclude(
				&controllers.Service_requestsController{},
			),
		),
		beego.NSNamespace("/clients",
			beego.NSInclude(
				&controllers.ClientsController{},
			),
		),
	)

	beego.AddNamespace(bs)
}
