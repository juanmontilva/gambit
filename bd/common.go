package bd

import (
	// database/sql es un estandar en golang pero la otra importacion que son los drivers no, hay que hacer un go get
	"database/sql"
	"fmt"
	"os"

	// este es el driver que se importa, es recomendado poner _ porque no me importa, solamente que funcione el driver de sql, no me la se toda, tengo qeu averiguar, esto esta en la documentacion, muy interesante, si no se pone el _ da error
	_ "github.com/go-sql-driver/mysql"
	"github.com/juanmontilva/gambit/models"
	"github.com/juanmontilva/gambit/secretm"
)

// SecretModel almacena las credenciales y configuración de la base de datos
// obtenidas desde AWS Secrets Manager
var SecretModel models.SecretRDSJson
var err error

// Db es la conexión global a la base de datos
// Se define como variable pública (mayúscula) para ser accesible desde otros paquetes
// Se usa un puntero para mejorar el rendimiento en operaciones de base de datos
var Db *sql.DB

// ReadSecret lee las credenciales de la base de datos desde AWS Secrets Manager
//
// Retorna:
//   - error: nil si la lectura fue exitosa, error en caso contrario
func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

// DbConnect establece la conexión con la base de datos MySQL
//
// Retorna:
//   - error: nil si la conexión fue exitosa, error en caso contrario
//
// La función:
// 1. Abre la conexión usando las credenciales almacenadas
// 2. Verifica la conexión mediante un ping
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Verifica que la conexión esté activa
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("conexion exitosa de la base de datos")
	return nil
}

// ConnStr genera la cadena de conexión para MySQL
//
// Parámetros:
//   - claves: Estructura que contiene las credenciales y configuración de la base de datos
//
// Retorna:
//   - string: Cadena de conexión formateada para MySQL
func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string

	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "gambit"

	// Genera la cadena de conexión en formato DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
