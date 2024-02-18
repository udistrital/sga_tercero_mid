package controllers

import (
	"net/url"
	"unicode"

	"github.com/astaxie/beego"
	"github.com/beego/beego/logs"
	"github.com/udistrital/sga_tercero_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/requestresponse"
)

// TerceroController operations for Tercero
type TerceroController struct {
	beego.Controller
}

// URLMapping ...
func (c *TerceroController) URLMapping() {
	c.Mapping("ActualizarPersona", c.ActualizarPersona)
	c.Mapping("GuardarPersona", c.GuardarPersona)
	c.Mapping("GuardarDatosComplementarios", c.GuardarDatosComplementarios)
	c.Mapping("GuardarDatosComplementariosParAcademico", c.GuardarDatosComplementariosParAcademico)
	c.Mapping("ConsultarPersona", c.ConsultarPersona)
	c.Mapping("GuardarDatosContacto", c.GuardarDatosContacto)
	c.Mapping("ConsultarDatosComplementarios", c.ConsultarDatosComplementarios)
	c.Mapping("ConsultarDatosContacto", c.ConsultarDatosContacto)
	c.Mapping("ConsultarDatosFamiliar", c.ConsultarDatosFamiliar)
	c.Mapping("ConsultarDatosFormacionPregrado", c.ConsultarDatosFormacionPregrado)
	c.Mapping("ActualizarDatosComplementarios", c.ActualizarDatosComplementarios)
	c.Mapping("ActualizarInfoFamiliar", c.ActualizarInfoFamiliar)
	c.Mapping("ConsultarInfoEstudiante", c.ConsultarInfoEstudiante)
	c.Mapping("GuardarAutor", c.GuardarAutor)
	c.Mapping("ConsultarExistenciaPersona", c.ConsultarExistenciaPersona)
	c.Mapping("ObtenerTercerosConNIT", c.ObtenerTercerosConNIT)
}

