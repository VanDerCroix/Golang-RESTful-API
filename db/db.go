package mysqldb

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "log"
    "fmt"
)
func main() {
    //Query2()
}

func Query(){
    db, err := sql.Open("mysql", CxStr)
    if err != nil{
        log.Printf(err.Error())
    }

    defer db.Close()

    // Execute the query
    query := "SELECT id, name, address FROM Employees"
    fmt.Println(query)
    rows, err := db.Query(query)//SELECT * FROM table
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

func QueryEmployees()([]Employee) {
    db, err := sql.Open("mysql", CxStr)
    if err != nil{
        log.Printf(err.Error())
    }

    defer db.Close()

    // Execute the query
    query := "SELECT id, name, address FROM Employees"
    fmt.Println(query)
    rows, err := db.Query(query)//SELECT * FROM table
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    emp := new(Employee)
    emps := []Employee{}
    for rows.Next() {
        err1 := rows.Scan(&emp.Id, &emp.Name, &emp.Mail)
        if err1 != nil{
            panic(err1.Error())
        }else{
            //log.Println("emp: ", id, name, mail)
            emps = append(emps, *emp)
        }
    }
    return emps
}

func QueryEmployee(id int)(Employee) {
    db, err := sql.Open("mysql", CxStr)
    if err != nil{
        log.Printf(err.Error())
    }

    defer db.Close()

    // Execute the query
    query := "SELECT id, name, address FROM Employees where id=?"
    fmt.Println(query)
    rows, err := db.Query(query,id)//SELECT * FROM table
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    emp := new(Employee)
    
    if rows.Next() {
        err1 := rows.Scan(&emp.Id, &emp.Name, &emp.Mail)
        if err1 != nil{
            panic(err1.Error())
        }
    }
    return *emp
}