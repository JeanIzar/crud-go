package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JeanIzar/crud-go/src/config"
	"github.com/JeanIzar/crud-go/src/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	IDCompra            uint      `json:"idcompra" gorm:"primaryKey"`
	Compracod           int       `json:"compracod"`
	Codigo              string    `json:"codigo" sql:"varchar(45)"`
	Compra_fecha_emi    time.Time `json:"comprafechaemi"`
	Compra_fecha_venc   time.Time `json:"comprafechavenc"`
	Compra_serie        string    `json:"compraserie" sql:"varchar(45)"`
	Compra_numer        string    `json:"compranumer" sql:"varchar(45)"`
	Compra_Descrip      string    `json:"compradescrip" sql:"type:text"`
	Compra_Obser        string    `json:"compraobser" sql:"type:text"`
	Compra_total_item   int       `json:"compratotalitem"`
	Compra_tipo_cambio  float32   `json:"compratipocambio" sql:"type:decimal(20,3)"`
	Compra_vbb          float32   `json:"compra_vbb" sql:"type:decimal(20,2)"`
	Compra_Vvitem       float32   `json:"compra_vvitem" sql:"type:decimal(20,2)"`
	Compra_Dscto_global float32   `json:"compra_dsctoglobal" sql:"decimal(20,2)"`
	Compra_Isc          float32   `json:"compra_isc" sql:"decimal(20,2)"`
	Compra_Vv           float32   `json:"compra_vv" sql:"decimal(20,2)"`
	Compra_Pv           float32   `json:"compra_pv" sql:"decimal(20,2)"`
	Compra_Cargo        float32   `json:"compra_cargo" sql:"decimal(20,2)"`
	Compra_Tot_v        float32   `json:"compra_totv" sql:"decimal(20,2)"`
	Compra_Tot_v_mn     float32   `json:"compra_totvmn" sql:"decimal(20,2)"`
	Compra_Oper_A       float32   `json:"compra_oa" sql:"decimal(20,2)"`
	Compra_Oper_I       float32   `json:"compra_oi" sql:"decimal(20,2)"`
	Compra_Oper_E       float32   `json:"compra_oe" sql:"decimal(20,2)"`
	Compra_Cant_Pago    int       `json:"compra_cantpago"`
	IDEstado            int       `json:"idestado"`
	IDProveedor         int       `json:"idproveedor"`
	ID_Tipodocumento    int       `json:"id_tipodocumento"`
	ID_Periodo          int       `json:"id_periodo"`
	ID_Moneda           int       `json:"id_moneda"`
	ID_Empresa          int       `json:"id_empresa"`
	ID_Local            int       `json:"id_local"`
	ID_Centro_Costo     int       `json:"id_centrocosto"`
	ID_Persona          int       `json:"id_persona"`
	Compra_Cant_Asiento int       `json:"compra_cantasiento"`
	Cod_Usua_Reg        int       `json:"cod_us_reg"`
	Cod_Usua_Act        int       `json:"cod_us_act"`
}

type todoResponse struct {
	todoRequest
	IDCompra            uint      `json:"idcompra" gorm:"primaryKey"`
	Compracod           int       `json:"compracod"`
	Codigo              string    `json:"codigo" sql:"varchar(45)"`
	Compra_fecha_emi    time.Time `json:"comprafechaemi"`
	Compra_fecha_venc   time.Time `json:"comprafechavenc"`
	Compra_serie        string    `json:"compraserie" sql:"varchar(45)"`
	Compra_numer        string    `json:"compranumer" sql:"varchar(45)"`
	Compra_Descrip      string    `json:"compradescrip" sql:"type:text"`
	Compra_Obser        string    `json:"compraobser" sql:"type:text"`
	Compra_total_item   int       `json:"compratotalitem"`
	Compra_tipo_cambio  float32   `json:"compratipocambio" sql:"type:decimal(20,3)"`
	Compra_vbb          float32   `json:"compra_vbb" sql:"type:decimal(20,2)"`
	Compra_Vvitem       float32   `json:"compra_vvitem" sql:"type:decimal(20,2)"`
	Compra_Dscto_global float32   `json:"compra_dsctoglobal" sql:"decimal(20,2)"`
	Compra_Isc          float32   `json:"compra_isc" sql:"decimal(20,2)"`
	Compra_Vv           float32   `json:"compra_vv" sql:"decimal(20,2)"`
	Compra_Pv           float32   `json:"compra_pv" sql:"decimal(20,2)"`
	Compra_Cargo        float32   `json:"compra_cargo" sql:"decimal(20,2)"`
	Compra_Tot_v        float32   `json:"compra_totv" sql:"decimal(20,2)"`
	Compra_Tot_v_mn     float32   `json:"compra_totvmn" sql:"decimal(20,2)"`
	Compra_Oper_A       float32   `json:"compra_oa" sql:"decimal(20,2)"`
	Compra_Oper_I       float32   `json:"compra_oi" sql:"decimal(20,2)"`
	Compra_Oper_E       float32   `json:"compra_oe" sql:"decimal(20,2)"`
	Compra_Cant_Pago    int       `json:"compra_cantpago"`
	IDEstado            int       `json:"idestado"`
	IDProveedor         int       `json:"idproveedor"`
	ID_Tipodocumento    int       `json:"id_tipodocumento"`
	ID_Periodo          int       `json:"id_periodo"`
	ID_Moneda           int       `json:"id_moneda"`
	ID_Empresa          int       `json:"id_empresa"`
	ID_Local            int       `json:"id_local"`
	ID_Centro_Costo     int       `json:"id_centrocosto"`
	ID_Persona          int       `json:"id_persona"`
	Compra_Cant_Asiento int       `json:"compra_cantasiento"`
	Cod_Usua_Reg        int       `json:"cod_us_reg"`
	Cod_Usua_Act        int       `json:"cod_us_act"`
}

