// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListProjectEKSCapacityTypesParams creates a new ListProjectEKSCapacityTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectEKSCapacityTypesParams() *ListProjectEKSCapacityTypesParams {
	return &ListProjectEKSCapacityTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectEKSCapacityTypesParamsWithTimeout creates a new ListProjectEKSCapacityTypesParams object
// with the ability to set a timeout on a request.
func NewListProjectEKSCapacityTypesParamsWithTimeout(timeout time.Duration) *ListProjectEKSCapacityTypesParams {
	return &ListProjectEKSCapacityTypesParams{
		timeout: timeout,
	}
}

// NewListProjectEKSCapacityTypesParamsWithContext creates a new ListProjectEKSCapacityTypesParams object
// with the ability to set a context for a request.
func NewListProjectEKSCapacityTypesParamsWithContext(ctx context.Context) *ListProjectEKSCapacityTypesParams {
	return &ListProjectEKSCapacityTypesParams{
		Context: ctx,
	}
}

// NewListProjectEKSCapacityTypesParamsWithHTTPClient creates a new ListProjectEKSCapacityTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectEKSCapacityTypesParamsWithHTTPClient(client *http.Client) *ListProjectEKSCapacityTypesParams {
	return &ListProjectEKSCapacityTypesParams{
		HTTPClient: client,
	}
}

/*
ListProjectEKSCapacityTypesParams contains all the parameters to send to the API endpoint

	for the list project e k s capacity types operation.

	Typically these are written to a http.Request.
*/
type ListProjectEKSCapacityTypesParams struct {

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project e k s capacity types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectEKSCapacityTypesParams) WithDefaults() *ListProjectEKSCapacityTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project e k s capacity types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectEKSCapacityTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project e k s capacity types params
func (o *ListProjectEKSCapacityTypesParams) WithTimeout(timeout time.Duration) *ListProjectEKSCapacityTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project e k s capacity types params
func (o *ListProjectEKSCapacityTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project e k s capacity types params
func (o *ListProjectEKSCapacityTypesParams) WithContext(ctx context.Context) *ListProjectEKSCapacityTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project e k s capacity types params
func (o *ListProjectEKSCapacityTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project e k s capacity types params
func (o *ListProjectEKSCapacityTypesParams) WithHTTPClient(client *http.Client) *ListProjectEKSCapacityTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project e k s capacity types params
func (o *ListProjectEKSCapacityTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the list project e k s capacity types params
func (o *ListProjectEKSCapacityTypesParams) WithProjectID(projectID string) *ListProjectEKSCapacityTypesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project e k s capacity types params
func (o *ListProjectEKSCapacityTypesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectEKSCapacityTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}