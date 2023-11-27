// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-swagger/models"
)

// PostUserCreateOKCode is the HTTP code returned for type PostUserCreateOK
const PostUserCreateOKCode int = 200

/*
PostUserCreateOK create the user operations

swagger:response postUserCreateOK
*/
type PostUserCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.Register `json:"body,omitempty"`
}

// NewPostUserCreateOK creates PostUserCreateOK with default headers values
func NewPostUserCreateOK() *PostUserCreateOK {

	return &PostUserCreateOK{}
}

// WithPayload adds the payload to the post user create o k response
func (o *PostUserCreateOK) WithPayload(payload *models.Register) *PostUserCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user create o k response
func (o *PostUserCreateOK) SetPayload(payload *models.Register) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
PostUserCreateDefault failed to create user

swagger:response postUserCreateDefault
*/
type PostUserCreateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUserCreateDefault creates PostUserCreateDefault with default headers values
func NewPostUserCreateDefault(code int) *PostUserCreateDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUserCreateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post user create default response
func (o *PostUserCreateDefault) WithStatusCode(code int) *PostUserCreateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post user create default response
func (o *PostUserCreateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post user create default response
func (o *PostUserCreateDefault) WithPayload(payload *models.Error) *PostUserCreateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post user create default response
func (o *PostUserCreateDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUserCreateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
