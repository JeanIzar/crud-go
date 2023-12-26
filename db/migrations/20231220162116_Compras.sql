-- migrate:up
create table compras_rf (
    created_at timestamp not null,
    updated_at timestamp,
    deleted_at timestamp,
    idcompra integer not null primary key,
    compracod integer not null,
    codigo varchar(45) not null,
    compra_fecha_emi date not null,
    compra_serie varchar(45) not null,
    compra_numer varchar(45) not null,
    compra_descuent text not null,
    compra_total_item integer,
    compra_tipo_cambio decimal(20,3),
    compra_vvb decimal(20,2),
    compra_oper_g decimal(20,2),
    compra_isc decimal(20,2),
    compra_vv decimal(20,2),
    compra_igv decimal(20,2),
    compra_pv decimal(20,2),
    compra_cargo decimal(20,2),
    compra_tot_v decimal(20,2),
    compra_oper_a decimal(20,2),
    compra_oper_i decimal(20,2),
    compra_oper_e decimal(20,2),
    id_estado integer not null,
    id_proveedor integer not null,
    txt_tipodocumento varchar(15) not null,
    txt_tipocomprobante varchar(12) not null,
    id_periodo integer not null,
    txt_moneda varchar(10) not null,
    id_empresa integer not null,
    id_centro_costo integer,
    compra_cant_asiento integer,
    igv decimal(4,2) not null,
    prov_ruc varchar(11),
    prov_nombre varchar(150),
    prov_direccion varchar(15),
    prov_contacto varchar(50),
    prov_telefono varchar(20),
    prov_correo varchar(50),
    cod_usua_reg integer not null,
    cod_usua_act integer not null
    )

-- migrate:down

