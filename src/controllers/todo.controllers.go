package controllers

import (
	"fmt"
	"net/http"

	"github.com/JeanIzar/crud-go/src/config"
	"github.com/JeanIzar/crud-go/src/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	IDCompra            uint    `json:"idcompra" gorm:"primaryKey"`
	Compracod           int     `json:"compracod"`
	Codigo              string  `json:"codigo" sql:"type:varchar(45)"`
	Compra_fecha_emi    string  `json:"comprafechaemi" sql:"type:date"`
	Compra_serie        string  `json:"compraserie" sql:"type:varchar(45)"`
	Compra_numer        string  `json:"compranumer" sql:"type:varchar(45)"`
	Compra_Descuent     string  `json:"compradescuent" sql:"type:text"`
	Compra_total_item   int     `json:"compratotalitem"`
	Compra_tipo_cambio  float32 `json:"compratipocambio" sql:"type:decimal(20,3)"`
	Compra_vvb          float32 `json:"compra_vvb" sql:"type:decimal(20,2)"`
	Compra_Oper_g       float32 `json:"compra_og" sql:"type:decimal(20,2)"`
	Compra_Isc          float32 `json:"compra_isc" sql:"type:decimal(20,2)"`
	Compra_Vv           float32 `json:"compra_vv" sql:"type:decimal(20,2)"`
	Compra_Igv          float32 `json:"compra_igv" sql:"type:decimal(20,2)"`
	Compra_Pv           float32 `json:"compra_pv" sql:"type:decimal(20,2)"`
	Compra_Cargo        float32 `json:"compra_cargo" sql:"type:decimal(20,2)"`
	Compra_Tot_v        float32 `json:"compra_totv" sql:"type:decimal(20,2)"`
	Compra_Oper_a       float32 `json:"compra_oa" sql:"type:decimal(20,2)"`
	Compra_Oper_i       float32 `json:"compra_oi" sql:"type:decimal(20,2)"`
	Compra_Oper_e       float32 `json:"compra_oe" sql:"type:decimal(20,2)"`
	IDEstado            int     `json:"idestado"`
	IDProveedor         int     `json:"idproveedor"`
	Txt_Tipodocumento   string  `json:"txt_tipodocumento" sql:"type:varchar(15)"`
	Txt_Tipocomprobante string  `json:"txt_tipocomprobante" sql:"type:varchar(12)"`
	ID_Periodo          int     `json:"id_periodo"`
	Txt_Moneda          string  `json:"txt_moneda" sql:"type:varchar(10)"`
	ID_Empresa          int     `json:"id_empresa"`
	ID_Centro_Costo     int     `json:"id_centrocosto"`
	Compra_Cant_Asiento int     `json:"compra_cantasiento"`
	Igv                 int     `json:"igv" sql:"type:decimal(4,2)"`
	Prov_Ruc            string  `json:"prov_ruc" sql:"type:varchar(11)"`
	Prov_Nombre         string  `json:"prov_nombre" sql:"type:varchar(150)"`
	Prov_Direccion      string  `json:"prov_direccion" sql:"type:varchar(15)"`
	Prov_Contacto       string  `json:"prov_contacto" sql:"type:varchar(50)"`
	Prov_Telefono       string  `json:"prov_telefono" sql:"type:varchar(20)"`
	Prov_Correo         string  `json:"prov_correo" sql:"type:varchar(50)"`
	Fech_Reg            string  `json:"fech_reg" sql:"type:datetime"`
	Fech_Act            string  `json:"fech_act" sql:"type:datetime"`
	Cod_Usua_Reg        int     `json:"cod_us_reg"`
	Cod_Usua_Act        int     `json:"cod_us_act"`
}

