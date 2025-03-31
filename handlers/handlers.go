package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juanmontilva/gambit/auth"
	"github.com/juanmontilva/gambit/routes"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	//ESTO LO VOY A VER EN POSTMAN O INSOMNIA
	fmt.Println("VOY A PROCESAR" + path + " > " + method)

	//se configura la ruta de apigateway

	//tengo dos datos, el string y el numerico que viene siendo idn, para eso se usa el strconv para transformar a int
	id := request.PathParameters["id"]
	//atoi necesita el error pero de verdad no me interesa si no tendria que poner si error es distinto de nil
	// OJO ESTO SE UTILIZA PARA EVITAR ERROR EN LOS SWITCH Y CASE POR SER STRING TRANSFORMO A INT
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validoAuthorization(path, method, headers)

	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcesoUser(body, path, method, user, id, request)

	case "prod":
		return ProcesoProducts(body, path, method, user, idn, request)

	case "stoc":
		return ProcesoStock(body, path, method, user, idn, request)

	case "addr":
		return ProcesoAdress(body, path, method, user, idn, request)

	case "cate":
		return ProcesoCategory(body, path, method, user, idn, request)

	case "orde":
		return ProcesoOrder(body, path, method, user, idn, request)

	}

	return 400, "Method Invalid"

}

// OJO TIENE QUE RETORNAR LO MISMO QUE AUTH EN EL PAQUETE AUTH BOOL, INT Y STRING
func validoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "TOKEN REQUERIDO"
	}

	todoOk, err, msg := auth.ValidoToken(token)
	if !todoOk {
		if err != nil {
			fmt.Println("ERROR EN EL TOKEN" + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("ERROR EN EL TOKEN " + msg)
			return false, 401, msg
		}
	}

	fmt.Println("token ok")
	return true, 200, msg

}

func ProcesoUser(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoProducts(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoCategory(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	switch method {
	case "POST":
		return routes.InsertCategory(body, user)
	case "PUT":
		return routes.UpdateCategory(body, user, id)
	}

	return 400, "Method Invalid"
}

func ProcesoStock(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoAdress(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method Invalid"
}

func ProcesoOrder(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {

	return 400, "Method Invalid"
}
