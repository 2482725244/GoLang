// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"helloworld/controllers"
)

//	func init() {
//		ns := beego.NewNamespace("/v1",
//			beego.NSNamespace("/object",
//				beego.NSInclude(
//					&controllers.ObjectController{},
//				),
//			),
//			beego.NSNamespace("/user",
//				beego.NSInclude(
//					&controllers.UserController{},
//				),
//			),
//		)
//		beego.AddNamespace(ns)
//	}
func init() {
	beego.Router("/webaseset", &controllers.WeBASEController{}, "post:Set")
	beego.Router("/webaseget", &controllers.WeBASEController{}, "get:Get")
	beego.Router("/sdkget", &controllers.SdkController{}, "get:Get")
	beego.Router("/sdkset", &controllers.SdkController{}, "post:Set")
}
