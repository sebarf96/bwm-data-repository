package user_service_test

import (
	"bwm-api-repository/src/database"
	"bwm-api-repository/src/models"
	"bwm-api-repository/src/repositories/user_repository"
	"bwm-api-repository/src/services/user_service"
	"github.com/go-kit/kit/log"
	"os"
	"testing"
	"time"
)

var (
	service user_service.UserService
)

func init() {
	logger := log.NewLogfmtLogger(os.Stderr)
	repository := user_repository.NewUserRepository(database.GetCollection("users"))
	service = user_service.NewUserService(logger, repository)
}

func TestServiceCreate(t *testing.T) {

	usr1 := models.User{
		Name:     "Seba",
		Surname:  "Remaggi",
		Email:    "seb@gmail.com",
		CreateAt: time.Now(),
	}

	res,err := service.Create(usr1)
	if err != nil {
		t.Error("Test fail")
	} else {
		t.Log("Test assert", res)
	}
}

func TestServiceRead(t *testing.T) {

	users, err := service.Read()
	if err != nil {
		t.Error("Ha fallado el test")
		t.Fail()
	}
	if len(users) == 0 {

		t.Error("No hay registros")
		t.Fail()
	} else {

		t.Log("Exito: ", users)
	}

}

func TestServiceUpdate(t *testing.T) {

	user := models.User{
		Name:  "Giorgio",
		Email: "gio@gmail.com",
	}

	err := service.Update(user, "60dafa96acff82038f718701")
	if err != nil {
		t.Error("Error al actualizar los datos")
		t.Fail()
	} else {
		t.Log("La prueba finalizó con exito")
	}
}

func TestDelete(t *testing.T) {

	err := service.Delete("60dbe5ea3b99308bda5153c1")
	if err != nil {
		t.Error("Error al actualizar los datos")
		t.Fail()
	} else {
		t.Log("La prueba finalizó con exito")
	}

}
