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
	query := "select f.idFacultadxsede, f.nombreF, a.nombreAut, ub.urlfoto from facultadxsede f join ubicacion ub on f.idUbicacion = ub.idUbicacion join autoridad a on f.idAutoridad = a.idAutoridad"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fac := new(Facultad)
	facs := []Facultad{}
	for rows.Next() {
		err1 := rows.Scan(&fac.Id, &fac.Nombre, &fac.Autoridad, &fac.URLFoto)
		if err1 != nil {
			panic(err1.Error())
		} else {
			facs = append(facs, *fac)
		}
	}
	return facs
}

func ConsultaFacultad(id string) Facultad {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select f.idFacultadxsede, f.nombreF, a.nombreAut, ub.urlfoto from facultadxsede f join ubicacion ub on f.idUbicacion = ub.idUbicacion join autoridad a on f.idAutoridad = a.idAutoridad where f.idFacultadxsede=?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fac := new(Facultad)

	if rows.Next() {
		err1 := rows.Scan(&fac.Id, &fac.Nombre, &fac.Autoridad, &fac.URLFoto)
		if err1 != nil {
			panic(err1.Error())
		}
	}
	return *fac
}

func ConsultaFacultadDetalles(id string) FacultadDetalles {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select f.idFacultadxsede, f.nombreF, a.nombreAut, f.idUbicacion from facultadxsede f join autoridad a on f.idAutoridad = a.idAutoridad where f.idFacultadxsede=?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fac := new(FacultadDetalles)

	if rows.Next() {
		err1 := rows.Scan(&fac.Id, &fac.Nombre, &fac.Autoridad, &fac.IdUbicacion)
		if err1 != nil {
			panic(err1.Error())
		}
	}
	//get ubicacion strucut
	fac.Ubicacion = ConsultaUbicacion(fac.IdUbicacion)
	//get escuales array
	fac.Escuelas = ConsultaEscuelasxFacu(fac.Id)
	return *fac
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

func ConsultaEscuelasxFacu(id string) []Escuela{
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idEscuela, nombreE, idAutoridad, idFacultadxsede, idAdministrador FROM escuela WHERE idFacultadxsede=?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
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

func ConsultaNoticias() []Noticia {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idNoticia, encabezado, cuerpo, fecha, idAdministrador FROM noticias"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	not := new(Noticia)
	nots := []Noticia{}
	for rows.Next() {
		err1 := rows.Scan(&not.Id, &not.Titulo, &not.Descripcion, &not.Fecha, &not.IdAdministrador)
		if err1 != nil {
			panic(err1.Error())
		} else {
			//log.Println("emp: ", id, name, mail)
			nots = append(nots, *not)
		}
	}
	return nots
}

func ConsultaUbicaciones() []Ubicacion {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idUbicacion, latitud, longitud, urlfoto FROM ubicacion"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	ubi := new(Ubicacion)
	ubis := []Ubicacion{}
	for rows.Next() {
		err1 := rows.Scan(&ubi.Id, &ubi.Latitud, &ubi.Longitud, &ubi.URLFoto)
		if err1 != nil {
			panic(err1.Error())
		} else {
			//log.Println("emp: ", id, name, mail)
			ubis = append(ubis, *ubi)
		}
	}
	return ubis
}

func ConsultaUbicacion(id string) Ubicacion {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idUbicacion, latitud, longitud, urlfoto FROM ubicacion WHERE idUbicacion=?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	ubi := new(Ubicacion)
	//ubis := []Ubicacion{}
	if rows.Next() {
		err1 := rows.Scan(&ubi.Id, &ubi.Latitud, &ubi.Longitud, &ubi.URLFoto)
		if err1 != nil {
			panic(err1.Error())
		}
	}
	return *ubi
}

func ConsultaAreasUniversidad() []AreasUniversidad {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "SELECT idAreasuniversidad, nombreAU, idUniversidad, idUbicacion, idAdministrador, urlfoto FROM areasuniversidad"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	asu := new(AreasUniversidad)
	asus := []AreasUniversidad{}
	for rows.Next() {
		err1 := rows.Scan(&asu.Id, &asu.Nombre, &asu.IdUniversidad, &asu.IdUbicacion, &asu.IdAdministrador, &asu.URL)
		if err1 != nil {
			panic(err1.Error())
		} else {
			//log.Println("emp: ", id, name, mail)
			asus = append(asus, *asu)
		}
	}
	return asus
}
