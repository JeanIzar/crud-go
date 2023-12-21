package models

import (
	"gorm.io/gorm"
)

// Defines todo table for database communications
type Todo struct {
	gorm.Model
	IDCompra            uint    //`json:"idcompra" gorm:"primaryKey"`
	Compracod           int     //`json:"compracod"`
	Codigo              string  //`json:"codigo" sql:"varchar(45)"`
	Compra_fecha_emi    string  //`json:"comprafechaemi" sql:"type:date"`
	Compra_fecha_venc   string  //`json:"comprafechavenc" sql:"type:date"`
	Compra_serie        string  //`json:"compraserie" sql:"varchar(45)"`
	Compra_numer        string  //`json:"compranumer" sql:"varchar(45)"`
	Compra_Descrip      string  //`json:"compradescrip" sql:"type:text"`
	Compra_Obser        string  //`json:"compraobser" sql:"type:text"`
	Compra_total_item   int     //`json:"compratotalitem"`
	Compra_tipo_cambio  float32 //`json:"compratipocambio" sql:"type:decimal(20,3)"`
	Compra_vbb          float32 //`json:"compra_vbb" sql:"type:decimal(20,2)"`
	Compra_dscto_item   float32 //`json:"compra_dscto_item" sql:"type:decimal(20,2)"`
	Compra_Oper_G       float32 //`json:"compra_og" sql:"type:decimal(20,2)"`
	Compra_Vvitem       float32 //`json:"compra_vvitem" sql:"type:decimal(20,2)"`
	Compra_Dscto_global float32 //`json:"compra_dsctoglobal" sql:"decimal(20,2)"`
	Compra_Isc          float32 //`json:"compra_isc" sql:"decimal(20,2)"`
	Compra_Vv           float32 //`json:"compra_vv" sql:"decimal(20,2)"`
	Compra_Igv          float32 //`json:"compra_igv" sql:"decimal(20,2)"`
	Compra_Pv           float32 //`json:"compra_pv" sql:"decimal(20,2)"`
	Compra_Cargo        float32 //`json:"compra_cargo" sql:"decimal(20,2)"`
	Compra_Tot_v        float32 //`json:"compra_totv" sql:"decimal(20,2)"`
	Compra_Tot_v_mn     float32 //`json:"compra_totvmn" sql:"decimal(20,2)"`
	Compra_Oper_A       float32 //`json:"compra_oa" sql:"decimal(20,2)"`
	Compra_Oper_I       float32 //`json:"compra_oi" sql:"decimal(20,2)"`
	Compra_Oper_E       float32 //`json:"compra_oe" sql:"decimal(20,2)"`
	Compra_Cant_Pago    int     //`json:"compra_cantpago"`
	IDEstado            int     //`json:"idestado"`
	IDProveedor         int     //`json:"idproveedor"`
	ID_Tipodocumento    int     //`json:"id_tipodocumento"`
	ID_Tipocomprobante  int     //`json:"id_tipocomprobante"`
	ID_Periodo          int     //`json:"id_periodo"`
	ID_Moneda           int     //`json:"id_moneda"`
	ID_Empresa          int     //`json:"id_empresa"`
	ID_Local            int     //`json:"id_local"`
	ID_Centro_Costo     int     //`json:"id_centrocosto"`
	ID_Persona          int     //`json:"id_persona"`
	Compra_Cant_Asiento int     //`json:"compra_cantasiento"`
	Cod_Usua_Reg        int     //`json:"cod_us_reg"`
	Cod_Usua_Act        int     //`json:"cod_us_act"`
}
