package routes


import(
	"encoding/json"
	"strconv"


	//"github.com/aws/aws-lambda-go/events"
	"github.com/juanmontilva/gambit/bd"
	"github.com/juanmontilva/gambit/models"
)



func InsertCategory(body string, User string) (int, string) {


		var t models.Category

		err := json.Unmarshal([]byte(body), &t)
		if err != nil{
			return 400, "error en los datos recibidos"+ err.Error()
		}


		if len(t.CategName) == 0 {
			return 400, "debe especificar el nombre (title) de la categoria"
		}


		if len(t.CategPath) == 0 {
			return 400, "debe especificar el path (ruta) de la categoria"
		}

		//esto es una validacion de administrador
		isAdmin, msg := bd.UserIsAdmin(User)
		if !isAdmin{
			return 400, msg
		}

		// se tiene que crear un erro2 porque generaria problemas en el primero, es bueno separarlo
		result, err2 := bd.InsertCategory(t)
		if err2 != nil{
			return 400, "OCURRIO UN ERROR AL INTENTAR REALIZAAR UN REGISTRO DE CATEGORIA "+ t.CategName + " > " + err2.Error()
		}



		return 200, "{ CategID "+ strconv.Itoa(int(result)) + "}"
}