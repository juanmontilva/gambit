package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// Ctx es el contexto global para las operaciones de AWS
// Se utiliza para manejar timeouts y cancelaciones en las llamadas a la API
var Ctx context.Context

// Cfg almacena la configuración de AWS
// Contiene las credenciales y la región por defecto
var Cfg aws.Config
var err error

// InicializoAws configura el contexto y la configuración de AWS
// Esta función debe ser llamada al inicio de la aplicación para establecer
// la conexión con los servicios de AWS
//
// La función:
// 1. Crea un contexto sin límites de tiempo
// 2. Carga la configuración por defecto de AWS
// 3. Establece la región por defecto como us-east-1
//
// Si hay un error en la configuración, la función entrará en pánico
func InicializoAws() {
	// Crea un contexto sin límites de tiempo para las operaciones de AWS
	Ctx = context.TODO()

	// Carga la configuración por defecto de AWS con la región especificada
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))

	if err != nil {
		panic("ERROR CONEXION DE LAMBDA .aws/config" + err.Error())
	}
}