type todoResponse struct {
	todoRequest
	IDCompra            uint    `json:"idcompra" gorm:"primaryKey"`
	Compracod           int     `json:"compracod"`
	Codigo              string  `json:"codigo" sql:"type:varchar(45)"`
	Compra_fecha_emi    string  `json:"comprafechaemi" sql:"type:date"`
	Compra_serie        string  `json:"compraserie" sql:"type:varchar(45)"`
	Compra_numer        string  `json:"compranumer" sql:"type:varchar(45)"`
	Compra_Descuent     string  `json:"compradescuent" sql:"type:text"`
	Compra_total_item   int     `json:"compratotalitem"`
	Compra_tipo_cambio  float32 `json:"compratipocambio" sql:"type:decimal(20,3)"`
	Compra_vvb          float32 `json:"compra_vvb" sql:"type:decimal(20,2)"`
	Compra_Oper_g       float32 `json:"compra_og" sql:"type:decimal(20,2)"`
	Compra_Isc          float32 `json:"compra_isc" sql:"type:decimal(20,2)"`
	Compra_Vv           float32 `json:"compra_vv" sql:"type:decimal(20,2)"`
	Compra_Igv          float32 `json:"compra_igv" sql:"type:decimal(20,2)"`
	Compra_Pv           float32 `json:"compra_pv" sql:"type:decimal(20,2)"`
	Compra_Cargo        float32 `json:"compra_cargo" sql:"type:decimal(20,2)"`
	Compra_Tot_v        float32 `json:"compra_totv" sql:"type:decimal(20,2)"`
	Compra_Oper_a       float32 `json:"compra_oa" sql:"type:decimal(20,2)"`
	Compra_Oper_i       float32 `json:"compra_oi" sql:"type:decimal(20,2)"`
	Compra_Oper_e       float32 `json:"compra_oe" sql:"type:decimal(20,2)"`
	IDEstado            int     `json:"idestado"`
	IDProveedor         int     `json:"idproveedor"`
	Txt_Tipodocumento   string  `json:"txt_tipodocumento" sql:"type:varchar(15)"`
	Txt_Tipocomprobante string  `json:"txt_tipocomprobante" sql:"type:varchar(12)"`
	ID_Periodo          int     `json:"id_periodo"`
	Txt_Moneda          string  `json:"txt_moneda" sql:"type:varchar(10)"`
	ID_Empresa          int     `json:"id_empresa"`
	ID_Centro_Costo     int     `json:"id_centrocosto"`
	Compra_Cant_Asiento int     `json:"compra_cantasiento"`
	Igv                 int     `json:"igv" sql:"type:decimal(4,2)"`
	Prov_Ruc            string  `json:"prov_ruc" sql:"type:varchar(11)"`
	Prov_Nombre         string  `json:"prov_nombre" sql:"type:varchar(150)"`
	Prov_Direccion      string  `json:"prov_direccion" sql:"type:varchar(15)"`
	Prov_Contacto       string  `json:"prov_contacto" sql:"type:varchar(50)"`
	Prov_Telefono       string  `json:"prov_telefono" sql:"type:varchar(20)"`
	Prov_Correo         string  `json:"prov_correo" sql:"type:varchar(50)"`
	Fech_Reg            string  `json:"fech_reg" sql:"type:datetime"`
	Fech_Act            string  `json:"fech_act" sql:"type:datetime"`
	Cod_Usua_Reg        int     `json:"cod_us_reg"`
	Cod_Usua_Act        int     `json:"cod_us_act"`
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
	todo.Compra_serie = data.Compra_serie
	todo.Compra_numer = data.Compra_numer
	todo.Compra_Descuent = data.Compra_Descuent
	todo.Compra_total_item = data.Compra_total_item
	todo.Compra_tipo_cambio = data.Compra_tipo_cambio
	todo.Compra_vvb = data.Compra_vvb
	todo.Compra_Oper_g = data.Compra_Oper_g
	todo.Compra_Isc = data.Compra_Isc
	todo.Compra_Vv = data.Compra_Vv
	todo.Compra_Igv = data.Compra_Igv
	todo.Compra_Pv = data.Compra_Pv
	todo.Compra_Cargo = data.Compra_Cargo
	todo.Compra_Tot_v = data.Compra_Tot_v
	todo.Compra_Oper_a = data.Compra_Oper_a
	todo.Compra_Oper_i = data.Compra_Oper_i
	todo.Compra_Oper_e = data.Compra_Oper_e
	todo.IDEstado = data.IDEstado
	todo.IDProveedor = data.IDProveedor
	todo.Txt_Tipodocumento = data.Txt_Tipodocumento
	todo.Txt_Tipocomprobante = data.Txt_Tipocomprobante
	todo.ID_Periodo = data.ID_Periodo
	todo.Txt_Moneda = data.Txt_Moneda
	todo.ID_Empresa = data.ID_Empresa
	todo.ID_Centro_Costo = data.ID_Centro_Costo
	todo.Compra_Cant_Asiento = data.Compra_Cant_Asiento
	todo.Igv = data.Igv
	todo.Prov_Ruc = data.Prov_Ruc
	todo.Prov_Nombre = data.Prov_Nombre
	todo.Prov_Direccion = data.Prov_Direccion
	todo.Prov_Contacto = data.Prov_Contacto
	todo.Prov_Telefono = data.Prov_Telefono
	todo.Prov_Telefono = data.Prov_Telefono
	todo.Fech_Reg = data.Fech_Reg
	todo.Fech_Act = data.Fech_Act
	todo.Cod_Usua_Reg = data.Cod_Usua_Reg
	todo.Cod_Usua_Act = data.Cod_Usua_Act

	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Hubo un error al registrar"})
		return
	}

	var response todoResponse
	response.IDCompra = todo.IDCompra
	response.Compracod = todo.Compracod
	response.Codigo = todo.Codigo
	response.Compra_fecha_emi = todo.Compra_fecha_emi
	response.Compra_serie = todo.Compra_serie
	response.Compra_numer = todo.Compra_numer
	response.Compra_Descuent = todo.Compra_Descuent
	response.Compra_total_item = todo.Compra_total_item
	response.Compra_tipo_cambio = todo.Compra_tipo_cambio
	response.Compra_vvb = todo.Compra_vvb
	response.Compra_Oper_g = todo.Compra_Oper_g
	response.Compra_Isc = todo.Compra_Isc
	response.Compra_Vv = todo.Compra_Vv
	response.Compra_Igv = todo.Compra_Igv
	response.Compra_Pv = todo.Compra_Pv
	response.Compra_Cargo = todo.Compra_Cargo
	response.Compra_Tot_v = todo.Compra_Tot_v
	response.Compra_Oper_a = todo.Compra_Oper_a
	response.Compra_Oper_i = todo.Compra_Oper_i
	response.Compra_Oper_e = todo.Compra_Oper_e
	response.IDEstado = todo.IDEstado
	response.IDProveedor = todo.IDProveedor
	response.Txt_Tipodocumento = todo.Txt_Tipodocumento
	response.Txt_Tipocomprobante = todo.Txt_Tipocomprobante
	response.ID_Periodo = todo.ID_Periodo
	response.Txt_Moneda = todo.Txt_Moneda
	response.ID_Empresa = todo.ID_Empresa
	response.ID_Centro_Costo = todo.ID_Centro_Costo
	response.Compra_Cant_Asiento = todo.Compra_Cant_Asiento
	response.Igv = todo.Igv
	response.Prov_Ruc = todo.Prov_Ruc
	response.Prov_Nombre = todo.Prov_Nombre
	response.Prov_Direccion = todo.Prov_Direccion
	response.Prov_Contacto = todo.Prov_Contacto
	response.Prov_Telefono = todo.Prov_Telefono
	response.Prov_Telefono = todo.Prov_Telefono
	response.Fech_Reg = todo.Fech_Reg
	response.Fech_Act = todo.Fech_Act
	response.Cod_Usua_Reg = todo.Cod_Usua_Reg
	response.Cod_Usua_Act = todo.Cod_Usua_Act

	context.JSON(http.StatusCreated, response)
}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	if err := db.Find(&todos).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos"})
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

