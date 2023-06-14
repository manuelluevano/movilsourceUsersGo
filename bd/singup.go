package bd

import (
	"fmt"

	"example.com/m/v2/models"
	"example.com/m/v2/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza el registro")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	//RESPUESTA SQL
	_, err = Db.Exec(sentencia)

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	fmt.Println("SignUp > Ejecucion Exitosa")

	return nil
}
