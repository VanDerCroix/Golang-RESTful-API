package mysqldb


// ---- CALIDAD SOS ---
type Usuario struct{
	DNIUsuario int
	NombreUsuario string
	Peso float32
	Talla float32
	Sexo string
	TipoSangre string
	MensajePrederteminado string
}

type Contacto struct{
	IdContacto int
	Usuario_DNIUsuario int
	NumeroTelef string
	NombreContacto string
}

type CentrosAtencion struct{
	Latitud float32
	Longitud float32
	NombreCentAten string
	Direccion string
	Telefono int
	URLFoto string
	Distrito string
}

type Alergia struct{
	IdAlergias int
	Usuario_DNIUsuario int
	NombreAlergia string
	Medicacion string	
}

type Categoria struct{
	IdCategoria int
	Nombre string
	URLFoto string
}

type Subcategoria struct{
	IdSubcategoria int
	Categoria_Id int
	Nombre string
	URLFoto string
}

type Recomendacion struct{
	IdRecomendacion int
	Parte int
	Descripcion string
	Sub_Categoria_Id int
}