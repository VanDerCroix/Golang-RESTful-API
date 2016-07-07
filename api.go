package main

import (
	datos "./db"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	//"strconv"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		/*&rest.Route{"GET", "/registros", Registros_handler},
		&rest.Route{"GET", "/empleados", Empleados_handler},
		&rest.Route{"GET", "/empleados/:id", Empleado_handler},*/

		// ---- ARQUITECTURABD ---
		&rest.Route{"GET", "/facultades", Facultades_handler},
		&rest.Route{"GET", "/facultades/:id", Facultad_handler},
		&rest.Route{"GET", "/escuelas", Escuelas_handler},
		&rest.Route{"GET", "/escuelasxfacu/:id", EscuelasxFacu_handler},
		&rest.Route{"GET", "/administradores", Administradores_handler},
		&rest.Route{"GET", "/noticias", Noticias_handler},
		&rest.Route{"GET", "/ubicacion", Ubicacion_handler},
		&rest.Route{"GET", "/areasuniversidad", AreasUniversidad_handler},
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9988", api.MakeHandler()))
}

/*func Registros_handler(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{
		"nombre": "barry allen",
		"correo": "aelosmit@gmail.com",
	})
}
func Ubicacion_handler(w rest.ResponseWriter, r *rest.Request) {
	idubi, err := strconv.Atoi(r.PathParam("id"))

	if err != nil {
		w.WriteJson(map[string]string{"Error": err.Error()})
	} else {
		ubicacion := datos.ConsultaUbicacion(idubi)
		w.WriteJson(ubicacion)
	}
}*/

// ---- ARQUITECTURABD ---
func Facultades_handler(w rest.ResponseWriter, r *rest.Request) {
	facultad := datos.ConsultaFacultades()
	w.WriteJson(facultad)
}

func Facultad_handler(w rest.ResponseWriter, r *rest.Request) {
	facuid := r.PathParam("id")

	facultad := datos.ConsultaFacultadDetalles(facuid)
	w.WriteJson(facultad)
}

func Escuelas_handler(w rest.ResponseWriter, r *rest.Request) {
	escuelas := datos.ConsultaEscuelas()
	w.WriteJson(escuelas)
}

func EscuelasxFacu_handler(w rest.ResponseWriter, r *rest.Request) {
	facuid := r.PathParam("id")

	escuelas := datos.ConsultaEscuelasxFacu(facuid)
	w.WriteJson(escuelas)
}

func Administradores_handler(w rest.ResponseWriter, r *rest.Request) {
	admins := datos.ConsultaAdministradores()
	w.WriteJson(admins)
}

func Noticias_handler(w rest.ResponseWriter, r *rest.Request) {
	news := datos.ConsultaNoticias()
	w.WriteJson(news)
}

func Ubicacion_handler(w rest.ResponseWriter, r *rest.Request) {
	ubis := datos.ConsultaUbicaciones()
	w.WriteJson(ubis)
}

func AreasUniversidad_handler(w rest.ResponseWriter, r *rest.Request) {
	asus := datos.ConsultaAreasUniversidad()
	w.WriteJson(asus)
}
