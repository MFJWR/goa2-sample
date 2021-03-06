// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Admin client HTTP transport
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package client

import (
	"context"
	"net/http"

	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// Client lists the Admin service endpoint HTTP clients.
type Client struct {
	// UserNumber Doer is the HTTP client used to make requests to the user_number
	// endpoint.
	UserNumberDoer goahttp.Doer

	// AdminListUser Doer is the HTTP client used to make requests to the admin
	// list user endpoint.
	AdminListUserDoer goahttp.Doer

	// AdminGetUser Doer is the HTTP client used to make requests to the admin get
	// user endpoint.
	AdminGetUserDoer goahttp.Doer

	// AdminCreateUser Doer is the HTTP client used to make requests to the admin
	// create user endpoint.
	AdminCreateUserDoer goahttp.Doer

	// AdminUpdateUser Doer is the HTTP client used to make requests to the admin
	// update user endpoint.
	AdminUpdateUserDoer goahttp.Doer

	// AdminDeleteUser Doer is the HTTP client used to make requests to the admin
	// delete user endpoint.
	AdminDeleteUserDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the Admin service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		UserNumberDoer:      doer,
		AdminListUserDoer:   doer,
		AdminGetUserDoer:    doer,
		AdminCreateUserDoer: doer,
		AdminUpdateUserDoer: doer,
		AdminDeleteUserDoer: doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// UserNumber returns an endpoint that makes HTTP requests to the Admin service
// user_number server.
func (c *Client) UserNumber() goa.Endpoint {
	var (
		decodeResponse = DecodeUserNumberResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUserNumberRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UserNumberDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("Admin", "user_number", err)
		}
		return decodeResponse(resp)
	}
}

// AdminListUser returns an endpoint that makes HTTP requests to the Admin
// service admin list user server.
func (c *Client) AdminListUser() goa.Endpoint {
	var (
		decodeResponse = DecodeAdminListUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAdminListUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AdminListUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("Admin", "admin list user", err)
		}
		return decodeResponse(resp)
	}
}

// AdminGetUser returns an endpoint that makes HTTP requests to the Admin
// service admin get user server.
func (c *Client) AdminGetUser() goa.Endpoint {
	var (
		decodeResponse = DecodeAdminGetUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAdminGetUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AdminGetUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("Admin", "admin get user", err)
		}
		return decodeResponse(resp)
	}
}

// AdminCreateUser returns an endpoint that makes HTTP requests to the Admin
// service admin create user server.
func (c *Client) AdminCreateUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeAdminCreateUserRequest(c.encoder)
		decodeResponse = DecodeAdminCreateUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAdminCreateUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AdminCreateUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("Admin", "admin create user", err)
		}
		return decodeResponse(resp)
	}
}

// AdminUpdateUser returns an endpoint that makes HTTP requests to the Admin
// service admin update user server.
func (c *Client) AdminUpdateUser() goa.Endpoint {
	var (
		encodeRequest  = EncodeAdminUpdateUserRequest(c.encoder)
		decodeResponse = DecodeAdminUpdateUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAdminUpdateUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AdminUpdateUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("Admin", "admin update user", err)
		}
		return decodeResponse(resp)
	}
}

// AdminDeleteUser returns an endpoint that makes HTTP requests to the Admin
// service admin delete user server.
func (c *Client) AdminDeleteUser() goa.Endpoint {
	var (
		decodeResponse = DecodeAdminDeleteUserResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAdminDeleteUserRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AdminDeleteUserDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("Admin", "admin delete user", err)
		}
		return decodeResponse(resp)
	}
}
