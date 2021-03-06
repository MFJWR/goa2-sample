// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Users endpoints
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package users

import (
	"context"

	goa "goa.design/goa"
)

// Endpoints wraps the "Users" service endpoints.
type Endpoints struct {
	ListUser   goa.Endpoint
	GetUser    goa.Endpoint
	CreateUser goa.Endpoint
	UpdateUser goa.Endpoint
	DeleteUser goa.Endpoint
}

// NewEndpoints wraps the methods of the "Users" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		ListUser:   NewListUserEndpoint(s),
		GetUser:    NewGetUserEndpoint(s),
		CreateUser: NewCreateUserEndpoint(s),
		UpdateUser: NewUpdateUserEndpoint(s),
		DeleteUser: NewDeleteUserEndpoint(s),
	}
}

// Use applies the given middleware to all the "Users" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ListUser = m(e.ListUser)
	e.GetUser = m(e.GetUser)
	e.CreateUser = m(e.CreateUser)
	e.UpdateUser = m(e.UpdateUser)
	e.DeleteUser = m(e.DeleteUser)
}

// NewListUserEndpoint returns an endpoint function that calls the method "list
// user" of service "Users".
func NewListUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.ListUser(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoa2SampleUserCollection(res, "default")
		return vres, nil
	}
}

// NewGetUserEndpoint returns an endpoint function that calls the method "get
// user" of service "Users".
func NewGetUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetUserPayload)
		res, err := s.GetUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoa2SampleUser(res, "default")
		return vres, nil
	}
}

// NewCreateUserEndpoint returns an endpoint function that calls the method
// "create user" of service "Users".
func NewCreateUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*CreateUserPayload)
		return s.CreateUser(ctx, p)
	}
}

// NewUpdateUserEndpoint returns an endpoint function that calls the method
// "update user" of service "Users".
func NewUpdateUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdateUserPayload)
		res, err := s.UpdateUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGoa2SampleUser(res, "default")
		return vres, nil
	}
}

// NewDeleteUserEndpoint returns an endpoint function that calls the method
// "delete user" of service "Users".
func NewDeleteUserEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*DeleteUserPayload)
		return nil, s.DeleteUser(ctx, p)
	}
}
