package mysqldb

// ---- ARQUITECTURABD ---
type Facultad struct {
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
}
