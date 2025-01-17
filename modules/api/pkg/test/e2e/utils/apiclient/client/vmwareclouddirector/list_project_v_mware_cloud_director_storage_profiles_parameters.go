// Code generated by go-swagger; DO NOT EDIT.

package vmwareclouddirector

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

// NewListProjectVMwareCloudDirectorStorageProfilesParams creates a new ListProjectVMwareCloudDirectorStorageProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectVMwareCloudDirectorStorageProfilesParams() *ListProjectVMwareCloudDirectorStorageProfilesParams {
	return &ListProjectVMwareCloudDirectorStorageProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectVMwareCloudDirectorStorageProfilesParamsWithTimeout creates a new ListProjectVMwareCloudDirectorStorageProfilesParams object
// with the ability to set a timeout on a request.
func NewListProjectVMwareCloudDirectorStorageProfilesParamsWithTimeout(timeout time.Duration) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	return &ListProjectVMwareCloudDirectorStorageProfilesParams{
		timeout: timeout,
	}
}

// NewListProjectVMwareCloudDirectorStorageProfilesParamsWithContext creates a new ListProjectVMwareCloudDirectorStorageProfilesParams object
// with the ability to set a context for a request.
func NewListProjectVMwareCloudDirectorStorageProfilesParamsWithContext(ctx context.Context) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	return &ListProjectVMwareCloudDirectorStorageProfilesParams{
		Context: ctx,
	}
}

// NewListProjectVMwareCloudDirectorStorageProfilesParamsWithHTTPClient creates a new ListProjectVMwareCloudDirectorStorageProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectVMwareCloudDirectorStorageProfilesParamsWithHTTPClient(client *http.Client) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	return &ListProjectVMwareCloudDirectorStorageProfilesParams{
		HTTPClient: client,
	}
}

/*
ListProjectVMwareCloudDirectorStorageProfilesParams contains all the parameters to send to the API endpoint

	for the list project v mware cloud director storage profiles operation.

	Typically these are written to a http.Request.
*/
type ListProjectVMwareCloudDirectorStorageProfilesParams struct {

	// Credential.
	Credential *string

	// Organization.
	Organization *string

	// Password.
	Password *string

	// Username.
	Username *string

	// VDC.
	VDC *string

	/* Dc.

	   KKP Datacenter to use for endpoint
	*/
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project v mware cloud director storage profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithDefaults() *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project v mware cloud director storage profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithTimeout(timeout time.Duration) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithContext(ctx context.Context) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithHTTPClient(client *http.Client) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithCredential(credential *string) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithOrganization adds the organization to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithOrganization(organization *string) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetOrganization(organization *string) {
	o.Organization = organization
}

// WithPassword adds the password to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithPassword(password *string) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetPassword(password *string) {
	o.Password = password
}

// WithUsername adds the username to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithUsername(username *string) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetUsername(username *string) {
	o.Username = username
}

// WithVDC adds the vDC to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithVDC(vDC *string) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetVDC(vDC)
	return o
}

// SetVDC adds the vDC to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetVDC(vDC *string) {
	o.VDC = vDC
}

// WithDC adds the dc to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithDC(dc string) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WithProjectID(projectID string) *ListProjectVMwareCloudDirectorStorageProfilesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project v mware cloud director storage profiles params
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectVMwareCloudDirectorStorageProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.Organization != nil {

		// header param Organization
		if err := r.SetHeaderParam("Organization", *o.Organization); err != nil {
			return err
		}
	}

	if o.Password != nil {

		// header param Password
		if err := r.SetHeaderParam("Password", *o.Password); err != nil {
			return err
		}
	}

	if o.Username != nil {

		// header param Username
		if err := r.SetHeaderParam("Username", *o.Username); err != nil {
			return err
		}
	}

	if o.VDC != nil {

		// header param VDC
		if err := r.SetHeaderParam("VDC", *o.VDC); err != nil {
			return err
		}
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
