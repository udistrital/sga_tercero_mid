// @APIVersion 1.0.0
// @Title Microservicio SGA MID - Terceros
// @Description Microservicio SGA MID del microservicio terceros
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/sga_mid_tercero/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/terceros",
			beego.NSInclude(
				&controllers.TerceroController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
