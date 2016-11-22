package main
	
import (
	datos "./db"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"strconv"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		&rest.Route{"GET", "/ejemplo1", Ejemplo1_handler},
		&rest.Route{"GET", "/ejemplo2", Ejemplo2_handler},
		// ---- CALIDAD SOS ---
		&rest.Route{"GET", "/usuario", Usuario_handler},
		&rest.Route{"GET", "/usuario/:dni/contacto", Contacto_handler},
		&rest.Route{"GET", "/usuario/:dni/alergia", Alergia_handler},
		&rest.Route{"GET", "/centrosatencion", Centros_Atencion_handler},
		&rest.Route{"GET", "/categorias", Categorias_handler},
		&rest.Route{"GET", "/subcategorias/:categoriaid", Sub_Categorias_handler},
		&rest.Route{"GET", "/recomendacion/:subcategoriaid", Recomendacion_handler},
		&rest.Route{"POST", "/postusuario", PostUsuario_handler},
		&rest.Route{"POST", "/postcontacto", PostContacto_handler},
		&rest.Route{"POST", "/postalergia", PostAlergia_handler},
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9988", api.MakeHandler()))
}

// ---- EJEMPLOS ---
func Ejemplo1_handler(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{
		"nombre": "barry allen",
		"correo": "aelosmit@gmail.com",
	})
}
func Ejemplo2_handler(w rest.ResponseWriter, r *rest.Request) {
/*	idubi, err := strconv.Atoi(r.PathParam("id"))

	if err != nil {
		w.WriteJson(map[string]string{"Error": err.Error()})
	} else {
		ubicacion := datos.ConsultaUbicacion(idubi)
		w.WriteJson(ubicacion)
	}*/
	// datos.Query()
	w.WriteJson(map[string]string{
		"nombre": "2",
		"correo": "asdit@gmail.com",
	})
}

// ---- CALIDAD SOS ---
/////////////////////////////////////////////////////////////////////////
func Usuario_handler(w rest.ResponseWriter, r *rest.Request) {
	usuarios := datos.ConsultaUsuarios()
	w.WriteJson(usuarios)
}

func Contacto_handler(w rest.ResponseWriter, r *rest.Request) {
	userdni := r.PathParam("dni")
	dni, _ := strconv.Atoi(userdni)
	contactos := datos.ConsultaContactos(dni)
	w.WriteJson(contactos)
}

func Alergia_handler(w rest.ResponseWriter, r *rest.Request) {
	userdni := r.PathParam("dni")
	dni, _ := strconv.Atoi(userdni)
	alergias := datos.ConsultaAlergias(dni)
	w.WriteJson(alergias)
}

//////////////////////////////////////////////////////////////////////////
func Centros_Atencion_handler(w rest.ResponseWriter, r *rest.Request){
centros_atencion := datos.ConsultaCentrosAtencion()
	w.WriteJson(centros_atencion)	
}

//////////////////////////////////////////////////////////////////////////

func Categorias_handler(w rest.ResponseWriter, r *rest.Request){
categorias := datos.ConsultaCategorias()
	w.WriteJson(centros_atencion)	
}

func Sub_Categorias_handler(w rest.ResponseWriter, r *rest.Request) {
	categoriaid := r.PathParam("categoriaid")
	id, _ := strconv.Atoi(categoriaid)
	subcategorias := datos.ConsultaSubcategorias(id)
	w.WriteJson(subcategorias)
}

func Recomendacion_handler(w rest.ResponseWriter, r *rest.Request) {
	subcategoriaid := r.PathParam("subcategoriaid")
	id, _ := strconv.Atoi(subcategoriaid)
	recomendacion := datos.ConsultaRecomendacion(id)
	w.WriteJson(recomendacion)
}

/////////////////////////////////////////////////////////////////////////

func PostUsuario_handler(w rest.ResponseWriter, r *rest.Request) {
	usuario := new(datos.Usuario)
	err := r.DecodeJsonPayload(&usuario)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datos.InsertarUsuario(*usuario)
	w.WriteJson(&usuario)
}

func PostContacto_handler(w rest.ResponseWriter, r *rest.Request) {
	contacto := new(datos.Contacto)
	err := r.DecodeJsonPayload(&contacto)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datos.InsertarContacto(*contacto)
	w.WriteJson(&contacto)
}

func PostAlergia_handler(w rest.ResponseWriter, r *rest.Request) {
	alergia := new(datos.Alergia)
	err := r.DecodeJsonPayload(&alergia)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datos.InsertarAlergia(*alergia)
	w.WriteJson(&alergia)
}