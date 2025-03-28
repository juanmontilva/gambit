package handlers

import(
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)


func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string){


	//ESTO LO VOY A VER EN POSTMAN O INSOMNIA
	fmt.Println("VOY A PROCESAR" + path +" > " +method)

	//se configura la ruta de apigateway

	//tengo dos datos, el string y el numerico que viene siendo idn, para eso se usa el strconv para transformar a int
	id := request.PathParameters["id"]
	//atoi necesita el error pero de verdad no me interesa si no tendria que poner si error es distinto de nil
	idn, _ := strconv.Atoi(id)






	return 400, "Method Invalid"













}