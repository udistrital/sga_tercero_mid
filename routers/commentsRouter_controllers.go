package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ActualizarDatosComplementarios",
            Router: "/actualizar_complementarios",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ActualizarPersona",
            Router: "/actualizar_persona",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ConsultarDatosComplementarios",
            Router: "/consultar_complementarios/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ConsultarDatosContacto",
            Router: "/consultar_contacto/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ConsultarDatosFamiliar",
            Router: "/consultar_familiar/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ConsultarDatosFormacionPregrado",
            Router: "/consultar_formacion_pregrado/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ConsultarInfoEstudiante",
            Router: "/consultar_info_solicitante/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ConsultarPersona",
            Router: "/consultar_persona/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ObtenerTercerosConNIT",
            Router: "/consultar_terceros_con_nit",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ConsultarExistenciaPersona",
            Router: "/existe_persona/:numeroDocumento",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GuardarAutor",
            Router: "/guardar_autor",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GuardarDatosComplementarios",
            Router: "/guardar_complementarios",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GuardarDatosComplementariosParAcademico",
            Router: "/guardar_complementarios_par",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GuardarDatosContacto",
            Router: "/guardar_datos_contacto",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "GuardarPersona",
            Router: "/guardar_persona",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/sga_mid_tercero/controllers:TerceroController"],
        beego.ControllerComments{
            Method: "ActualizarInfoFamiliar",
            Router: "/info_familiar",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
