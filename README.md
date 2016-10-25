# Api Rest + Mysql Cx

Set the variable CxStr on the package mysqldb

```go
var CxStr string = "username:password@tcp(uri.com:3306)/dbname"
```

POST USAGE
insert a new user

```
curl -X POST -H "Content-Type: application/json" -d "{\"DNIUsuario\": 61845192,\"NombreUsuario\": \"Chino Registro de Prueba\",\"Peso\": 51.15,\"Talla\": 180.15,\"Sexo\": \"M\",\"TipoSangre\": \"O-\",\"MensajePrederteminado\": \"habla gud\"}" http://54.190.14.3:9988/postusuario
```

USER LIST
http://54.190.14.3:9988/usuario