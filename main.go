package main

import (
	"context"
	"os"

	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/juanmontilva/gambit/awsgo"

	"github.com/juanmontilva/gambit/bd"
	"github.com/juanmontilva/gambit/handlers"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializoAws()

	if !ValidoParametros() {
		panic("Error en LOS PARAMETROS DEBE ENVIAR 'SecretName', 'UrlPrefix'")
	}

	//esto me permite manejar una API DE MANERA 100% OPTIMA

	var res *events.APIGatewayProxyResponse

	//OJO SE PUEDE AGREGAR EN EL path el prefix directamente, se podria optimizar memoria pero no lo hago para entender mejor
	prefix := os.Getenv("UrlPrefix")

	path := strings.Replace(request.RawPath, prefix, "", -1)

	method := request.RequestContext.HTTP.Method

	body := request.Body

	header := request.Headers

	//CON ESTA PARTE LA API SE MANEJARA SI O SI DE UNA MANERA MUY OPTIMA res, prefix, path, method, body, header, ojo leer documentacion de lambda porque posiblemente no funcione gorilla mux

	bd.ReadSecret()

	status, message := handlers.Manejadores(path, method, body, header, request)

	headesResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headesResp,
	}

	return res, nil

}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	//necesito investigar el porque no hace falta utilizar, supuestamente viene integrado con cognito, esto puede ser codigo no util sin necesidad, se eliminara del panic !ValidoParametros

	// _, traeParametro = os.LookupEnv("UserPoolId")
	// if !traeParametro {
	// 	return traeParametro
	// }

	// _, traeParametro = os.LookupEnv("Region")
	// if !traeParametro {
	// 	return traeParametro
	// }

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro

}
