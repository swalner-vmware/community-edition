// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// GetAzureEndpointOKCode is the HTTP code returned for type GetAzureEndpointOK
const GetAzureEndpointOKCode int = 200

/*GetAzureEndpointOK Successful retrieval of Azure account parameters

swagger:response getAzureEndpointOK
*/
type GetAzureEndpointOK struct {

	/*
	  In: Body
	*/
	Payload *models.AzureAccountParams `json:"body,omitempty"`
}

// NewGetAzureEndpointOK creates GetAzureEndpointOK with default headers values
func NewGetAzureEndpointOK() *GetAzureEndpointOK {

	return &GetAzureEndpointOK{}
}

// WithPayload adds the payload to the get azure endpoint o k response
func (o *GetAzureEndpointOK) WithPayload(payload *models.AzureAccountParams) *GetAzureEndpointOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure endpoint o k response
func (o *GetAzureEndpointOK) SetPayload(payload *models.AzureAccountParams) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureEndpointOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureEndpointBadRequestCode is the HTTP code returned for type GetAzureEndpointBadRequest
const GetAzureEndpointBadRequestCode int = 400

/*GetAzureEndpointBadRequest Bad Request

swagger:response getAzureEndpointBadRequest
*/
type GetAzureEndpointBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureEndpointBadRequest creates GetAzureEndpointBadRequest with default headers values
func NewGetAzureEndpointBadRequest() *GetAzureEndpointBadRequest {

	return &GetAzureEndpointBadRequest{}
}

// WithPayload adds the payload to the get azure endpoint bad request response
func (o *GetAzureEndpointBadRequest) WithPayload(payload *models.Error) *GetAzureEndpointBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure endpoint bad request response
func (o *GetAzureEndpointBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureEndpointBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureEndpointUnauthorizedCode is the HTTP code returned for type GetAzureEndpointUnauthorized
const GetAzureEndpointUnauthorizedCode int = 401

/*GetAzureEndpointUnauthorized Incorrect credentials

swagger:response getAzureEndpointUnauthorized
*/
type GetAzureEndpointUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureEndpointUnauthorized creates GetAzureEndpointUnauthorized with default headers values
func NewGetAzureEndpointUnauthorized() *GetAzureEndpointUnauthorized {

	return &GetAzureEndpointUnauthorized{}
}

// WithPayload adds the payload to the get azure endpoint unauthorized response
func (o *GetAzureEndpointUnauthorized) WithPayload(payload *models.Error) *GetAzureEndpointUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure endpoint unauthorized response
func (o *GetAzureEndpointUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureEndpointUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureEndpointInternalServerErrorCode is the HTTP code returned for type GetAzureEndpointInternalServerError
const GetAzureEndpointInternalServerErrorCode int = 500

/*GetAzureEndpointInternalServerError Internal server error

swagger:response getAzureEndpointInternalServerError
*/
type GetAzureEndpointInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureEndpointInternalServerError creates GetAzureEndpointInternalServerError with default headers values
func NewGetAzureEndpointInternalServerError() *GetAzureEndpointInternalServerError {

	return &GetAzureEndpointInternalServerError{}
}

// WithPayload adds the payload to the get azure endpoint internal server error response
func (o *GetAzureEndpointInternalServerError) WithPayload(payload *models.Error) *GetAzureEndpointInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure endpoint internal server error response
func (o *GetAzureEndpointInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureEndpointInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
