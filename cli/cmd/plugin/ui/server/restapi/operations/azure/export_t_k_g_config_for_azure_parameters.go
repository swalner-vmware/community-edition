// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/vmware-tanzu/community-edition/cli/cmd/plugin/ui/server/models"
)

// NewExportTKGConfigForAzureParams creates a new ExportTKGConfigForAzureParams object
//
// There are no default values defined in the spec.
func NewExportTKGConfigForAzureParams() ExportTKGConfigForAzureParams {

	return ExportTKGConfigForAzureParams{}
}

// ExportTKGConfigForAzureParams contains all the bound params for the export t k g config for azure operation
// typically these are obtained from a http.Request
//
// swagger:parameters exportTKGConfigForAzure
type ExportTKGConfigForAzureParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*parameters to generate TKG configuration file for Azure
	  Required: true
	  In: body
	*/
	Params *models.AzureRegionalClusterParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewExportTKGConfigForAzureParams() beforehand.
func (o *ExportTKGConfigForAzureParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AzureRegionalClusterParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("params", "body", ""))
			} else {
				res = append(res, errors.NewParseError("params", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Params = &body
			}
		}
	} else {
		res = append(res, errors.Required("params", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
