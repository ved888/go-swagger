// Code generated by go-swagger; DO NOT EDIT.

package movie

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-swagger/models"
)

// GetMovieGetIDOKCode is the HTTP code returned for type GetMovieGetIDOK
const GetMovieGetIDOKCode int = 200

/*
GetMovieGetIDOK get the movie operations

swagger:response getMovieGetIdOK
*/
type GetMovieGetIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Movie `json:"body,omitempty"`
}

// NewGetMovieGetIDOK creates GetMovieGetIDOK with default headers values
func NewGetMovieGetIDOK() *GetMovieGetIDOK {

	return &GetMovieGetIDOK{}
}

// WithPayload adds the payload to the get movie get Id o k response
func (o *GetMovieGetIDOK) WithPayload(payload *models.Movie) *GetMovieGetIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie get Id o k response
func (o *GetMovieGetIDOK) SetPayload(payload *models.Movie) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieGetIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetMovieGetIDDefault failed to get movie

swagger:response getMovieGetIdDefault
*/
type GetMovieGetIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMovieGetIDDefault creates GetMovieGetIDDefault with default headers values
func NewGetMovieGetIDDefault(code int) *GetMovieGetIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMovieGetIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get movie get ID default response
func (o *GetMovieGetIDDefault) WithStatusCode(code int) *GetMovieGetIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get movie get ID default response
func (o *GetMovieGetIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get movie get ID default response
func (o *GetMovieGetIDDefault) WithPayload(payload *models.Error) *GetMovieGetIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie get ID default response
func (o *GetMovieGetIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieGetIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
