package models

// Defines todo table for database communications
type Todo struct {
	IDCompra            uint    //`json:"idcompra" gorm:"primaryKey"`
	Compracod           int     //`json:"compracod"`
	Codigo              string  //`json:"codigo" sql:"type:varchar(45)"`
	Compra_fecha_emi    string  //`json:"comprafechaemi" sql:"type:date"`
	Compra_serie        string  //`json:"compraserie" sql:"type:varchar(45)"`
	Compra_numer        string  //`json:"compranumer" sql:"type:varchar(45)"`
	Compra_Descuent     string  //`json:"compradescuent" sql:"type:text"`
	Compra_total_item   int     //`json:"compratotalitem"`
	Compra_tipo_cambio  float32 //`json:"compratipocambio" sql:"type:decimal(20,3)"`
	Compra_vvb          float32 //`json:"compra_vvb" sql:"type:decimal(20,2)"`
	Compra_Oper_g       float32 //`json:"compra_og" sql:"type:decimal(20,2)"`
	Compra_Isc          float32 //`json:"compra_isc" sql:"type:decimal(20,2)"`
	Compra_Vv           float32 //`json:"compra_vv" sql:"type:decimal(20,2)"`
	Compra_Igv          float32 //`json:"compra_igv" sql:"type:decimal(20,2)"`
	Compra_Pv           float32 //`json:"compra_pv" sql:"type:decimal(20,2)"`
	Compra_Cargo        float32 //`json:"compra_cargo" sql:"type:decimal(20,2)"`
	Compra_Tot_v        float32 //`json:"compra_totv" sql:"type:decimal(20,2)"`
	Compra_Oper_a       float32 //`json:"compra_oa" sql:"type:decimal(20,2)"`
	Compra_Oper_i       float32 //`json:"compra_oi" sql:"type:decimal(20,2)"`
	Compra_Oper_e       float32 //`json:"compra_oe" sql:"type:decimal(20,2)"`
	IDEstado            int     //`json:"idestado"`
	IDProveedor         int     //`json:"idproveedor"`
	Txt_Tipodocumento   string  //`json:"txt_tipodocumento" sql:"type:varchar(15)"`
	Txt_Tipocomprobante string  //`json:"txt_tipocomprobante" sql:"type:varchar(12)"`
	ID_Periodo          int     //`json:"id_periodo"`
	Txt_Moneda          string  //`json:"txt_moneda" sql:"type:varchar(10)"`
	ID_Empresa          int     //`json:"id_empresa"`
	ID_Centro_Costo     int     //`json:"id_centrocosto"`
	Compra_Cant_Asiento int     //`json:"compra_cantasiento"`
	Igv                 int     //`json:"igv" sql:"type:decimal(4,2)"`
	Prov_Ruc            string  //`json:"prov_ruc" sql:"type:varchar(11)"`
	Prov_Nombre         string  //`json:"prov_nombre" sql:"type:varchar(150)"`
	Prov_Direccion      string  //`json:"prov_direccion" sql:"type:varchar(15)"`
	Prov_Contacto       string  //`json:"prov_contacto" sql:"type:varchar(50)"`
	Prov_Telefono       string  //`json:"prov_telefono" sql:"type:varchar(20)"`
	Prov_Correo         string  //`json:"prov_correo" sql:"type:varchar(50)"`
	Fech_Reg            string  //`json:"fech_reg" sql:"type:datetime"`
	Fech_Act            string  //`json:"fech_act" sql:"type:datetime"`
	Cod_Usua_Reg        int     //`json:"cod_us_reg"`
	Cod_Usua_Act        int     //`json:"cod_us_act"`
}

func (Todo) TableName() string {
	return "compras_rf"
}
