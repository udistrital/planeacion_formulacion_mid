package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego"
	formulacionhelper "github.com/udistrital/planeacion_formulacion_mid/helpers/formulacionHelper"
	"github.com/udistrital/planeacion_formulacion_mid/models"
	"github.com/udistrital/utils_oas/request"
	"golang.org/x/sync/errgroup"
)

func ClonarFormato(id string, body []byte) (interface{}, error) {

	var respuesta map[string]interface{}
	var respuestaHijos map[string]interface{}
	var hijos []map[string]interface{}
	var planFormato map[string]interface{}
	var parametros map[string]interface{}
	var resPost map[string]interface{}
	var resLimpia map[string]interface{}

	plan := make(map[string]interface{})
	clienteHttp := &http.Client{}
	url := "http://" + beego.AppConfig.String("PlanesService") + "/plan/"

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &respuesta); err == nil {

		request.LimpiezaRespuestaRefactor(respuesta, &planFormato)
		json.Unmarshal(body, &parametros)

		plan["nombre"] = "" + planFormato["nombre"].(string)
		plan["descripcion"] = planFormato["descripcion"].(string)
		plan["tipo_plan_id"] = planFormato["tipo_plan_id"].(string)
		plan["aplicativo_id"] = planFormato["aplicativo_id"].(string)
		plan["activo"] = planFormato["activo"]
		plan["formato"] = false
		plan["vigencia"] = parametros["vigencia"].(string)
		plan["dependencia_id"] = parametros["dependencia_id"].(string)
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd" // En formulación
		plan["formato_id"] = id
		plan["nueva_estructura"] = true

		aux, err := json.Marshal(plan)
		if err != nil {
			return nil, errors.New("error del servicio ClonarFormato: Error codificado " + err.Error())
		}

		peticion, err := http.NewRequest("POST", url, bytes.NewBuffer(aux))
		if err != nil {
			return nil, errors.New("error del servicio ClonarFormato: Error creando peticion " + err.Error())
		}
		peticion.Header.Set("Content-Type", "application/json; charset=UTF-8")
		respuesta, err := clienteHttp.Do(peticion)
		if err != nil {
			return nil, errors.New("error del servicio ClonarFormato: Error haciendo peticion " + err.Error())
		}

		defer respuesta.Body.Close()

		cuerpoRespuesta, err := io.ReadAll(respuesta.Body)
		if err != nil {
			return nil, errors.New("error del servicio ClonarFormato: Error leyendo peticion " + err.Error())
		}

		json.Unmarshal(cuerpoRespuesta, &resPost)
		resLimpia = resPost["Data"].(map[string]interface{})
		padre := resLimpia["_id"].(string)

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &respuestaHijos); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
			formulacionhelper.ClonarHijos(hijos, padre)
		} else {
			return nil, errors.New("error del servicio ClonarFormato: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}

	} else {
		return nil, errors.New("error del servicio ClonarFormato: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}

	return resPost, nil
}

func ClonarPI_PED(id string, body []byte) (interface{}, error) {

	var respuesta map[string]interface{}
	var respuestaHijos map[string]interface{}
	var hijos []map[string]interface{}
	var planFormato map[string]interface{}
	var parametros map[string]interface{}

	plan := make(map[string]interface{})
	clienteHttp := &http.Client{}
	url := "http://" + beego.AppConfig.String("PlanesService") + "/plan/"

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &respuesta); err == nil {

		request.LimpiezaRespuestaRefactor(respuesta, &planFormato)
		json.Unmarshal(body, &parametros)

		plan["nombre"] = "" + planFormato["nombre"].(string) + " - CLONADO"
		plan["descripcion"] = planFormato["descripcion"].(string)
		plan["tipo_plan_id"] = planFormato["tipo_plan_id"].(string)
		plan["aplicativo_id"] = planFormato["aplicativo_id"].(string)
		plan["activo"] = planFormato["activo"]
		plan["formato"] = planFormato["formato"]

		var resPost map[string]interface{}
		var resLimpia map[string]interface{}

		aux, err := json.Marshal(plan)
		if err != nil {
			log.Fatalf("Error codificado: %v", err)
		}

		peticion, err := http.NewRequest("POST", url, bytes.NewBuffer(aux))
		if err != nil {
			log.Fatalf("Error creando peticion: %v", err)
		}
		peticion.Header.Set("Content-Type", "application/json; charset=UTF-8")
		respuesta, err := clienteHttp.Do(peticion)
		if err != nil {
			log.Fatalf("Error haciendo peticion: %v", err)
		}

		defer respuesta.Body.Close()

		cuerpoRespuesta, err := io.ReadAll(respuesta.Body)
		if err != nil {
			log.Fatalf("Error leyendo peticion: %v", err)
		}

		json.Unmarshal(cuerpoRespuesta, &resPost)
		resLimpia = resPost["Data"].(map[string]interface{})
		padre := resLimpia["_id"].(string)

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &respuestaHijos); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
			formulacionhelper.ClonarHijos(hijos, padre)
		}
		return resPost, nil

	}
	return nil, errors.New("error del servicio ClonarPI_PED: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido")
}

