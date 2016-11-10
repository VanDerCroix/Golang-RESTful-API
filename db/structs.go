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


// ---- ARQUITECTURABD ---
/*type Facultad struct {
	Id              string
	Nombre          string
	Autoridad       string
	URLFoto		    string
}

type FacultadDetalles struct {
	Id string
	Nombre string
	Autoridad string
	IdUbicacion string
	Ubicacion Ubicacion
	Escuelas []Escuela
}

type Administrador struct {
	Id         string
	Nombre     string
	Correo     string
	Contrasena string
}

type Autoridad struct {
	Id     string
	Nombre string
}

type Escuela struct {
	Id              string
	Nombre          string
	IdAutoridad     string
	IdFacultad      string
	IdAdministrador string
}

type Noticia struct {
	Id              string
	Titulo          string
	Descripcion     string
	Fecha      		string
	IdAdministrador string
}

type Ubicacion struct {
	Id              string
	Latitud         string
	Longitud  	   	string
	URLFoto    		string
}

type AreasUniversidad struct{
	Id 				string
	Nombre 			string
	IdUniversidad 	string
	IdUbicacion 	string
	IdAdministrador string
	URL string
}*/
