package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeReverseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(reverseReq)
		v, err := svc.Reverse(req.S)
		if err != nil {
			return reverseRes{v, err.Error()}, err
		}

		return reverseRes{v, ""}, nil
	}
}
