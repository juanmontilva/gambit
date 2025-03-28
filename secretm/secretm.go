package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/juanmontilva/gambit/awsgo"
	"github.com/juanmontilva/gambit/models"
)

// GetSecret recupera un secreto almacenado en AWS Secrets Manager
//
// Parámetros:
//   - nombreSecret: Nombre del secreto a recuperar de AWS Secrets Manager
//
// Retorna:
//   - models.SecretRDSJson: Estructura con las credenciales y configuración de la base de datos
//   - error: nil si la operación fue exitosa, error en caso contrario
//
// La función:
// 1. Crea un cliente para AWS Secrets Manager
// 2. Recupera el valor del secreto especificado
// 3. Deserializa el JSON del secreto en la estructura SecretRDSJson
func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println("> Pido Secreto" + nombreSecret)

	// Crea un cliente para AWS Secrets Manager usando la configuración global
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	// Recupera el valor del secreto especificado
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	// Deserializa el JSON del secreto en la estructura SecretRDSJson
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(">Lectura secret ok" + nombreSecret)

	return datosSecret, nil
}
