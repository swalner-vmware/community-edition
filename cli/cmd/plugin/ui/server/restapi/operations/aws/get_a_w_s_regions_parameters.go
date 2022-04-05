// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetAWSRegionsParams creates a new GetAWSRegionsParams object
//
// There are no default values defined in the spec.
func NewGetAWSRegionsParams() GetAWSRegionsParams {

	return GetAWSRegionsParams{}
}

// GetAWSRegionsParams contains all the bound params for the get a w s regions operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAWSRegions
type GetAWSRegionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAWSRegionsParams() beforehand.
func (o *GetAWSRegionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
