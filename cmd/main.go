package main

import (
db	"bwm-api-repository/src/database"
	ep "bwm-api-repository/src/endpoint"
	"bwm-api-repository/src/handler"
	"bwm-api-repository/src/models"
	"bwm-api-repository/src/repositories/user_repository"
	"bwm-api-repository/src/services/user_service"

	"github.com/go-kit/kit/log"
)

func main() {

	repository := user_repository.NewUserRepository(db.GetCollection("users"))
	service := user_service.NewUserService(log.NewNopLogger(), repository)
	handler.NewHandler(getEndpoints(service),":8080")

}
func getEndpoints(svc user_service.UserService) models.Eps {
	epCreate := ep.MakeUserCreateEndpoint(svc)
	epRead := ep.MakeUserReadEndpoint(svc)
	epUpdate := ep.MakeUserUpdateEndpoint(svc)
	epDelete := ep.MakeUserDeleteEndpoint(svc)

	return models.Eps{
		CreateEP: epCreate,
		ReadEP:   epRead,
		UpdateEP: epUpdate,
		DeleteEP: epDelete,
	}
}