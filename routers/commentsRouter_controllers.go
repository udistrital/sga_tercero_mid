package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "GuardarPersona",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ActualizarPersona",
			Router:           "/",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ConsultarPersona",
			Router:           "/:tercero_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ConsultarDatosComplementarios",
			Router:           "/:tercero_id/complementarios",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ConsultarDatosContacto",
			Router:           "/:tercero_id/contacto",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ConsultarDatosFamiliar",
			Router:           "/:tercero_id/familiar",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ConsultarDatosFormacionPregrado",
			Router:           "/:tercero_id/formacion-pregrado",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ConsultarInfoEstudiante",
			Router:           "/:tercero_id/info-solicitante",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "GuardarAutor",
			Router:           "/autores",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ActualizarDatosComplementarios",
			Router:           "/complementarios",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "GuardarDatosComplementarios",
			Router:           "/complementarios",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "GuardarDatosComplementariosParAcademico",
			Router:           "/complementarios-par",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "GuardarDatosContacto",
			Router:           "/contacto",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ConsultarExistenciaPersona",
			Router:           "/existencia/:numeroDocumento",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
		beego.ControllerComments{
			Method:           "ActualizarInfoFamiliar",
			Router:           "/info-familiar",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
