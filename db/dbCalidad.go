package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"strconv"
)

func main() {

}

func ConsultaUsuarios() []Usuario {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select DNIUsuario, NombreUsuario, Peso, Talla, Sexo, TipoSangre, MensajePredeterminado from Usuario"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	user := new(Usuario)
	users := []Usuario{}
	for rows.Next() {
		err1 := rows.Scan(&user.DNIUsuario, &user.NombreUsuario, &user.Peso, &user.Talla, &user.Sexo, &user.TipoSangre, &user.MensajePrederteminado)
		if err1 != nil {
			panic(err1.Error())
		} else {
			users = append(users, *user)
		}
	}
	return users
}

func ConsultaContactos(dni int) []Contacto {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select NombreContacto, NumeroTelef from Contacto where Usuario_DNIUsuario = ?"
	fmt.Println(query)
	rows, err := db.Query(query, dni) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	contacto := new(Contacto)
	contactos := []Contacto{}
	for rows.Next() {
		err1 := rows.Scan(&contacto.IdContacto, &contacto.Usuario_DNIUsuario, &contacto.NombreContacto, &contacto.NumeroTelef)
		if err1 != nil {
			panic(err1.Error())
		} else {
			contactos = append(contactos, *contacto)
		}
	}
	return contactos
}

func ConsultaAlergias(dni int) []Alergia {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select IdAlergias, NombreAlergia, Medicacion from Alergia where Usuario_DNIUsuario = ?"
	fmt.Println(query)
	rows, err := db.Query(query, dni) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	alergia := new(Alergia)
	alergias := []Alergia{}
	for rows.Next() {
		err1 := rows.Scan(&alergia.IdAlergias, &alergia.NombreAlergia, &alergia.Medicacion)
		if err1 != nil {
			panic(err1.Error())
		} else {
			alergias = append(alergias, *alergia)
		}
	}
	return alergias
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
	
func ConsultaCentrosAtencion() []CentrosAtencion {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select c.Latitud, c.Longitud, c.NombreCentAten, c.Direccion, c.Telefono, c.URLFoto, d.NombreDistrito from Centros_de_Atencion c join Distrito d on c.Distrito_Id = d.Id order by d.Id"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	centro := new(CentrosAtencion)
	centros := []CentrosAtencion{}
	for rows.Next() {
		err1 := rows.Scan(&centro.Latitud, &centro.Longitud, &centro.NombreCentAten, &centro.Direccion, &centro.Telefono, &centro.URLFoto, &centro.Distrito)
		if err1 != nil {
			panic(err1.Error())
		} else {
			centros = append(centros, *centro)
		}
	}
	return centros
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func ConsultaCategorias() []Categoria {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select IdCategoria, Nombre, URLFoto from Categoria"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	categoria := new(Categoria)
	categorias := []Categoria{}
	for rows.Next() {
		err1 := rows.Scan(&categoria.IdCategoria, &categoria.Nombre, &categoria.URLFoto)
		if err1 != nil {
			panic(err1.Error())
			} else {
			categorias = append(categorias, *categoria)
		}
	}
	return categorias
}

func ConsultaSubcategorias(id int) []Subcategoria {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select IdSubcategoria, Categoria_Id, Nombre, URLFoto from Sub_Categoria where Categoria_Id = ?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	subcategoria := new(Subcategoria)
	subcategorias := []Subcategoria{}
	for rows.Next() {
		err1 := rows.Scan(&subcategoria.IdSubcategoria, &subcategoria.Categoria_Id, &subcategoria.Nombre, &subcategoria.URLFoto)
		if err1 != nil {
			panic(err1.Error())
		} else {
			subcategorias = append(subcategorias, *subcategoria)
		}
	}
	return subcategorias
}

func ConsultaRecomendacion(id int) []Recomendacio {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select IdRecomendacion, Parte, Descripcion, Sub_Categoria_Id from Recomendacion where Sub_Categoria_Id = ?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	recomendacio := new(Recomendacio)
	recomendacios := []Recomendacio{}
	for rows.Next() {
		err1 := rows.Scan(&recomendacio.IdRecomendacion, &recomendacio.Parte, &recomendacio.Descripcion, &subcategoria.Sub_Categoria_Id)
		if err1 != nil {
			panic(err1.Error())
		} else {
			recomendacios = append(recomendacios, *recomendacio)
		}
	}
	return recomendacios
}

///////////////////////////////////////////////////////////////////////////////////////////////////////

func InsertarUsuario(usr Usuario) {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	res := mustExec(db, "INSERT INTO Usuario(DNIUsuario, NombreUsuario, Peso, Talla, Sexo, TipoSangre, MensajePredeterminado)  VALUES (?, ?, ?, ?, ?, ?, ?)", usr.DNIUsuario, usr.NombreUsuario, usr.Peso, usr.Talla, usr.Sexo, usr.TipoSangre, usr.MensajePrederteminado)
	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("res.RowsAffected() returned error: %s", err.Error())
	}
	if count != 1 {
		log.Printf("expected 1 affected row, got %d", count)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("res.LastInsertId() returned error: %s", err.Error())
	}
	if id != 0 {
		log.Printf("expected InsertId 0, got %d", id)
	}
}

func InsertarContacto(usr Contacto) {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	res := mustExec(db, "INSERT INTO Contacto(Usuario_DNIUsuario, NombreContacto, NumeroTelef) VALUES (?, ?, ?)", usr.Usuario_DNIUsuario, usr.NombreContacto, usr.NumeroTelef)
	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("res.RowsAffected() returned error: %s", err.Error())
	}
	if count != 1 {
		log.Printf("expected 1 affected row, got %d", count)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("res.LastInsertId() returned error: %s", err.Error())
	}
	if id != 0 {
		log.Printf("expected InsertId 0, got %d", id)
	}
}

func InsertarAlergia(usr Alergia) {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	res := mustExec(db, "INSERT INTO Alergias(Usuario_DNIUsuario, NombreAlergia, Medicacion) VALUES (?, ?, ?)", usr.Usuario_DNIUsuario, usr.NombreAlergia, usr.Medicacion)
	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("res.RowsAffected() returned error: %s", err.Error())
	}
	if count != 1 {
		log.Printf("expected 1 affected row, got %d", count)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("res.LastInsertId() returned error: %s", err.Error())
	}
	if id != 0 {
		log.Printf("expected InsertId 0, got %d", id)
	}
}

func mustExec(db *sql.DB, query string, args ...interface{}) (res sql.Result) {
	res, err := db.Exec(query, args...)
	if err != nil {
		fail("exec", query, err)
	}
	return res
}

func fail(method, query string, err error) {
	if len(query) > 300 {
		query = "[query too large to print]"
	}
	// dbt.Fatalf("error on %s %s: %s", method, query, err.Error())
	log.Printf("error on %s %s: %s", method, query, err.Error())
}
