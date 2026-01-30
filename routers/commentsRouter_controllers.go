package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GuardarActividad",
			Router:           "/actividad/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GetAllActividades",
			Router:           "/actividad/:id/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "ActualizarActividad",
			Router:           "/actividad/:id/:index",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "DeleteActividad",
			Router:           "/actividad/:id/:index/desactivar",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "PonderacionActividades",
			Router:           "/actividad/ponderacion/:plan",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "CalculosDocentes",
			Router:           "/calculos_docentes",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "CambioCargoIdVinculacionTercero",
			Router:           "/cargo_vinculacion/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "ClonarPI_PED",
			Router:           "/clonar-pi-ped/:id",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "EstructuraPlanes",
			Router:           "/estructura_planes/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "ClonarFormato",
			Router:           "/formato/:id/clonar",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GetArbolArmonizacion",
			Router:           "/get_arbol_armonizacion/:id/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GetPlanesUnidadesComun",
			Router:           "/get_planes_unidades_comun/:id",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "HabilitarFechas",
			Router:           "/habilitar_fechas",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GetAllIdentificacion",
			Router:           "/identificacion/:id/:idTipo",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GuardarIdentificacion",
			Router:           "/identificacion/:id/:idTipo",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "DeleteIdentificacion",
			Router:           "/identificacion/:id/:idTipo/:index",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "VerificarIdentificaciones",
			Router:           "/identificacion/verificacion/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GetPlan",
			Router:           "/plan/:id/:index",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "VersionarPlan",
			Router:           "/plan/:id/versionar",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GetPlanVersiones",
			Router:           "/plan/versiones/:unidad/:vigencia/:nombre",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "Planes",
			Router:           "/planes",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "PlanesDeAccion",
			Router:           "/planes_accion",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "PlanesDeAccionPorUnidad",
			Router:           "/planes_accion/:unidad_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "PlanesEnFormulacion",
			Router:           "/planes_formulacion",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GetRubros",
			Router:           "/rubros",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "VinculacionTercero",
			Router:           "/tercero/:tercero_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "GetUnidades",
			Router:           "/unidades",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "VinculacionTerceroByEmail",
			Router:           "/vinculacion_tercero_email/:tercero_email",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
		beego.ControllerComments{
			Method:           "VinculacionTerceroByIdentificacion",
			Router:           "/vinculacion_tercero_identificacion/:identificacion",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "ActualizarActividad",
			Router:           "/actividad/:id/:index",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "ActualizarTablaActividad",
			Router:           "/actividad/:id/:index/tabla",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "ActualizarProyectoGeneral",
			Router:           "/actualizar_proyecto/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "ArmonizarInversion",
			Router:           "/armonizar/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "GuardarDocumentos",
			Router:           "/documentos",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "ProgMagnitudesPlan",
			Router:           "/magnitudes/:id/:index",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "MagnitudesProgramadas",
			Router:           "/magnitudes/:id/:indexMeta",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "VerificarMagnitudesProgramadas",
			Router:           "/magnitudes/:id/verificar",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "AllMetasPlan",
			Router:           "/metas/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "GuardarMeta",
			Router:           "/metas/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "ActualizarMetaPlan",
			Router:           "/metas/:id/:index",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "InactivarMeta",
			Router:           "/metas/:id/:index/inactivar",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "CrearGrupoMeta",
			Router:           "/metas/grupo",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "ActualizarPresupuestoMeta",
			Router:           "/metas/presupuestos/:id/:index",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "CrearPlan",
			Router:           "/planes",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "GetPlanId",
			Router:           "/planes/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "GetPlan",
			Router:           "/planes/:id/informacion",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "VersionarPlan",
			Router:           "/planes/:id/versionar",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "GetProyectoId",
			Router:           "/proyecto/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "AddProyecto",
			Router:           "/proyectos",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "GetAllProyectos",
			Router:           "/proyectos/:aplicativo_id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "EditProyecto",
			Router:           "/proyectos/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "GetMetasProyect",
			Router:           "/proyectos/:id/metas",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
		beego.ControllerComments{
			Method:           "ActualizarSubgrupoDetalle",
			Router:           "/subgrupo/detalle/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
