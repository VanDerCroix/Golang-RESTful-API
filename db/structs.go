package mysqldb

type Employee struct {
	Id   int
	Name string
	Mail string
}

// ---- ARQUITECTURABD ---

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

type Facultad struct {
	Id              string
	Nombre          string
	IdAutoridad     string
	IdUniversidad   string
	IdUbicacion     string
	IdAdministrador string
}

type Ubicacion struct {
	Id       string
	Latitud  string
	Longitud string
	Foto     []byte
}
