package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
)

func UpdateOrCreateInfoComplementaria(tipoInfo string, infoComp map[string]interface{}, idTercero float64) (map[string]interface{}, bool) {
	resp := map[string]interface{}{}
	ok := false

	if infoComp[tipoInfo].(map[string]interface{})["hasId"] != nil {
		idInfComp := infoComp[tipoInfo].(map[string]interface{})["hasId"].(float64)
		var updateInfoComp map[string]interface{}
		errUpdtInfoComp := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"info_complementaria_tercero/"+fmt.Sprintf("%v", idInfComp), &updateInfoComp)
		if errUpdtInfoComp == nil && updateInfoComp["Status"] != 404 {
			dataToUpdate := infoComp[tipoInfo].(map[string]interface{})["data"].(map[string]interface{})
			updateInfoComp["InfoComplementariaId"] = dataToUpdate

			var updateAnswer map[string]interface{}
			errupdateAnswer := request.SendJson("http://"+beego.AppConfig.String("TercerosService")+"info_complementaria_tercero/"+fmt.Sprintf("%.f", idInfComp), "PUT", &updateAnswer, updateInfoComp)
			if errupdateAnswer == nil {
				resp = updateAnswer
				ok = true
			}
		}
	} else {
		newInfo := map[string]interface{}{
			"TerceroId":            map[string]interface{}{"Id": idTercero},
			"InfoComplementariaId": infoComp[tipoInfo].(map[string]interface{})["data"].(map[string]interface{}),
			"Activo":               true,
		}
		var createinfo map[string]interface{}
		errCreateInfo := request.SendJson("http://"+beego.AppConfig.String("TercerosService")+"info_complementaria_tercero", "POST", &createinfo, newInfo)
		if errCreateInfo == nil && fmt.Sprintf("%v", createinfo) != "map[]" && createinfo["Id"] != nil {
			resp = createinfo
			ok = true
		}
	}

	return resp, ok
}

type APIResponse struct {
	Body []byte
	Err  error
}

func SendRequestToCRUDAPI(endpoint string, data interface{}, method string) APIResponse {

	APICRUDURL := beego.AppConfig.String("router.APICRUD")

	rutaCompleta := APICRUDURL + endpoint
	logs.Info("Ruta completa: ", rutaCompleta)

	// Inicializar la respuesta
	var apiResp APIResponse

	// Convertir los datos a JSON
	reqBody, err := json.Marshal(data)
	if err != nil {
		apiResp.Err = err
	}

	// Configurar la solicitud HTTP
	req, err := http.NewRequest(method, rutaCompleta, bytes.NewBuffer(reqBody))
	if err != nil {
		apiResp.Err = err
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear cliente HTTP y enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		apiResp.Err = err
	}
	defer resp.Body.Close()

	// Leer la respuesta de la API CRUD
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		apiResp.Err = err
	}

	// logs.Info("Respuesta de la API CRUD: ", string(respBody))

	// Asignar la respuesta al campo Body de la estructura APIResponse
	apiResp.Body = respBody

	return apiResp
}
