package models

// SecretRDSJson representa la estructura de datos para las credenciales
// y configuración de conexión a la base de datos RDS almacenadas en AWS Secrets Manager
//
// Campos:
//   - Username: Nombre de usuario para la conexión a la base de datos
//   - Password: Contraseña para la conexión a la base de datos
//   - Engine: Tipo de motor de base de datos (ej: mysql, postgres)
//   - Host: Dirección del servidor de base de datos
//   - Port: Puerto de conexión a la base de datos
//   - DbClusterIdentifier: Identificador del cluster de RDS
type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"Port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}

// Signup representa la estructura de datos para el registro de un nuevo usuario
// Esta estructura se utiliza para procesar los eventos de registro de Cognito
//
// Campos:
//   - UserEmail: Dirección de correo electrónico del usuario
//   - UserUUID: Identificador único del usuario generado por Cognito
type Signup struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