func GuardarActividad(id string, datos []byte) (interface{}, error) {
	var body map[string]interface{}
	var res map[string]interface{}
	var entrada map[string]interface{}
	var resPlan map[string]interface{}
	var plan map[string]interface{}
	var armonizacionExecuted bool = false

	json.Unmarshal(datos, &body)
	entrada = body["entrada"].(map[string]interface{})
	armonizacion := body["armo"]
	armonizacionPI := body["armoPI"]
	maxIndex := formulacionhelper.GetIndexActividad(entrada)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &resPlan); err != nil {
		return nil, errors.New("error del servicio GuardarActividad: La solicitud getPlan contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
	request.LimpiezaRespuestaRefactor(resPlan, &plan)
	if plan["estado_plan_id"] != "614d3ad301c7a200482fabfd" {
		var res map[string]interface{}
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, "PUT", &res, plan); err != nil {
			return nil, errors.New("error del servicio GuardarActividad: La solicitud de actualizacion estado contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}
	}

	for key, element := range entrada {

		var respuesta map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})

		if element != "" {
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+key, &respuesta); err != nil {
				return nil, errors.New("error del servicio GuardarActividad: La solicitud de get subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
			}
			request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
			subgrupo_detalle = respuestaLimpia[0]
			actividad := make(map[string]interface{})

			if subgrupo_detalle["dato_plan"] == nil {
				actividad["index"] = maxIndex + 1
				actividad["dato"] = element
				actividad["activo"] = true
				i := strconv.Itoa(actividad["index"].(int))
				dato_plan[i] = actividad

				b, _ := json.Marshal(dato_plan)
				str := string(b)
				subgrupo_detalle["dato_plan"] = str
				if !armonizacionExecuted {
					armonizacion_dato := make(map[string]interface{})
					aux := make(map[string]interface{})
					aux["armonizacionPED"] = armonizacion
					aux["armonizacionPI"] = armonizacionPI
					armonizacion_dato[i] = aux
					c, _ := json.Marshal(armonizacion_dato)
					strArmonizacion := string(c)
					subgrupo_detalle["armonizacion_dato"] = strArmonizacion
					armonizacionExecuted = true
				}
			} else {
				dato_plan_str := subgrupo_detalle["dato_plan"].(string)
				json.Unmarshal([]byte(dato_plan_str), &dato_plan)

				actividad["index"] = maxIndex + 1
				actividad["dato"] = element
				actividad["activo"] = true
				i := strconv.Itoa(actividad["index"].(int))
				dato_plan[i] = actividad
				b, _ := json.Marshal(dato_plan)
				str := string(b)
				subgrupo_detalle["dato_plan"] = str

				if !armonizacionExecuted {
					armonizacion_dato := make(map[string]interface{})
					if subgrupo_detalle["armonizacion_dato"] != nil {
						armonizacion_dato_str := subgrupo_detalle["armonizacion_dato"].(string)
						json.Unmarshal([]byte(armonizacion_dato_str), &armonizacion_dato)
					}
					aux := make(map[string]interface{})
					aux["armonizacionPED"] = armonizacion
					aux["armonizacionPI"] = armonizacionPI
					armonizacion_dato[i] = aux
					c, _ := json.Marshal(armonizacion_dato)
					strArmonizacion := string(c)
					subgrupo_detalle["armonizacion_dato"] = strArmonizacion
					armonizacionExecuted = true

				}
			}
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
				return nil, errors.New("error del servicio GuardarActividad: La solicitud de actualizando subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
			}
		}
	}

	return entrada, nil
}