// ActualizarPersona ...
// @Title ActualizarPersona
// @Description Actualizar datos de persona
// @Param	body		body 	{}	true		"body for Actualizar datos de persona content"
// @Success	200	{}
// @Failure	403	body is empty
// @router / [put]
func (c *TerceroController) ActualizarPersona() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.ActualizarPersona(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// GuardarPersona ...
// @Title PostPersona
// @Description Guardar Persona
// @Param	body		body 	{}	true		"body for Guardar Persona content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TerceroController) GuardarPersona() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.GuardarPersona(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// GuardarDatosComplementarios ...
// @Title GuardarDatosComplementarios
// @Description Guardar Datos Complementarios Persona
// @Param	body		body 	{}	true		"body for Guardar Datos Complementarios Persona content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router /complementarios [post]
func (c *TerceroController) GuardarDatosComplementarios() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.GuardarDatosComplementarios(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// GuardarDatosComplementariosParAcademico ...
// @Title GuardarDatosComplementariosParAcademico
// @Description Guardar Datos Complementarios Persona ParAcademico
// @Param	body		body 	{}	true		"body for Guardar Datos Complementarios Persona content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router /complementarios-par [post]
func (c *TerceroController) GuardarDatosComplementariosParAcademico() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.GuardarDatosComplementariosParAcademico(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ActualizarDatosComplementarios ...
// @Title ActualizarDatosComplementarios
// @Description ActualizarDatosComplementarios
// @Param	body	body 	{}	true		"body for Actualizar los datos complementarios content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /complementarios [put]
func (c *TerceroController) ActualizarDatosComplementarios() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.ActualizarDatosComplementarios(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ConsultarExistenciaPersona ...
// @Title ConsultarExistenciaPersona
// @Description get ConsultarExistenciaPersona by NumeroIdentificacion
// @Param	numeroDocumento	path	int 	true	"numero documento del tercero"
// @Success 200 {}
// @Failure 404 not found resource
// @router /existencia/:numeroDocumento [get]
func (c *TerceroController) ConsultarExistenciaPersona() {
	defer errorhandler.HandlePanic(&c.Controller)

	numeroDocumento := c.Ctx.Input.Param(":numeroDocumento")

	resultado, err := services.ConsultarExistenciaPersona(numeroDocumento)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ConsultarPersona ...
// @Title ConsultarPersona
// @Description get ConsultaPersona by id
// @Param	tercero_id	path	int	true	"Id del tercero"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:tercero_id [get]
func (c *TerceroController) ConsultarPersona() {
	defer errorhandler.HandlePanic(&c.Controller)

	//Id del tercero
	idTercero := c.Ctx.Input.Param(":tercero_id")

	resultado, err := services.ConsultarPersona(idTercero)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// GuardarDatosContacto ...
// @Title PostrDatosContacto
// @Description Guardar DatosContacto
// @Param	body		body 	{}	true		"body for Guardar DatosContacto content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router /contacto [post]
func (c *TerceroController) GuardarDatosContacto() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.GuardarDatosContacto(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ConsultarDatosComplementarios ...
// @Title ConsultarDatosComplementarios
// @Description get ConsultarDatosComplementarios by id
// @Param	tercero_id	path	int	true	"Id del ente"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:tercero_id/complementarios [get]
func (c *TerceroController) ConsultarDatosComplementarios() {
	defer errorhandler.HandlePanic(&c.Controller)

	idTercero := c.Ctx.Input.Param(":tercero_id")

	resultado, err := services.ConsultarDatosComplementarios(idTercero)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ConsultarDatosContacto ...
// @Title ConsultarDatosContacto
// @Description get ConsultarDatosContacto by id
// @Param	tercero_id	path	int	true	"Id del Tercero"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:tercero_id/contacto [get]
func (c *TerceroController) ConsultarDatosContacto() {
	defer errorhandler.HandlePanic(&c.Controller)

	//Id de la persona
	idTercero := c.Ctx.Input.Param(":tercero_id")

	resultado, err := services.ConsultarDatosContacto(idTercero)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ConsultarDatosFamiliar ...
// @Title ConsultarDatosFamiliar
// @Description get ConsultarDatosFamiliar by id
// @Param	tercero_id	path	int	true	"Id del Tercero"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:tercero_id/familiar [get]
func (c *TerceroController) ConsultarDatosFamiliar() {
	defer errorhandler.HandlePanic(&c.Controller)

	idTercero := c.Ctx.Input.Param(":tercero_id")

	resultado, err := services.ConsultarDatosFamiliar(idTercero)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ConsultarDatosFormacionPregrado ...
// @Title ConsultarDatosFormacionPregrado
// @Description get ConsultarDatosFormacionPregrado by id
// @Param	tercero_id	path	int	true	"Id del Tercero"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:tercero_id/formacion-pregrado [get]
func (c *TerceroController) ConsultarDatosFormacionPregrado() {
	defer errorhandler.HandlePanic(&c.Controller)

	//Id de la persona
	idTercero := c.Ctx.Input.Param(":tercero_id")
	// resultado datos complementarios persona

	resultado, err := services.ConsultarDatosFormacionPregrado(idTercero)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ActualizarInfoFamiliar ...
// @Title ActualizarInfoFamiliar
// @Description Actualiza la informacion familiar del tercero
// @Param	body	body 	{}	true		"body for Actualizar la info familiar del tercero content"
// @Success 200 {}
// @Failure 403 body is empty
// @router /info-familiar [put]
func (c *TerceroController) ActualizarInfoFamiliar() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.ActualizarInfoFamiliar(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ConsultarPersona ...
// @Title ConsultarInfoSolicitante
// @Description get ConsultarInfoSolicitante by id
// @Param	tercero_id	path	int	true	"Id del tercero"
// @Success 200 {}
// @Failure 404 not found resource
// @router /:tercero_id/info-solicitante [get]
func (c *TerceroController) ConsultarInfoEstudiante() {
	defer errorhandler.HandlePanic(&c.Controller)

	idTercero := c.Ctx.Input.Param(":tercero_id")

	resultado, err := services.ConsultarInfoEstudiante(idTercero)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// GuardarAutor ...
// @Title PostAutor
// @Description Guardar autor
// @Param	body		body 	{}	true		"body for Guardar autor content"
// @Success 201 {int}
// @Failure 400 the request contains incorrect syntax
// @router /autores [post]
func (c *TerceroController) GuardarAutor() {
	defer errorhandler.HandlePanic(&c.Controller)

	data := c.Ctx.Input.RequestBody

	resultado, err := services.GuardarAutor(data)

	if err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err.Error())
	}

	c.ServeJSON()
}

// ObtenerTercerosConNIT maneja la solicitud para obtener una lista de terceros.
// La búsqueda se puede realizar tanto por el NIT (Número de Identificación Tributaria) como por el nombre completo.
// Si se busca por NIT, la función intenta encontrar coincidencias en los números de identificación.
// Si se busca por nombre, intenta encontrar coincidencias en los nombres completos de los terceros.
// La función retorna una lista de terceros, cada uno con su NIT, nombre completo, y un label.
// Este label es una combinación del NIT y el nombre, dependiendo del tipo de búsqueda realizada.
// @Title ObtenerTerceroConNIT
// @Description Retorna una lista de terceros con su NIT y nombre completo.
//
//	La búsqueda puede realizarse por NIT o por nombre completo.
//	El resultado incluye un label que es una combinación de NIT y nombre, dependiendo del criterio de búsqueda.
//
// @Success 200 {array} TerceroConNIT "Lista de terceros con NIT, nombre completo y label correspondiente."
// @Failure 400 "bad request" en caso de una solicitud incorrecta o problemas en la consulta.
// @router /consultar_terceros_con_nit [get]
func (c *TerceroController) ObtenerTercerosConNIT() {
	var query string
	var queryUrl string
	// order: desc,asc
	if v := c.GetString("query"); v != "" {
		query = url.QueryEscape(v)
	}
	// Se arma la query
	if query != "" {
		if esNumerico(query) {
			// Búsqueda por número
			queryUrl = "datos_identificacion?query=TipoDocumentoId:7,Numero__icontains:" + query
		} else {
			// Búsqueda por nombre
			queryUrl = "datos_identificacion?query=TipoDocumentoId:7,TerceroId.NombreCompleto__icontains:" + query
		}
	} else {
		queryUrl = "datos_identificacion?query=TipoDocumentoId:7&limit=0"
	}

	var tercerosConNIT []map[string]interface{}
	//Consultar terceros con nit
	errTerceroConNIT := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+queryUrl, &tercerosConNIT)
	if errTerceroConNIT == nil {
		if tercerosConNIT != nil && len(tercerosConNIT) > 0 {
			type TerceroConNIT struct {
				NIT            string `json:"NIT"`
				NombreCompleto string `json:"NombreCompleto"`
				Label          string `json:"Label"`
			}
			var resultado []TerceroConNIT
			for _, tercero := range tercerosConNIT {
				if terceroData, ok := tercero["TerceroId"].(map[string]interface{}); ok {
					var label string
					if esNumerico(query) {
						label = tercero["Numero"].(string) + " - " + terceroData["NombreCompleto"].(string)
					} else {
						label = terceroData["NombreCompleto"].(string) + " - " + tercero["Numero"].(string)
					}

					terceroConNIT := TerceroConNIT{
						NombreCompleto: terceroData["NombreCompleto"].(string),
						NIT:            tercero["Numero"].(string),
						Label:          label,
					}
					resultado = append(resultado, terceroConNIT)
				}
			}
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Request successful", "Data": resultado}
		}
	} else {
		logs.Error(errTerceroConNIT)
		c.Data["json"] = map[string]interface{}{"Success": false, "Status": "404", "Message": "Data not found", "Data": nil}
		c.Data["system"] = errTerceroConNIT
		c.Abort("404")
	}
	c.ServeJSON()
}

func esNumerico(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
