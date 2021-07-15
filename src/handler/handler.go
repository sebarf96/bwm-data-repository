package handler

import (
	m "bwm-api-repository/src/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)


func NewHandler(ep m.Eps, port string) {

	createHandler := httptransport.NewServer(ep.CreateEP, decodeCreateRequest, encodeResponse)
	readHandler := httptransport.NewServer(ep.DeleteEP, decodeReadRequest, encodeResponse)
	updateHandler := httptransport.NewServer(ep.UpdateEP, decodeUpdateRequest, encodeResponse)
	deleteHandler := httptransport.NewServer(ep.DeleteEP, decodeDeleteRequest, encodeResponse)

	http.Handle("/create", createHandler)
	http.Handle("/read", readHandler)
	http.Handle("/update/{id}", updateHandler)
	http.Handle("/delete/{id}", deleteHandler)

	log.Fatal(http.ListenAndServe(port, nil))
	log.Println("Listen at port: ", port)

}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request m.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.CreateAt = time.Now()
	return request, nil
}

func decodeReadRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, _ := mux.Vars(r)["id"]
	fmt.Println(id)

	return nil, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
