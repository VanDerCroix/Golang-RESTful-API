package main

import (
	datos "./db"
	"log"
	"net/http"
	"github.com/ant0ine/go-json-rest/rest"
	"strconv"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		&rest.Route{"GET","/registros", Registros_handler},
		&rest.Route{"GET","/empleados", Empleados_handler},
		&rest.Route{"GET","/empleados/:id", Empleado_handler},
		)

	if err != nil{
		log.Fatal(err)
	}
	//db.Query()
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9988", api.MakeHandler()))
}

func Registros_handler(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{
			"nombre":  "barry allen",
			"correo":  "aelosmit@gmail.com",
		})
	/**/
	/*
	rpta := datos.Query()
	log.Println(rpta)
	w.WriteJson(rpta)
	*/
}

//52.35.136.144
//curl -i http://52.35.136.144:9988/registros
func Empleados_handler(w rest.ResponseWriter, r *rest.Request) {
	empleados := datos.QueryEmployees()
	w.WriteJson(empleados)
}
func Empleado_handler(w rest.ResponseWriter, r *rest.Request) {
	idemp, err := strconv.Atoi(r.PathParam("id"))

	if err != nil {
		w.WriteJson(map[string]string{"Error": err.Error()})
	} else {
		empleados := datos.QueryEmployee(idemp)
		w.WriteJson(empleados)		
	}

}