package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := stringSvc{}

	revHandler := httptransport.NewServer(
		makeReverseEndpoint(svc),
		decodeReverseRequest,
		encodeResponse,
	)

	http.Handle("/reverse", revHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodeReverseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request reverseReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
