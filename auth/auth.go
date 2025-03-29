package auth

import (

	//todos los jwt vienen en base 64, no vale la pena traer el paquete completo encoding porque es muy pesado, es mejor optimizar lo necesario
	"encoding/base64"
	"encoding/json"

	//necesito fmt para mostrar al cloudwatch
	"fmt"
	"strings"
	"time"
)

// ESTO ES LO QUE SE VA A DESESTRUCTURAR
type TokenJson struct {
	Sub       string
	Event_Id  string
	Token_Use string
	Scope     string
	Auth_Time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidoToken(token string) (bool, error, string) {
	//se tiene que dividir el token en las 3 partes, el separador es el punto .
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("TOKEN NO VALIDO, NO VINO CON LAS 3 PARTES DEL TOKEN")
		return false, nil, "Token no valido"
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("NO SE PUEDE DECODIFICAR LA PARTE DEL TOKEN", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJson
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("NO SE PUEDE DECODIFICAR EN LA ESTRUCTURA JSON", err.Error())
		return false, err, err.Error()
	}

	//siempre es bueno mantener una seguridad propia y no depender de la seguridad de api gateway, siempre es bueno mantener comentarios en el cloudwatch
	ahora := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	//funcion importante del time, permite evitar comparacion de mayor o menor en ahora y tm, las pone igual

	if tm.Before(ahora) {
		fmt.Println("FECHA EXPIRACION TOKEN = " + tm.String())
		fmt.Println("TOKEN EXPIRADO")
		return false, err, "TOKEN EXPIRADO!"
	}

	//es importante retornar en una funcion string a tkj.username por el siguiente motivo, amazon es muy diferente y si se envia directo como respuesta puede generar problema por ese valor
	return true, nil, string(tkj.Username)

}
