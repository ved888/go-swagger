// Code generated by go-swagger; DO NOT EDIT.

package movie

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetMovieGetParams creates a new GetMovieGetParams object
// no default values defined in spec.
func NewGetMovieGetParams() GetMovieGetParams {

	return GetMovieGetParams{}
}

// GetMovieGetParams contains all the bound params for the get movie get operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetMovieGet
type GetMovieGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*get movie by given id
	  Required: true
	  In: query
	*/
	ID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMovieGetParams() beforehand.
func (o *GetMovieGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qID, qhkID, _ := qs.GetOK("id")
	if err := o.bindID(qID, qhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindID binds and validates parameter ID from query.
func (o *GetMovieGetParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("id", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("id", "query", raw); err != nil {
		return err
	}

	o.ID = raw

	return nil
}