func CreateTodo(context *gin.Context) {
	var data todoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}
	todo.IDCompra = data.IDCompra
	todo.Compracod = data.Compracod
	todo.Codigo = data.Codigo
	todo.Compra_fecha_emi = data.Compra_fecha_emi
	todo.Compra_fecha_venc = data.Compra_fecha_venc
	todo.Compra_serie = data.Compra_serie
	todo.Compra_numer = data.Compra_numer
	todo.Compra_Descrip = data.Compra_Descrip
	todo.Compra_Obser = data.Compra_Obser
	todo.Compra_total_item = data.Compra_total_item
	todo.Compra_tipo_cambio = data.Compra_tipo_cambio
	todo.Compra_vbb = data.Compra_vbb
	todo.Compra_Vvitem = data.Compra_Vvitem
	todo.Compra_Dscto_global = data.Compra_Dscto_global
	todo.Compra_Isc = data.Compra_Isc
	todo.Compra_Vv = data.Compra_Vv
	todo.Compra_Pv = data.Compra_Pv
	todo.Compra_Cargo = data.Compra_Cargo
	todo.Compra_Tot_v = data.Compra_Tot_v
	todo.Compra_Tot_v_mn = data.Compra_Tot_v_mn
	todo.Compra_Oper_A = data.Compra_Oper_A
	todo.Compra_Oper_I = data.Compra_Oper_I
	todo.Compra_Oper_E = data.Compra_Oper_E
	todo.Compra_Cant_Pago = data.Compra_Cant_Pago
	todo.IDEstado = data.IDEstado
	todo.IDProveedor = data.IDProveedor
	todo.ID_Tipodocumento = data.ID_Tipodocumento
	todo.ID_Periodo = data.ID_Periodo
	todo.ID_Moneda = data.ID_Moneda
	todo.ID_Empresa = data.ID_Empresa
	todo.ID_Local = data.ID_Local
	todo.ID_Centro_Costo = data.ID_Centro_Costo
	todo.ID_Persona = data.ID_Persona
	todo.Compra_Cant_Asiento = data.Compra_Cant_Asiento
	todo.Cod_Usua_Reg = data.Cod_Usua_Reg
	todo.Cod_Usua_Act = data.Cod_Usua_Act

	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.IDCompra = todo.IDCompra
	response.Compracod = todo.Compracod
	response.Codigo = todo.Codigo
	response.Compra_fecha_emi = todo.Compra_fecha_emi
	response.Compra_fecha_venc = todo.Compra_fecha_venc
	response.Compra_serie = todo.Compra_serie
	response.Compra_numer = todo.Compra_numer
	response.Compra_Descrip = todo.Compra_Descrip
	response.Compra_Obser = todo.Compra_Obser
	response.Compra_total_item = todo.Compra_total_item
	response.Compra_tipo_cambio = todo.Compra_tipo_cambio
	response.Compra_vbb = todo.Compra_vbb
	response.Compra_Vvitem = todo.Compra_Vvitem
	response.Compra_Dscto_global = todo.Compra_Dscto_global
	response.Compra_Isc = todo.Compra_Isc
	response.Compra_Vv = todo.Compra_Vv
	response.Compra_Pv = todo.Compra_Pv
	response.Compra_Cargo = todo.Compra_Cargo
	response.Compra_Tot_v = todo.Compra_Tot_v
	response.Compra_Tot_v_mn = todo.Compra_Tot_v_mn
	response.Compra_Oper_A = todo.Compra_Oper_A
	response.Compra_Oper_I = todo.Compra_Oper_I
	response.Compra_Oper_E = todo.Compra_Oper_E
	response.Compra_Cant_Pago = todo.Compra_Cant_Pago
	response.IDEstado = todo.IDEstado
	response.IDProveedor = todo.IDProveedor
	response.ID_Tipodocumento = todo.ID_Tipodocumento
	response.ID_Periodo = todo.ID_Periodo
	response.ID_Moneda = todo.ID_Moneda
	response.ID_Empresa = todo.ID_Empresa
	response.ID_Local = todo.ID_Local
	response.ID_Centro_Costo = todo.ID_Centro_Costo
	response.ID_Persona = todo.ID_Persona
	response.Compra_Cant_Asiento = todo.Compra_Cant_Asiento
	response.Cod_Usua_Reg = todo.Cod_Usua_Reg
	response.Cod_Usua_Act = todo.Cod_Usua_Act

	context.JSON(http.StatusCreated, response)
}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// var response []todoResponse
	// for _, todo := range todos {
	// 	response = append(response, todoResponse{
	// 		ID:          todo.ID,
	// 		name:        todo.Name,
	// 		description: todo.Description,
	// 	})
	// }

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})

}

