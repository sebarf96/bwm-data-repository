package handler

import (
	"bwm-api-repository/src/models"
	"context"
	"encoding/json"
	"net/http"
)
/*
func NewHandler(
	svc user_service.UserService,
	createEP endpoint.Endpoint,
	readEP endpoint.Endpoint,
	updateEP endpoint.Endpoint,
	deleteEP endpoint.Endpoint) (*http.Server{


	createHandler :=httptransport.NewServer(createEP ,decodeCreateRequest, encodeResponse,)
	http.Handle("/create",createHandler)
	return createHandler

}
*/
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request models.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}