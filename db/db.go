package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//Query2()
}

func Query() {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT * FROM administrador"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//return rows

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}

// ---- ARQUITECTURABD ---
func ConsultaFacultades() []Facultad {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idFacultadxsede, nombreF, idAutoridad, idUniversidad, idUbicacion, idAdministrador FROM facultadxsede"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fac := new(Facultad)
	facs := []Facultad{}
	for rows.Next() {
		err1 := rows.Scan(&fac.Id, &fac.Nombre, &fac.IdAutoridad, &fac.IdUniversidad, &fac.IdUbicacion, &fac.IdAdministrador)
		if err1 != nil {
			panic(err1.Error())
		} else {
			//log.Println("emp: ", id, name, mail)
			facs = append(facs, *fac)
		}
	}
	return facs
}


func ConsultaEscuelas() []Escuela {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idEscuela, nombreE, idAutoridad, idFacultadxsede, idAdministrador FROM escuela"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	esc := new(Escuela)
	escs := []Escuela{}
	for rows.Next() {
		err1 := rows.Scan(&esc.Id, &esc.Nombre, &esc.IdAutoridad, &esc.IdFacultad, &esc.IdAdministrador)
		if err1 != nil {
			panic(err1.Error())
		} else {
			//log.Println("emp: ", id, name, mail)
			escs = append(escs, *esc)
		}
	}
	return escs
}

func ConsultaAdministradores() []Administrador {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idadministrador, nombreadm, correo, contrasenha FROM administrador"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	adm := new(Administrador)
	adms := []Administrador{}
	for rows.Next() {
		err1 := rows.Scan(&adm.Id, &adm.Nombre, &adm.Correo, &adm.Contrasena)
		if err1 != nil {
			panic(err1.Error())
		} else {
			//log.Println("emp: ", id, name, mail)
			adms = append(adms, *adm)
		}
	}
	return adms
}

func ConsultaUbicacion(id string) Ubicacion {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idubicacion, latitud, longitud, foto  FROM ubicacion WHERE idubicacion=?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	ubi := new(Ubicacion)

	if rows.Next() {
		err1 := rows.Scan(&ubi.Id, &ubi.Latitud, &ubi.Longitud, &ubi.Foto)
		if err1 != nil {
			panic(err1.Error())
		}
	}
	return *ubi
}