func GetPlan(id string, index string) (interface{}, error) {
	var res map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &hijos)
		formulacionhelper.Limpia()
		tree := formulacionhelper.BuildTreeFa(hijos, index)
		return tree, nil
	} else {
		return nil, errors.New("error del servicio GetPlan: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
}

func ActualizarActividad(id string, index string, datos []byte) (interface{}, error) {
	var res map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}

	_ = id
	json.Unmarshal(datos, &body)
	entrada = body["entrada"].(map[string]interface{})
	armonizacion := body["armo"]
	armonizacionPI := body["armoPI"]
	for key, element := range entrada {
		var respuesta map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var armonizacion_dato map[string]interface{}
		var id_subgrupoDetalle string
		keyStr := strings.Split(key, "_")

		if len(keyStr) > 1 && keyStr[1] == "o" {
			id_subgrupoDetalle = keyStr[0]
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
				return nil, errors.New("error del servicio ActualizarActividad: La solicitud get subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
			}
			request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

			subgrupo_detalle = respuestaLimpia[0]
			if subgrupo_detalle["dato_plan"] != nil {
				actividad := make(map[string]interface{})
				dato_plan_str := subgrupo_detalle["dato_plan"].(string)
				json.Unmarshal([]byte(dato_plan_str), &dato_plan)
				for index_actividad := range dato_plan {
					if index_actividad == index {
						aux_actividad := dato_plan[index_actividad].(map[string]interface{})
						actividad["index"] = index_actividad
						actividad["dato"] = aux_actividad["dato"]
						actividad["activo"] = aux_actividad["activo"]
						actividad["observacion"] = element

						dato_plan[index_actividad] = actividad
					}
				}
				b, _ := json.Marshal(dato_plan)
				str := string(b)
				subgrupo_detalle["dato_plan"] = str
			}

			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
				return nil, errors.New("error del servicio ActualizarActividad: La solicitud de actualizando subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
			}

			continue
		}
		id_subgrupoDetalle = key
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
			return nil, errors.New("error del servicio ActualizarActividad: La solicitud get subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

		subgrupo_detalle = respuestaLimpia[0]
		if subgrupo_detalle["armonizacion_dato"] != nil {
			dato_armonizacion_str := subgrupo_detalle["armonizacion_dato"].(string)
			json.Unmarshal([]byte(dato_armonizacion_str), &armonizacion_dato)
			if armonizacion_dato[index] != nil {
				aux := make(map[string]interface{})
				aux["armonizacionPED"] = armonizacion
				aux["armonizacionPI"] = armonizacionPI
				armonizacion_dato[index] = aux
			}
			c, _ := json.Marshal(armonizacion_dato)
			strArmonizacion := string(c)
			subgrupo_detalle["armonizacion_dato"] = strArmonizacion

		}

		nuevoDato := true
		actividad := make(map[string]interface{})

		if subgrupo_detalle["dato_plan"] != nil {
			dato_plan_str := subgrupo_detalle["dato_plan"].(string)
			json.Unmarshal([]byte(dato_plan_str), &dato_plan)

			for index_actividad := range dato_plan {
				if index_actividad == index {
					nuevoDato = false
					aux_actividad := dato_plan[index_actividad].(map[string]interface{})
					actividad["index"] = index_actividad
					actividad["dato"] = element
					actividad["activo"] = aux_actividad["activo"]
					actividad["observacion"] = aux_actividad["observacion"]
					dato_plan[index_actividad] = actividad
				}
			}
		}

		if nuevoDato {
			actividad["index"] = index
			actividad["dato"] = element
			actividad["activo"] = true
			dato_plan[index] = actividad
		}

		b, _ := json.Marshal(dato_plan)
		str := string(b)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
			return nil, errors.New("error del servicio ActualizarActividad: La solicitud de actualizando subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}

	}
	return entrada, nil
}

func DeleteActividad(id string, index string) (interface{}, error) {
	var res map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &hijos)
		formulacionhelper.RecorrerHijos(hijos, index)
		return "Actividades Inactivas", nil
	} else {
		return nil, errors.New("error del servicio DeleteActividad: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
}

func GetAllActividades(id string) (interface{}, error) {
	var res map[string]interface{}
	var hijos []map[string]interface{}
	var tabla map[string]interface{}
	var auxHijos []interface{}
	formulacionhelper.LimpiaTabla()
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &hijos)
		for i := 0; i < len(hijos); i++ {
			auxHijos = append(auxHijos, hijos[i]["_id"])
		}
		tabla = formulacionhelper.GetTabla(auxHijos)
		return tabla, nil
	} else {
		return nil, errors.New("error del servicio GetAllActividades: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
}

func GetArbolArmonizacion(datos []byte) (interface{}, error) {
	var entrada map[string][]string
	var arregloId []string
	var armonizacion []map[string]interface{}

	json.Unmarshal(datos, &entrada)
	arregloId = entrada["Data"]
	for i := 0; i < len(arregloId); i++ {
		armonizacion = append(armonizacion, formulacionhelper.GetArmonizacion(arregloId[i]))
	}
	return armonizacion, nil
}

func GuardarIdentificacion(id string, tipoIdenti string, datos []byte) (interface{}, error) {

	var entrada map[string]interface{}
	var res map[string]interface{}
	var resJ map[string]interface{}
	var respuesta []map[string]interface{}
	var idStr string
	var identificacion map[string]interface{}

	json.Unmarshal(datos, &entrada)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+id+",tipo_identificacion_id:"+tipoIdenti, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &respuesta)

		if tipoIdenti == "61897518f6fc97091727c3c3" { // ? Recurso docente unicamente
			if len(respuesta) > 0 {
				if strings.Contains(respuesta[0]["dato"].(string), "ids_detalle") {
					identificacion = respuesta[0]
					dato_json := map[string]interface{}{}
					json.Unmarshal([]byte(identificacion["dato"].(string)), &dato_json)

					iddetail := ""
					identificacionDetalle := map[string]interface{}{}
					errIdentificacionDetalle := error(nil)

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhf"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rhf"]
						identificacionDetalle = map[string]interface{}{}
						errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut)
						if errIdentificacionDetalle != nil || identificacionDetalle["Status"] != "200" {
							return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud actualizando detalle identificacion \"rhf\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
						}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud consultando detalle identificacion \"rhf\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhv_pre"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rhv_pre"]
						identificacionDetalle = map[string]interface{}{}
						errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut)
						if errIdentificacionDetalle != nil || identificacionDetalle["Status"] != "200" {
							return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud actualizando detalle identificacion \"rhv_pre\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
						}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud consultando detalle identificacion \"rhv_pre\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhv_pos"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rhv_pos"]
						identificacionDetalle = map[string]interface{}{}
						errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut)
						if errIdentificacionDetalle != nil || identificacionDetalle["Status"] != "200" {
							return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud actualizando detalle identificacion \"rhv_pos\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
						}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud consultando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rubros"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rubros"]
						identificacionDetalle = map[string]interface{}{}
						errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut)
						if errIdentificacionDetalle != nil || identificacionDetalle["Status"] != "200" {
							return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud actualizando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
						}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud consultando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rubros_pos"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rubros_pos"]
						identificacionDetalle = map[string]interface{}{}
						errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut)
						if errIdentificacionDetalle != nil || identificacionDetalle["Status"] != "200" {
							return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud actualizando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
						}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud consultando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

				} else {
					// ? -- Transicional mientras migran datos --
					identificacion = respuesta[0]
					detalles := map[string]interface{}{
						"rhf":        "",
						"rhv_pre":    "",
						"rhv_pos":    "",
						"rubros":     "",
						"rubros_pos": "",
					}
					identificacionDetalle := map[string]interface{}{}
					errIdentificacionDetalle := error(nil)

					data := map[string]interface{}{
						"dato": entrada["rhf"],
					}
					errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "201" {
						detalles["rhf"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud guardando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					data = map[string]interface{}{
						"dato": entrada["rhv_pre"],
					}
					errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "201" {
						detalles["rhv_pre"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud guardando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					data = map[string]interface{}{
						"dato": entrada["rhv_pos"],
					}
					errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "201" {
						detalles["rhv_pos"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud guardando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					data = map[string]interface{}{
						"dato": entrada["rubros"],
					}
					errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "201" {
						detalles["rubros"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud guardando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					data = map[string]interface{}{
						"dato": entrada["rubros_pos"],
					}
					errIdentificacionDetalle = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "201" {
						detalles["rubros_pos"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud guardando detalle identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

					bt, _ := json.Marshal(map[string]interface{}{"ids_detalle": detalles})
					identificacion["dato"] = string(bt)

					identificacionAns := map[string]interface{}{}
					errIdentificacion := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion/"+identificacion["_id"].(string), "PUT", &identificacionAns, identificacion)
					if errIdentificacion != nil && identificacionAns["Status"] != "200" {
						return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud actualizando identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}

				}
			} else {
				return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud no contiene un dato identificacion" + err.Error())
			}
		} else {
			jsonString, _ := json.Marshal(respuesta[0]["_id"])
			json.Unmarshal(jsonString, &idStr)

			identificacion = respuesta[0]
			b, _ := json.Marshal(entrada)
			str := string(b)

			identificacion["dato"] = str
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion/"+idStr, "PUT", &resJ, identificacion); err != nil {
				return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud actualizando identificacion contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
			}
		}

	} else {
		return nil, errors.New("error del servicio GuardarIdentificacion: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
	return "Registro de identificación", nil
}

func GetAllIdentificacion(id string, tipoIdenti string) (interface{}, error) {
	var respuesta []map[string]interface{}
	var res map[string]interface{}
	var identificacion map[string]interface{}
	var dato map[string]interface{}
	var data_identi []map[string]interface{}

	if tipoIdenti == "61897518f6fc97091727c3c3" { // ? Recurso docente unicamente
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+id+",tipo_identificacion_id:"+tipoIdenti, &res); err == nil {
			request.LimpiezaRespuestaRefactor(res, &respuesta)
			if len(respuesta) > 0 {
				if strings.Contains(respuesta[0]["dato"].(string), "ids_detalle") { // ? Nuevo método fraccionado
					identificacion = respuesta[0]
					dato_json := map[string]interface{}{}
					json.Unmarshal([]byte(identificacion["dato"].(string)), &dato_json)
					result := dato_json["ids_detalle"].(map[string]interface{})

					iddetail := ""
					identificacionDetalle := map[string]interface{}{}
					errIdentificacionDetalle := error(nil)

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhf"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						for key := range dato {
							element := dato[key].(map[string]interface{})
							if element["activo"] == true {
								datos = append(datos, element)
							}
						}
						if len(datos) > 0 {
							result["rhf"] = datos
						} else {
							result["rhf"] = "{}"
						}
					} else {
						result["rhf"] = "{}"
					}

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhv_pre"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						for key := range dato {
							element := dato[key].(map[string]interface{})
							if element["activo"] == true {
								datos = append(datos, element)
							}
						}
						if len(datos) > 0 {
							result["rhv_pre"] = datos
						} else {
							result["rhv_pre"] = "{}"
						}
					} else {
						result["rhv_pre"] = "{}"
					}

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhv_pos"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						for key := range dato {
							element := dato[key].(map[string]interface{})
							if element["activo"] == true {
								datos = append(datos, element)
							}
						}
						if len(datos) > 0 {
							result["rhv_pos"] = datos
						} else {
							result["rhv_pos"] = "{}"
						}
					} else {
						result["rhv_pos"] = "{}"
					}

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rubros"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						for key := range dato {
							element := dato[key].(map[string]interface{})
							if element["activo"] == true {
								datos = append(datos, element)
							}
						}
						if len(datos) > 0 {
							result["rubros"] = datos
						} else {
							result["rubros"] = "{}"
						}
					} else {
						result["rubros"] = "{}"
					}

					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rubros_pos"].(string)
					errIdentificacionDetalle = request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle)
					if errIdentificacionDetalle == nil && identificacionDetalle["Status"] == "200" && identificacionDetalle["Data"] != nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						for key := range dato {
							element := dato[key].(map[string]interface{})
							if element["activo"] == true {
								datos = append(datos, element)
							}
						}
						if len(datos) > 0 {
							result["rubros_pos"] = datos
						} else {
							result["rubros_pos"] = "{}"
						}
					} else {
						result["rubros_pos"] = "{}"
					}

					return result, nil

				} else {
					identificacion = respuesta[0]
					if identificacion["dato"] != nil && identificacion["dato"] != "{}" { // ? Antiguo método unificado
						result := make(map[string]interface{})
						dato_str := identificacion["dato"].(string)
						json.Unmarshal([]byte(dato_str), &dato)

						var identi map[string]interface{} = nil
						dato_aux := ""

						dato_aux = dato["rhf"].(string)
						if dato_aux == "{}" {
							result["rhf"] = "{}"
						} else {
							json.Unmarshal([]byte(dato_aux), &identi)
							for key := range identi {
								element := identi[key].(map[string]interface{})
								if element["activo"] == true {
									data_identi = append(data_identi, element)
								}
							}
							result["rhf"] = data_identi
						}

						identi = nil
						data_identi = nil

						dato_aux = dato["rhv_pre"].(string)
						if dato_aux == "{}" {
							result["rhv_pre"] = "{}"
						} else {
							json.Unmarshal([]byte(dato_aux), &identi)
							for key := range identi {
								element := identi[key].(map[string]interface{})
								if element["activo"] == true {
									data_identi = append(data_identi, element)
								}
							}
							result["rhv_pre"] = data_identi
						}

						identi = nil
						data_identi = nil

						dato_aux = dato["rhv_pos"].(string)
						if dato_aux == "{}" {
							result["rhv_pos"] = "{}"
						} else {
							json.Unmarshal([]byte(dato_aux), &identi)
							for key := range identi {
								element := identi[key].(map[string]interface{})
								if element["activo"] == true {
									data_identi = append(data_identi, element)
								}
							}
							result["rhv_pos"] = data_identi
						}

						identi = nil
						data_identi = nil

						if dato["rubros"] != nil {
							dato_aux = dato["rubros"].(string)
							if dato_aux == "{}" {
								result["rubros"] = "{}"
							} else {
								json.Unmarshal([]byte(dato_aux), &identi)
								for key := range identi {
									element := identi[key].(map[string]interface{})
									if element["activo"] == true {
										data_identi = append(data_identi, element)
									}
								}
								result["rubros"] = data_identi
							}
						}

						identi = nil
						data_identi = nil

						if dato["rubros_pos"] != nil {
							dato_aux = dato["rubros_pos"].(string)
							if dato_aux == "{}" {
								result["rubros_pos"] = "{}"
							} else {
								json.Unmarshal([]byte(dato_aux), &identi)
								for key := range identi {
									element := identi[key].(map[string]interface{})
									if element["activo"] == true {
										data_identi = append(data_identi, element)
									}
								}
								result["rubros_pos"] = data_identi
							}
						}

						return result, nil

					} else {
						return "", nil
					}
				}
			} else {
				return "", nil
			}
		} else {
			return nil, errors.New("error del servicio GetAllIdentificacion: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}
	} else {
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+id+",tipo_identificacion_id:"+tipoIdenti, &res); err == nil {
			request.LimpiezaRespuestaRefactor(res, &respuesta)
			identificacion = respuesta[0]

			if identificacion["dato"] != nil {
				dato_str := identificacion["dato"].(string)
				json.Unmarshal([]byte(dato_str), &dato)
				for key := range dato {
					element := dato[key].(map[string]interface{})
					if element["activo"] == true {
						data_identi = append(data_identi, element)
					}
				}

				sort.SliceStable(data_identi, func(i, j int) bool {
					a, _ := strconv.Atoi(data_identi[i]["index"].(string))
					b, _ := strconv.Atoi(data_identi[j]["index"].(string))
					return a < b
				})

				return data_identi, nil

			} else {
				return "", nil
			}

		} else {
			return nil, errors.New("error del servicio GetAllIdentificacion: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}

	}
}

func DeleteIdentificacion(id string, idTipo string, index string) (interface{}, error) {
	var idStr string
	var res map[string]interface{}
	var respuesta []map[string]interface{}
	var identificacion map[string]interface{}
	var dato map[string]interface{}
	var resJ map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+id+",tipo_identificacion_id:"+idTipo, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &respuesta)
		identificacion = respuesta[0]

		jsonString, _ := json.Marshal(respuesta[0]["_id"])
		json.Unmarshal(jsonString, &idStr)

		if identificacion["dato"] != nil {
			dato_str := identificacion["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato)
			for key := range dato {
				intVar, _ := strconv.Atoi(key)
				intVar = intVar + 1
				strr := strconv.Itoa(intVar)
				if strr == index {
					element := dato[key].(map[string]interface{})
					element["activo"] = false
					dato[key] = element
				}
			}
			b, _ := json.Marshal(dato)
			str := string(b)
			identificacion["dato"] = str
		}
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion/"+idStr, "PUT", &resJ, identificacion); err != nil {
			return nil, errors.New("error del servicio DeleteIdentificacion: La solicitud eliminando identificacion \"idStr\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}
		return "Identificación Inactiva", nil
	} else {
		return nil, errors.New("error del servicio DeleteIdentificacion: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
}

func VersionarPlan(id string) (interface{}, error) {
	var respuesta map[string]interface{}
	var respuestaHijos map[string]interface{}
	var respuestaIdentificacion map[string]interface{}
	var hijos []map[string]interface{}
	var identificaciones []map[string]interface{}
	var planPadre map[string]interface{}
	var respuestaPost map[string]interface{}
	var planVersionado map[string]interface{}
	plan := make(map[string]interface{})

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &respuesta); err == nil {

		request.LimpiezaRespuestaRefactor(respuesta, &planPadre)

		plan["nombre"] = planPadre["nombre"].(string)
		plan["descripcion"] = planPadre["descripcion"].(string)
		plan["tipo_plan_id"] = planPadre["tipo_plan_id"].(string)
		plan["aplicativo_id"] = planPadre["aplicativo_id"].(string)
		plan["activo"] = planPadre["activo"]
		plan["formato"] = false
		plan["vigencia"] = planPadre["vigencia"].(string)
		plan["dependencia_id"] = planPadre["dependencia_id"].(string)
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"
		plan["padre_plan_id"] = id

		if value, ok := planPadre["formato_id"].(string); ok {
			plan["formato_id"] = value
		}

		if _, ok := planPadre["nueva_estructura"]; ok {
			plan["nueva_estructura"] = true
		}

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan", "POST", &respuestaPost, plan); err != nil {
			return nil, errors.New("error del servicio VersionarPlan: La solicitud versionando plan \"plan[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}
		planVersionado = respuestaPost["Data"].(map[string]interface{})

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+id, &respuestaIdentificacion); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaIdentificacion, &identificaciones)
			if len(identificaciones) != 0 {
				formulacionhelper.VersionarIdentificaciones(identificaciones, planVersionado["_id"].(string))
			}
		} else {
			return nil, errors.New("error del servicio VersionarPlan: La solicitud obteniendo identificación contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &respuestaHijos); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
			formulacionhelper.VersionarHijos(hijos, planVersionado["_id"].(string))
		} else {
			return nil, errors.New("error del servicio VersionarPlan: La solicitud obteniendo subgrupo detalle contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}

		var resPadres map[string]interface{}
		var planesPadre []map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=dependencia_id:"+plan["dependencia_id"].(string)+",vigencia:"+plan["vigencia"].(string)+",formato:false,nombre:"+url.QueryEscape(plan["nombre"].(string)), &resPadres); err == nil {
			request.LimpiezaRespuestaRefactor(resPadres, &planesPadre)
			for _, padre := range planesPadre {
				var resActualizacion map[string]interface{}
				if padre["_id"].(string) != planVersionado["_id"].(string) {
					padre["activo"] = false
					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+padre["_id"].(string), "PUT", &resActualizacion, padre); err != nil {
						return nil, errors.New("error del servicio VersionarPlan: La solicitud versionando plan contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
					}
				}
			}
		} else {
			return nil, errors.New("error del servicio VersionarPlan: La solicitud obteniendo plan contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}
		return planVersionado, nil
	} else {
		return nil, errors.New("error del servicio VersionarPlan: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
}

func GetPlanVersiones(unidad string, vigencia string, nombre string) (interface{}, error) {
	var respuesta map[string]interface{}
	var versiones []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=dependencia_id:"+unidad+",vigencia:"+vigencia+",formato:false,nombre:"+nombre, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &versiones)
		versionesOrdenadas := formulacionhelper.OrdenarVersiones(versiones)
		for _, version := range versionesOrdenadas {
			padre_plan_id, ok := version["padre_plan_id"]
			if ok {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/reformulacion?query=plan_id:"+padre_plan_id.(string), &respuesta); err == nil {
					if len(respuesta["Data"].([]interface{})) > 0 {
						version["reformulacion"] = true
					} else {
						version["reformulacion"] = false
					}
				} else {
					return nil, errors.New("error del servicio GetPlanVersiones: No se pudo hacer la solicitud a planes_crud" + err.Error())
				}
			} else {
				version["reformulacion"] = false
			}
		}
		return versionesOrdenadas, nil
	} else {
		return nil, errors.New("error del servicio GetPlanVersiones: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
}

func PonderacionActividades(plan string) (interface{}, error) {
	var respuesta map[string]interface{}
	var respuestaDetalle map[string]interface{}
	var respuestaLimpiaDetalle []map[string]interface{}
	var subgrupoDetalle map[string]interface{}
	var hijos []map[string]interface{}
	var hijosFiltrado []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+plan, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &hijos)
		for i := 0; i < len(hijos); i++ {
			if hijos[i]["activo"] == true {
				hijosFiltrado = append(hijosFiltrado, hijos[i])
			}
		}

		for i := 0; i < len(hijosFiltrado); i++ {
			if strings.Contains(strings.ToUpper(hijosFiltrado[i]["nombre"].(string)), "PONDERACIÓN") && strings.Contains(strings.ToUpper(hijosFiltrado[i]["nombre"].(string)), "ACTIVIDAD") || strings.Contains(strings.ToUpper(hijosFiltrado[i]["nombre"].(string)), "PONDERACIÓN") {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+hijosFiltrado[i]["_id"].(string), &respuestaDetalle); err == nil {
					request.LimpiezaRespuestaRefactor(respuestaDetalle, &respuestaLimpiaDetalle)
					subgrupoDetalle = respuestaLimpiaDetalle[0]

					if subgrupoDetalle["dato_plan"] != nil {
						var suma float64 = 0
						datoPlan := make(map[string]map[string]interface{})
						json.Unmarshal([]byte(subgrupoDetalle["dato_plan"].(string)), &datoPlan)

						ponderacionActividades := make(map[string]interface{})

						for j, dato := range datoPlan {
							if dato["activo"] != false && len(dato) != 0 {
								ponderacionActividades["Actividad "+(j)] = dato["dato"]
								suma += dato["dato"].(float64)
								suma = math.Round(suma*100) / 100
							}
						}

						ponderacionActividades["Total"] = suma
						return ponderacionActividades, nil
					}
				} else {
					return nil, errors.New("error del servicio PonderacionActividades: La solicitud subgrupo_detalle plan \"plan\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
				}
			}
		}
	} else {
		return nil, errors.New("error del servicio PonderacionActividades: La solicitud subgrupo_hijos plan \"plan\" contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
	return nil, errors.New("error del servicio PonderacionActividades: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido")
}

func GetRubros() (interface{}, error) {
	var respuesta map[string]interface{}
	var rubros []interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanCuentasService")+"/arbol_rubro", &respuesta); err != nil {
		return nil, errors.New("error del servicio GetRubros: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	} else {
		rubros = respuesta["Body"].([]interface{})
		for i := 0; i < len(rubros); i++ {
			if strings.ToUpper(rubros[i].(map[string]interface{})["Nombre"].(string)) == "GASTOS" {
				aux := rubros[i].(map[string]interface{})
				hojas := formulacionhelper.GetHijosRubro(aux["Hijos"].([]interface{}))
				if len(hojas) != 0 {
					return hojas, nil
				} else {
					return "", nil
				}
			}
		}
	}
	return nil, errors.New("error del servicio GetRubros: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido")
}

func GetUnidades() (interface{}, error) {
	var respuesta []map[string]interface{}
	var unidades []map[string]interface{}

	tiposDependencia := []string{"2", "3", "4", "5", "6", "7", "8", "11", "13", "15", "28", "33"}
	for _, tipoDep := range tiposDependencia {
		respuesta = nil
		err := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"/dependencia_tipo_dependencia?limit=0&query=TipoDependenciaId:"+tipoDep, &respuesta)
		if err != nil {
			return nil, errors.New("error del servicio GetUnidades: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}
		for i := 0; i < len(respuesta); i++ {
			if tipoDep == "15" {
				aux := respuesta[i]["DependenciaId"]
				if strings.Contains(aux.(map[string]interface{})["Nombre"].(string), "DOCTORADO") {
					aux := respuesta[i]["DependenciaId"].(map[string]interface{})
					aux["TipoDependencia"] = respuesta[i]["TipoDependenciaId"]
					unidades = append(unidades, aux)
				}
				continue
			}

			aux := respuesta[i]["DependenciaId"].(map[string]interface{})
			aux["TipoDependencia"] = respuesta[i]["TipoDependenciaId"]
			unidades = append(unidades, aux)
		}
	}

	tipoDependenciaDependencia := map[string][]string{
		"10": {"92", "96", "97", "209"},
		"14": {"42", "171"},
		"33": {"222"},
	}
	for tipoDep, deps := range tipoDependenciaDependencia {
		for _, dep := range deps {
			respuesta = nil
			err := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"/dependencia_tipo_dependencia?limit=0&query=TipoDependenciaId:"+tipoDep+",DependenciaId:"+dep, &respuesta)
			if err != nil {
				return nil, errors.New("error del servicio GetUnidades: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
			}
			for i := 0; i < len(respuesta); i++ {
				aux := respuesta[i]["DependenciaId"].(map[string]interface{})
				aux["TipoDependencia"] = respuesta[i]["TipoDependenciaId"]
				unidades = append(unidades, aux)
			}
		}
	}

	return unidades, nil
}

func VinculacionTercero(terceroId string) (interface{}, error) {
	var vinculaciones []models.Vinculacion
	var resultado []models.Vinculacion
	if err := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"/vinculacion?query=Activo:true,TerceroPrincipalId:"+terceroId, &vinculaciones); err != nil {
		return nil, errors.New("error del servicio VinculacionTercero: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	} else {
		for i := 0; i < len(vinculaciones); i++ {
			if vinculaciones[i].CargoId == 319 || vinculaciones[i].CargoId == 312 || vinculaciones[i].CargoId == 320 {
				resultado = append(resultado, vinculaciones[i])
			}
		}
		return resultado, nil
	}
}

func Planes() (interface{}, error) {
	var respuesta map[string]interface{}
	var planes []map[string]interface{}
	var planesPED []map[string]interface{}
	var planesPI []map[string]interface{}
	var arregloPlanes []map[string]interface{}
	var auxArregloPlanes []map[string]interface{}
	var finalRes interface{}
	var mutex sync.Mutex
	wge := new(errgroup.Group)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=formato:true", &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &planes)

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=tipo_plan_id:6239117116511e20405d408b", &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &planesPI)
		} else {
			return nil, errors.New("error del servicio Planes: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=tipo_plan_id:616513b91634adfaffed52bf", &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &planesPED)
		} else {
			return nil, errors.New("error del servicio Planes: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}

		auxArregloPlanes = append(auxArregloPlanes, planes...)
		auxArregloPlanes = append(auxArregloPlanes, planesPI...)
		auxArregloPlanes = append(auxArregloPlanes, planesPED...)

		for i := 0; i < len(auxArregloPlanes); i++ {
			i := i
			wge.Go(func() error {
				plan := auxArregloPlanes[i]
				tipoPlanId := plan["tipo_plan_id"].(string)

				var res map[string]interface{}
				var tipoPlanes []map[string]interface{}
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/tipo-plan?query=_id:"+tipoPlanId, &res); err == nil {
					request.LimpiezaRespuestaRefactor(res, &tipoPlanes)
					tipoPlan := tipoPlanes[0]
					nombreTipoPlan := tipoPlan["nombre"]
					planesTipo := make(map[string]interface{})
					planesTipo["_id"] = plan["_id"]
					planesTipo["nombre"] = plan["nombre"]
					planesTipo["descripcion"] = plan["descripcion"]
					planesTipo["tipo_plan_id"] = tipoPlanId
					planesTipo["formato"] = plan["formato"]
					planesTipo["vigencia"] = plan["vigencia"]
					planesTipo["dependencia_id"] = plan["dependencia_id"]
					planesTipo["aplicativo_id"] = plan["aplicativo_id"]
					planesTipo["activo"] = plan["activo"]
					planesTipo["nombre_tipo_plan"] = nombreTipoPlan

					mutex.Lock()
					defer mutex.Unlock()
					arregloPlanes = append(arregloPlanes, planesTipo)

					if arregloPlanes != nil {
						// Mapa auxiliar para rastrear IDs únicos
						idMap := make(map[string]bool)

						// Nuevo slice para almacenar elementos únicos
						uniquePlanes := []map[string]interface{}{}

						for _, plan := range arregloPlanes {
							id := plan["_id"].(string)
							if !idMap[id] {
								// Si el ID no ha sido visto antes, añadirlo al mapa y al slice único
								idMap[id] = true
								uniquePlanes = append(uniquePlanes, plan)
							}
						}
						finalRes = uniquePlanes
					}

				} else {
					return errors.New("error del servicio Planes: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
				}
				return nil
			})
		}
		if err := wge.Wait(); err != nil {
			return nil, errors.New("error del servicio Planes: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}

	} else {
		return nil, errors.New("error del servicio Planes: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}

	return finalRes, nil
}

func VerificarIdentificaciones(id string) (interface{}, error) {
	var respuesta map[string]interface{}
	var respuestaPlan map[string]interface{}
	var respuestaDependencia []map[string]interface{}
	var dependencia map[string]interface{}
	var plan map[string]interface{}
	var identificaciones []map[string]interface{}
	var bandera bool

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &respuestaPlan); err == nil {
		request.LimpiezaRespuestaRefactor(respuestaPlan, &plan)

		if err := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"/dependencia_tipo_dependencia?query=DependenciaId:"+plan["dependencia_id"].(string), &respuestaDependencia); err == nil {
			dependencia = respuestaDependencia[0]

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+id, &respuesta); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &identificaciones)

				tipoDependencia := dependencia["TipoDependenciaId"].(map[string]interface{})
				id := dependencia["DependenciaId"].(map[string]interface{})["Id"]
				if (tipoDependencia["Id"] == 2.00 || id == 67.00) && id != 8.0 {
					bandera = formulacionhelper.VerificarDataIdentificaciones(identificaciones, "facultad")
				} else {
					bandera = formulacionhelper.VerificarDataIdentificaciones(identificaciones, "unidad")
				}

			} else {
				return nil, errors.New("error del servicio VerificarIdentificaciones: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
			}
		} else {
			return nil, errors.New("error del servicio VerificarIdentificaciones: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
		}
	} else {
		return nil, errors.New("error del servicio VerificarIdentificaciones: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}

	return bandera, nil
}

func PlanesEnFormulacion() (interface{}, error) {
	if resumenPlanesActivos, err := formulacionhelper.ObtenerPlanesFormulacion(); err != nil {
		return nil, errors.New("error del servicio PlanesEnFormulacion consultando ObtenerPlanesFormulacion: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido")
	} else {
		return resumenPlanesActivos, nil
	}
}

func CalculosDocentes(datos []byte) (interface{}, error) {
	//Obtener respuesta del body
	var body map[string]interface{}
	if err := json.Unmarshal(datos, &body); err != nil {
		return nil, errors.New("error del servicio CalculosDocentes: Error al decodificar el cuerpo de la solicitud" + err.Error())
	}

	// Obtener Desagregado
	body["vigencia"] = body["vigencia"].(float64) - 1
	bodyResolucionesDocente := formulacionhelper.ConstruirCuerpoRD(body)
	respuestaPost, err := formulacionhelper.GetDesagregado(bodyResolucionesDocente)
	if err != nil {
		return nil, errors.New("error del servicio CalculosDocentes: Error al obtener desagregado" + err.Error())
	}
	result := respuestaPost["Data"].([]interface{})

	//Peticion GET hacia Parametros Service
	vigenciaStr := strconv.FormatFloat(body["vigencia"].(float64), 'f', 0, 64)
	salarioMinimo, err := formulacionhelper.GetSalarioMinimo(vigenciaStr)
	if err != nil {
		return nil, errors.New("error del servicio CalculosDocentes: Error al obtener salario minimo" + err.Error())
	}

	// Objeto para hacer los cálculos necesarios
	data := body
	data["resolucionDocente"] = result[0].(map[string]interface{})
	data["salarioMinimo"] = salarioMinimo["Valor"]
	delete(body, "vigencia")
	delete(body, "categoria")
	delete(body, "tipo")

	// Realizar los calculos
	dataFinal, err := formulacionhelper.GetCalculos(data)
	if err != nil {
		return nil, errors.New("error del servicio CalculosDocentes: Error al intentar realizar los calculos" + err.Error())
	}

	return dataFinal, nil
}

func EstructuraPlanes(id string) (interface{}, error) {
	//Obtener plantilla por id
	plantilla, err := formulacionhelper.GetPlantilla(id)
	if err != nil {
		return nil, errors.New("error del servicio EstructuraPlanes: Error al obtener plantilla" + err.Error())
	}

	//Obtener los planes en estado "En formulacion" asociados a la plantilla
	planes, err := formulacionhelper.GetPlanesPorNombre(plantilla["nombre"].(string))
	if err != nil {
		return nil, errors.New("error del servicio EstructuraPlanes: Error al obtener planes asociados a plantilla" + err.Error())
	}

	//Obtener el formato de la plantilla
	formatoPLantilla, err := formulacionhelper.GetFormato(id)
	if err != nil {
		return nil, errors.New("error del servicio EstructuraPlanes: Error al obtener formato de plantilla" + err.Error())
	}

	//Obtener lista plana del formato
	listaPlantilla, err := formulacionhelper.ConvArbolAListaPlana(formatoPLantilla[0], id, true)
	if err != nil {
		return nil, errors.New("error del servicio EstructuraPlanes: Error al obtener el valor de referencia" + err.Error())
	}

	//Obtener los formatos de los planes y comparar con el formato de la plantilla
	for _, plan := range planes {
		planId := plan["_id"].(string)
		formatoPlan, err := formulacionhelper.GetFormato(planId)
		if err != nil {
			return nil, errors.New("error del servicio EstructuraPlanes: Error al obtener formato de plan" + err.Error())
		}
		listaPlan, err := formulacionhelper.ConvArbolAListaPlana(formatoPlan[0], planId, false)
		if err == nil {
			formulacionhelper.ActualizarEstructuraPlan(listaPlantilla, listaPlan, planId)
		}
	}

	return "La estructura de los planes fue actualizada correctamente", nil
}

func DefinirFechas(datos []byte) (interface{}, error) {
	var body map[string]interface{}
	var res interface{}

	// Decodificar JSON desde el cuerpo de la solicitud
	err := json.Unmarshal(datos, &body)
	if err != nil {
		return nil, errors.New("error del servicio DefinirFechasFormulacionSeguimiento: Error al decodificar JSON" + err.Error())
	}
	res = formulacionhelper.DefinirFechas(body)
	return res, nil
}

func GetPlanesUnidadesComun(id string, datos []byte) (interface{}, error) {
	var body_unidades map[string]interface{}
	var periodo_seguimiento map[string]interface{}
	var res map[string]interface{}
	err := json.Unmarshal(datos, &body_unidades)
	if err != nil {
		return nil, errors.New("error del servicio GetPlanesUnidadesComun: Error al decodificar el cuerpo de la solicitud" + err.Error())
	}
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/periodo-seguimiento/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &periodo_seguimiento)
		unidadesValidadas := formulacionhelper.ValidarUnidadesPlanes(periodo_seguimiento, body_unidades)

		// Verificar el resultado
		if len(unidadesValidadas) > 0 {
			planesInteres := periodo_seguimiento["planes_interes"]
			return planesInteres, nil
		} else {
			// c.Data["json"] = map[string]interface{}{"Success": true, "Status": "404", "Message": "Successful", "Data": "Not found"}
			return "No hay unidades en la intersección", nil
		}
		// c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": planesInteres}
	} else {
		return nil, errors.New("error del servicio GetPlanesUnidadesComun: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	}
}

func GetPlanesDeAccion() (interface{}, error) {
	if datos, err := formulacionhelper.ObtenerPlanesAccion(); err == nil {
		return datos, nil
	} else {
		return nil, errors.New("error del servicio GetPlanesDeAccion" + err.Error())
	}
}

func GetPlanesDeAccionPorUnidad(unidadID string) (interface{}, error) {
	var planes []map[string]interface{}

	if datos, err := formulacionhelper.ObtenerPlanesAccion(); err == nil {
		for _, plan := range datos {
			if plan["dependencia_id"] == unidadID {
				planes = append(planes, plan)
			}
		}
		return planes, nil
	} else {
		return nil, errors.New("error del servicio GetPlanesDeAccion" + err.Error())
	}
}

func GetVinculacionTerceroByEmail(terceroEmail string) (interface{}, error) {
	var vinculaciones []models.Vinculacion

	idNoRegistra, idJefeOficina, idAsistenteDependencia, err := formulacionhelper.ObtenerIdParametros()
	if err != nil {
		panic(map[string]interface{}{"funcion": "VinculacionTercero", "err": "Error get parametros", "status": "400", "log": err})
	}

	if err := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"/vinculacion?query=Activo:true,TerceroPrincipalId__UsuarioWSO2:"+terceroEmail, &vinculaciones); err != nil {
		return nil, errors.New("error del servicio VinculacionTercero: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido" + err.Error())
	} else {
		var vinculacionesResponse []models.Vinculacion
		for i := 0; i < len(vinculaciones); i++ {
			if vinculaciones[i].CargoId == int(idJefeOficina) || vinculaciones[i].CargoId == int(idAsistenteDependencia) || vinculaciones[i].CargoId == int(idNoRegistra) {
				vinculacionesResponse = append(vinculacionesResponse, vinculaciones[i])
			}
		}
		if len(vinculacionesResponse) == 0 {
			return nil, errors.New("error del servicio VinculacionTercero: No se encontró la vinculación")
		}
		return vinculacionesResponse, nil
	}
}

func obtenerCorreoPlaneacion() (string, error) {
	var respuestaPeticion map[string]interface{}
	var correoPlaneacion map[string]interface{}
	baseURL := "http://" + beego.AppConfig.String("ParametrosService") + "/parametro_periodo?query="
	err := request.GetJson(baseURL+"ParametroId.CodigoAbreviacion:CORREO_OAP,ParametroId.TipoParametroId.CodigoAbreviacion:P_SISGPLAN,Activo:true", &respuestaPeticion)
	jsonData, ok := respuestaPeticion["Data"].([]interface{})[0].(map[string]interface{})["Valor"].(string)

	if err != nil || !ok {
		return "", fmt.Errorf("no se pudo obtener el correo de la Oficina de Asesora de Planeación del módulo de parámetros, comuníquese con computo@udistrital.edu.co")
	}
	err1 := json.Unmarshal([]byte(jsonData), &correoPlaneacion)
	if err1 != nil {
		return "", fmt.Errorf("Error al deserializar el JSON de Correo Oficina Planeacion: ", err)
	}
	return correoPlaneacion["Valor"].(string), nil
}

func CambioCargoIdVinculacionTercero(idVinculacion string, cuerpo []byte) (*models.Vinculacion, error) {
	const ROL_ASISTENTE_PLANEACION = "ASISTENTE_PLANEACION"
	var vinculacionPlaneacion, rolAsistentePlaneacion bool
	var vinculacion []models.Vinculacion
	var body map[string]interface{}
	json.Unmarshal(cuerpo, &body)

	correoPlaneacion, errorPeticion := obtenerCorreoPlaneacion()
	if errorPeticion != nil {
		panic(map[string]interface{}{"funcion": "CambioCargoIdVinculacionTercero", "err": errorPeticion.Error(), "status": "400", "log": errorPeticion})
	}

	idNoRegistra, idJefeOficina, idAsistenteDependencia, err := formulacionhelper.ObtenerIdParametros()
	if err != nil {
		return nil, errors.New("error del servicio CambioCargoIdVinculacionTercero: Error al obtener parametros" + err.Error())
	}

	err = request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"/vinculacion?query=Activo:true,Id:"+idVinculacion, &vinculacion)
	if err != nil || vinculacion[0].CargoId == 0 {
		return nil, errors.New("error del servicio CambioCargoIdVinculacionTercero: Error al obtener vinculacion" + err.Error())
	}

	vinculaciones := body["user"].(map[string]interface{})["Vinculacion"].([]interface{})
	vinculacionSeleccionada := body["user"].(map[string]interface{})["VinculacionSeleccionadaId"]

	for _, vinculacion := range vinculaciones {
		if vinculacion.(map[string]interface{})["Id"].(float64) == vinculacionSeleccionada {
			if DependenciaCorreo, ok := vinculacion.(map[string]interface{})["DependenciaCorreo"].(string); ok {
				vinculacionPlaneacion = DependenciaCorreo == correoPlaneacion
				rolAsistentePlaneacion = (body["rol"].(string) == ROL_ASISTENTE_PLANEACION)
				if rolAsistentePlaneacion && !vinculacionPlaneacion && body["vincular"] == true {
					return nil, fmt.Errorf("El usuario no puede tener el rol de %s si no pertenece a la Oficina de Asesora de Planeación", ROL_ASISTENTE_PLANEACION)
				}
			} else {
				return nil, fmt.Errorf("no se pudo obtener la Dependencia asociada a la vinculación")
			}
		}
	}

	if vinculacion[0].CargoId == int(idJefeOficina) || vinculacion[0].CargoId == int(idAsistenteDependencia) || vinculacion[0].CargoId == int(idNoRegistra) {
		if body["vincular"] == true {
			vinculacion[0].CargoId = int(idAsistenteDependencia)
		} else {
			vinculacion[0].CargoId = int(idNoRegistra)
		}
		vinculacion[0].FechaCreacion, err = formulacionhelper.FormatearFecha(vinculacion[0].FechaCreacion)
		if err != nil {
			return nil, errors.New("error del servicio CambioCargoIdVinculacionTercero: Error al formatear fecha" + err.Error())
		}
		vinculacion[0].FechaModificacion = time.Now().Format(time.RFC3339)
		if err := request.SendJson("http://"+beego.AppConfig.String("TercerosService")+"/vinculacion/"+idVinculacion, "PUT", &vinculacion[0], vinculacion[0]); err != nil {
			return nil, errors.New("error del servicio CambioCargoIdVinculacionTercero: Error al actualizar vinculacion" + err.Error())
		}
		return &vinculacion[0], nil
	}

	return nil, errors.New("No se encontró la vinculación")
}

func VinculacionTerceroByIdentificacion(identificacionTercero string) (interface{}, error) {
	var vinculaciones []models.Vinculacion
	var tercero []models.DatosIdentificacion

	idNoRegistra, idJefeOficina, idAsistenteDependencia, err := formulacionhelper.ObtenerIdParametros()
	if err != nil {
		return nil, errors.New("error del servicio VinculacionTerceroByIdentificacion: Error get parametros" + err.Error())
	}

	s := "Numero:" + identificacionTercero + ",Activo:true"
	if err := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"/datos_identificacion?query="+url.QueryEscape(s), &tercero); err != nil || tercero[0].TerceroID.ID == 0 {
		return nil, errors.New("error del servicio VinculacionTerceroByIdentificacion: Error get tercero" + err.Error())
	}

	TerceroIdStr := fmt.Sprintf("%d", tercero[0].TerceroID.ID)
	if err := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"/vinculacion?query=Activo:true,TerceroPrincipalId:"+TerceroIdStr, &vinculaciones); err != nil {
		return nil, errors.New("error del servicio VinculacionTerceroByIdentificacion: Error get vinculacion" + err.Error())
	} else {
		var vinculacionesResponse []models.Vinculacion
		for i := 0; i < len(vinculaciones); i++ {
			if vinculaciones[i].CargoId == int(idJefeOficina) || vinculaciones[i].CargoId == int(idAsistenteDependencia) || vinculaciones[i].CargoId == int(idNoRegistra) {
				vinculacionesResponse = append(vinculacionesResponse, vinculaciones[i])
			}
		}
		return vinculacionesResponse, nil
	}
}
