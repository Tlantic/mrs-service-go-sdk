package mrs_service_go_sdk

import (
	"net/http"
	"io"
	"fmt"
)

type Client struct {
	client   	*http.Client
	Organization 	string
	Application   	string
	BaseApi		string
	Log      	io.Writer
	AppId		string
}

type TokenResponse struct {
	RefreshToken string `json:"refreshToken"`
	Token        string `json:"accessToken"`
	Type         string `json:"tokenType"`
	ExpiresIn    int64  `json:"expiresIn"`
}

type Setting struct {
	Name string
}

type ErrorResponse struct {
	Response        *http.Response `json:"-"`
	Status          string         `json:"name"`
	Message         interface{}     `json:"message"`
}


func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v", r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}