package tools

import (
	"fmt"
	"time"
)

// FechaMysql genera una cadena de fecha y hora formateada para MySQL
//
// Retorna:
//   - string: Fecha y hora en formato "YYYY-MM-DDThh:mm:ss"
//
// La función:
// 1. Obtiene la fecha y hora actual
// 2. Formatea cada componente con dos dígitos (mes, día, hora, minuto, segundo)
// 3. Retorna la cadena formateada compatible con MySQL
//
// Nota: El formato de dos dígitos es requerido por MySQL para evitar errores
// de formato en las operaciones de base de datos
func FechaMysql() string {
	t := time.Now()
	// ESTO ES IMPORTANTE SE FORZA PARA QUE SI O SI LO FORMATEE EN 2 DIGITOS POR FORMATO CORRECTO EN MYSQL PARA EVITAR ERROR
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// esto es muy importante para el motor de mysql en golang para que se formatee, muy importante