func UpdateTodo(context *gin.Context) {
	var data todoRequest
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	if err := context.BindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}

	todoById := db.Where("id = ?", idTodo).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Todo not found"})
		return
	}

	todo.IDCompra = data.IDCompra
	todo.Compracod = data.Compracod
	todo.Codigo = data.Codigo
	todo.Compra_fecha_emi = data.Compra_fecha_emi
	todo.Compra_fecha_venc = data.Compra_fecha_venc
	todo.Compra_serie = data.Compra_serie
	todo.Compra_numer = data.Compra_numer
	todo.Compra_Descrip = data.Compra_Descrip
	todo.Compra_Obser = data.Compra_Obser
	todo.Compra_total_item = data.Compra_total_item
	todo.Compra_tipo_cambio = data.Compra_tipo_cambio
	todo.Compra_vbb = data.Compra_vbb
	todo.Compra_Vvitem = data.Compra_Vvitem
	todo.Compra_Dscto_global = data.Compra_Dscto_global
	todo.Compra_Isc = data.Compra_Isc
	todo.Compra_Vv = data.Compra_Vv
	todo.Compra_Pv = data.Compra_Pv
	todo.Compra_Cargo = data.Compra_Cargo
	todo.Compra_Tot_v = data.Compra_Tot_v
	todo.Compra_Tot_v_mn = data.Compra_Tot_v_mn
	todo.Compra_Oper_A = data.Compra_Oper_A
	todo.Compra_Oper_I = data.Compra_Oper_I
	todo.Compra_Oper_E = data.Compra_Oper_E
	todo.Compra_Cant_Pago = data.Compra_Cant_Pago
	todo.IDEstado = data.IDEstado
	todo.IDProveedor = data.IDProveedor
	todo.ID_Tipodocumento = data.ID_Tipodocumento
	todo.ID_Periodo = data.ID_Periodo
	todo.ID_Moneda = data.ID_Moneda
	todo.ID_Empresa = data.ID_Empresa
	todo.ID_Local = data.ID_Local
	todo.ID_Centro_Costo = data.ID_Centro_Costo
	todo.ID_Persona = data.ID_Persona
	todo.Compra_Cant_Asiento = data.Compra_Cant_Asiento
	todo.Cod_Usua_Reg = data.Cod_Usua_Reg
	todo.Cod_Usua_Act = data.Cod_Usua_Act

	result := db.Save(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response todoResponse
	response.IDCompra = todo.IDCompra
	response.Compracod = todo.Compracod
	response.Codigo = todo.Codigo
	response.Compra_fecha_emi = todo.Compra_fecha_emi
	response.Compra_fecha_venc = todo.Compra_fecha_venc
	response.Compra_serie = todo.Compra_serie
	response.Compra_numer = todo.Compra_numer
	response.Compra_Descrip = todo.Compra_Descrip
	response.Compra_Obser = todo.Compra_Obser
	response.Compra_total_item = todo.Compra_total_item
	response.Compra_tipo_cambio = todo.Compra_tipo_cambio
	response.Compra_vbb = todo.Compra_vbb
	response.Compra_Vvitem = todo.Compra_Vvitem
	response.Compra_Dscto_global = todo.Compra_Dscto_global
	response.Compra_Isc = todo.Compra_Isc
	response.Compra_Vv = todo.Compra_Vv
	response.Compra_Pv = todo.Compra_Pv
	response.Compra_Cargo = todo.Compra_Cargo
	response.Compra_Tot_v = todo.Compra_Tot_v
	response.Compra_Tot_v_mn = todo.Compra_Tot_v_mn
	response.Compra_Oper_A = todo.Compra_Oper_A
	response.Compra_Oper_I = todo.Compra_Oper_I
	response.Compra_Oper_E = todo.Compra_Oper_E
	response.Compra_Cant_Pago = todo.Compra_Cant_Pago
	response.IDEstado = todo.IDEstado
	response.IDProveedor = todo.IDProveedor
	response.ID_Tipodocumento = todo.ID_Tipodocumento
	response.ID_Periodo = todo.ID_Periodo
	response.ID_Moneda = todo.ID_Moneda
	response.ID_Empresa = todo.ID_Empresa
	response.ID_Local = todo.ID_Local
	response.ID_Centro_Costo = todo.ID_Centro_Costo
	response.ID_Persona = todo.ID_Persona
	response.Compra_Cant_Asiento = todo.Compra_Cant_Asiento
	response.Cod_Usua_Reg = todo.Cod_Usua_Reg
	response.Cod_Usua_Act = todo.Cod_Usua_Act

	context.JSON(http.StatusCreated, response)
}

func DeleteTodo(context *gin.Context) {
	todo := models.Todo{}
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	delete := db.Where("id = ?", idTodo).Unscoped().Delete(&todo)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idTodo,
	})

}