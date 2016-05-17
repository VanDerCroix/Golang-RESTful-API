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
		&rest.Route{"GET", "/escuelas", Escuelas_handler},
		&rest.Route{"GET", "/administradores", Administradores_handler},
		&rest.Route{"GET", "/ubicacion/:id", Ubicacion_handler},
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
func Escuelas_handler(w rest.ResponseWriter, r *rest.Request) {
	escuelas := datos.ConsultaEscuelas()
	w.WriteJson(escuelas)
}

func Administradores_handler(w rest.ResponseWriter, r *rest.Request) {
	admins := datos.ConsultaAdministradores()
	w.WriteJson(admins)
}

func Ubicacion_handler(w rest.ResponseWriter, r *rest.Request) {
	idubi := r.PathParam("id")

	ubicacion := datos.ConsultaUbicacion(idubi)
	w.WriteJson(ubicacion)
}
