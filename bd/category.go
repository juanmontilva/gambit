package bd

import (
	"database/sql"
	"fmt"
	"strings"

	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/juanmontilva/gambit/models"
	"github.com/juanmontilva/gambit/tools"
)

func InsertCategory(c models.Category) (int64, error) {

	fmt.Println("COMIENZA REGISTRO DE INSERTCATEGORY")

	err := DbConnect()
	if err != nil {
		return 0, err
	}

	defer Db.Close()

	sentencia := "INSERT INTO category (Categ_Name, Categ_Path) VALUES('" + c.CategName + "','" + c.CategPath + "')"

	var result sql.Result

	result, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()

	if err2 != nil {
		return 0, err2
	}

	fmt.Println("INSERT CATEGORY EJECUCION EXITOSA")
	return LastInsertId, nil

}

func UpdateCategory(c models.Category) error {

	fmt.Println("COMIENZA REGISTRO DE UPDATECATEGORY")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "UPDATE category SET "

	if len(c.CategName) > 0 {
		sentencia += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}

	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(sentencia, "SET ") {
			sentencia += ", "
		}
		sentencia += "Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}

	sentencia += " WHERE Categ_Id = " + strconv.Itoa(c.CategId)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("UPDATE CATEGORY > EJECUCION EXITOSA")

	return nil

}