func GetIdTodo(context *gin.Context) {
	todo := models.Todo{}
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	compra := db.Where("id_compra = ?", idTodo).Unscoped().First(&todo)
	fmt.Println(compra)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todo,
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

	//Opcion 1
	todo := models.Todo{}
	todoById := db.Where("id_compra=?", idTodo).First(&todo)
	if todoById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Compra no ubicada"})
		return
	}

	todo.IDCompra = data.IDCompra
	todo.Compracod = data.Compracod
	todo.Codigo = data.Codigo
	todo.Compra_fecha_emi = data.Compra_fecha_emi
	todo.Compra_serie = data.Compra_serie
	todo.Compra_numer = data.Compra_numer
	todo.Compra_Descuent = data.Compra_Descuent
	todo.Compra_total_item = data.Compra_total_item
	todo.Compra_tipo_cambio = data.Compra_tipo_cambio
	todo.Compra_vvb = data.Compra_vvb
	todo.Compra_Oper_g = data.Compra_Oper_g
	todo.Compra_Isc = data.Compra_Isc
	todo.Compra_Vv = data.Compra_Vv
	todo.Compra_Igv = data.Compra_Igv
	todo.Compra_Pv = data.Compra_Pv
	todo.Compra_Cargo = data.Compra_Cargo
	todo.Compra_Tot_v = data.Compra_Tot_v
	todo.Compra_Oper_a = data.Compra_Oper_a
	todo.Compra_Oper_i = data.Compra_Oper_i
	todo.Compra_Oper_e = data.Compra_Oper_e
	todo.IDEstado = data.IDEstado
	todo.IDProveedor = data.IDProveedor
	todo.Txt_Tipodocumento = data.Txt_Tipodocumento
	todo.Txt_Tipocomprobante = data.Txt_Tipocomprobante
	todo.ID_Periodo = data.ID_Periodo
	todo.Txt_Moneda = data.Txt_Moneda
	todo.ID_Empresa = data.ID_Empresa
	todo.ID_Centro_Costo = data.ID_Centro_Costo
	todo.Compra_Cant_Asiento = data.Compra_Cant_Asiento
	todo.Igv = data.Igv
	todo.Prov_Ruc = data.Prov_Ruc
	todo.Prov_Nombre = data.Prov_Nombre
	todo.Prov_Direccion = data.Prov_Direccion
	todo.Prov_Contacto = data.Prov_Contacto
	todo.Prov_Telefono = data.Prov_Telefono
	todo.Prov_Telefono = data.Prov_Telefono
	todo.Fech_Reg = data.Fech_Reg
	todo.Fech_Act = data.Fech_Act
	todo.Cod_Usua_Reg = data.Cod_Usua_Reg
	todo.Cod_Usua_Act = data.Cod_Usua_Act

	//Opcion1
	result := db.Where("id_compra = ?", idTodo).Updates(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Ocurrio un error al actualizar"})
		return
	}

	var response todoResponse
	response.IDCompra = todo.IDCompra
	response.Compracod = todo.Compracod
	response.Codigo = todo.Codigo
	response.Compra_fecha_emi = todo.Compra_fecha_emi
	response.Compra_serie = todo.Compra_serie
	response.Compra_numer = todo.Compra_numer
	response.Compra_Descuent = todo.Compra_Descuent
	response.Compra_total_item = todo.Compra_total_item
	response.Compra_tipo_cambio = todo.Compra_tipo_cambio
	response.Compra_vvb = todo.Compra_vvb
	response.Compra_Oper_g = todo.Compra_Oper_g
	response.Compra_Isc = todo.Compra_Isc
	response.Compra_Vv = todo.Compra_Vv
	response.Compra_Igv = todo.Compra_Igv
	response.Compra_Pv = todo.Compra_Pv
	response.Compra_Cargo = todo.Compra_Cargo
	response.Compra_Tot_v = todo.Compra_Tot_v
	response.Compra_Oper_a = todo.Compra_Oper_a
	response.Compra_Oper_i = todo.Compra_Oper_i
	response.Compra_Oper_e = todo.Compra_Oper_e
	response.IDEstado = todo.IDEstado
	response.IDProveedor = todo.IDProveedor
	response.Txt_Tipodocumento = todo.Txt_Tipodocumento
	response.Txt_Tipocomprobante = todo.Txt_Tipocomprobante
	response.ID_Periodo = todo.ID_Periodo
	response.Txt_Moneda = todo.Txt_Moneda
	response.ID_Empresa = todo.ID_Empresa
	response.ID_Centro_Costo = todo.ID_Centro_Costo
	response.Compra_Cant_Asiento = todo.Compra_Cant_Asiento
	response.Igv = todo.Igv
	response.Prov_Ruc = todo.Prov_Ruc
	response.Prov_Nombre = todo.Prov_Nombre
	response.Prov_Direccion = todo.Prov_Direccion
	response.Prov_Contacto = todo.Prov_Contacto
	response.Prov_Telefono = todo.Prov_Telefono
	response.Prov_Telefono = todo.Prov_Telefono
	response.Fech_Reg = todo.Fech_Reg
	response.Fech_Act = todo.Fech_Act
	response.Cod_Usua_Reg = todo.Cod_Usua_Reg
	response.Cod_Usua_Act = todo.Cod_Usua_Act

	context.JSON(http.StatusCreated, response)
}

func DeleteTodo(context *gin.Context) {
	todo := models.Todo{}
	reqParamId := context.Param("idTodo")
	idTodo := cast.ToUint(reqParamId)

	delete := db.Where("id_compra = ?", idTodo).Unscoped().Delete(&todo)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    idTodo,
	})

}
