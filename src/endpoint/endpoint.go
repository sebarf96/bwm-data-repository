package endpoint

import (
	"bwm-api-repository/src/models"
	"bwm-api-repository/src/services/user_service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeUserCreateEndpoint(svc user_service.UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.User)
		err := svc.Create(req)
		if err != nil {
			return err, nil
		}
		return err, nil
	}
}

func MakeUserUpdateEndpoint(svc user_service.UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.User)
		err := svc.Update(req, "")
		if err != nil {
			return err, nil
		}
		return err, nil
	}
}
func MakeUserDeleteEndpoint(svc user_service.UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		err := svc.Delete(req)
		if err != nil {
			return err, nil
		}
		return err, nil
	}
}
func MakeUserReadEndpoint(svc user_service.UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		res, err := svc.Read()
		if err != nil {
			return err, nil
		}
		return res, nil
	}
}
