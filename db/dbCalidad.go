package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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
	query := "select DNIUsuario, NombreUsuario from Usuario"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	user := new(Usuario)
	users := []Usuario{}
	for rows.Next() {
		err1 := rows.Scan(&user.DNIUsuario, &user.NombreUsuario)
		if err1 != nil {
			panic(err1.Error())
		} else {
			users = append(users, *user)
		}
	}
	return users
}

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

func ConsultaContactos() []Contacto {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select Usuario_DNIUsuario, NumeroTelef, NombreContacto from Contacto"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	user := new(Contacto)
	users := []Contacto{}
	for rows.Next() {
		err1 := rows.Scan(&user.Usuario_DNIUsuario, &user.NumeroTelef, &user.NombreContacto)
		if err1 != nil {
			panic(err1.Error())
		} else {
			users = append(users, *user)
		}
	}
	return users
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