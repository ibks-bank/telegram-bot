package app

import "context"

type SignInRequest struct {
	Token string
}

type SignInResponse struct {
	Token string
}

func (a *app) SignIn(ctx context.Context, request *SignInRequest) (*SignInResponse, error) {
	return &SignInResponse{Token: request.Token}, nil
}
