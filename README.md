# Api Rest + Mysql Cx

Set the variable CxStr on the package mysqldb

```go
var CxStr string = "username:password@tcp(uri.com:3306)/dbname"
```

POST USAGE
insert a new user

```
curl -X POST -H "Content-Type: application/json" -d "{\"DNIUsuario\": 49186284,\"NombreUsuario\": \"Chico Registro de Prueba\",\"Peso\": 50.15,\"Talla\": 170.15,\"Sexo\": \"M\",\"TipoSangre\": \"RH+\",\"MensajePrederteminado\": \"aiudaaaa\"}" http://127.0.0.1:9988/postusuario
```