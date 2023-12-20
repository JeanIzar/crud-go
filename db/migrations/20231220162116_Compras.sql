-- migrate:up
create table compras_rf (
    idcompra integer,
    compracod integer,
    codigo varchar(45),
    comprafecemis date,
    comprafecvenc date,
    compraserie varchar(45),
    compranumer varchar(45),
    compradesc text,
    compraobs text,
    compratotalitem integer,
    compra_tipocambio decimal(20,3),
    compra_vvb decimal(20,2),
    compra_vvitem decimal(20,2),
    compra_dsctoglobal decimal(20,2),
    compra_isc decimal(20,2),
    compra_vv decimal(20,2),
    compra_igv decimal(20,2),
    compra_pv decimal(20,2),
    compra_cargo decimal(20,2),
    compra_totv decimal(20,2),
    compra_totvmn decimal(20,2),
    compra_oa decimal(20,2),
    compra_oi decimal(20,2),
    compra_oe decimal(20,2),
    compra_cantpago integer,
    idestado integer,
    idproveedor integer,
    id_tipodocumento integer,
    idperiodo integer,
    idmoneda integer,
    idempresa integer,
    idlocal integer,
    idcentrocosto integer,
    idpersona integer,
    compra_cant_asiento integer,
    fechreg datetime,
    fechact datetime,
    cod_us_reg integer,
    cod_us_act integer,
    fech_created_at timestamp,
    fech_updated_at timestamp,
    fech_deleted_at timestamp
    )

-- migrate:down

