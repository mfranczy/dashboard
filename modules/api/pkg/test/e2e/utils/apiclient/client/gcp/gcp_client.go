// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gcp API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gcp API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListGCPDiskTypes(params *ListGCPDiskTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPDiskTypesOK, error)

	ListGCPDiskTypesNoCredentials(params *ListGCPDiskTypesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPDiskTypesNoCredentialsOK, error)

	ListGCPDiskTypesNoCredentialsV2(params *ListGCPDiskTypesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPDiskTypesNoCredentialsV2OK, error)

	ListGCPNetworks(params *ListGCPNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPNetworksOK, error)

	ListGCPNetworksNoCredentials(params *ListGCPNetworksNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPNetworksNoCredentialsOK, error)

	ListGCPNetworksNoCredentialsV2(params *ListGCPNetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPNetworksNoCredentialsV2OK, error)

	ListGCPSizes(params *ListGCPSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSizesOK, error)

	ListGCPSizesNoCredentials(params *ListGCPSizesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSizesNoCredentialsOK, error)

	ListGCPSizesNoCredentialsV2(params *ListGCPSizesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSizesNoCredentialsV2OK, error)

	ListGCPSubnetworks(params *ListGCPSubnetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSubnetworksOK, error)

	ListGCPSubnetworksNoCredentials(params *ListGCPSubnetworksNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSubnetworksNoCredentialsOK, error)

	ListGCPSubnetworksNoCredentialsV2(params *ListGCPSubnetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSubnetworksNoCredentialsV2OK, error)

	ListGCPZones(params *ListGCPZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPZonesOK, error)

	ListGCPZonesNoCredentials(params *ListGCPZonesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPZonesNoCredentialsOK, error)

	ListGCPZonesNoCredentialsV2(params *ListGCPZonesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPZonesNoCredentialsV2OK, error)

	ListProjectGCPNetworks(params *ListProjectGCPNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGCPNetworksOK, error)

	ListProjectGCPSubnetworks(params *ListProjectGCPSubnetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGCPSubnetworksOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListGCPDiskTypes Lists disk types from GCP
*/
func (a *Client) ListGCPDiskTypes(params *ListGCPDiskTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPDiskTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPDiskTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPDiskTypes",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/gcp/disktypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPDiskTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPDiskTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPDiskTypesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPDiskTypesNoCredentials Lists disk types from GCP
*/
func (a *Client) ListGCPDiskTypesNoCredentials(params *ListGCPDiskTypesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPDiskTypesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPDiskTypesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPDiskTypesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/disktypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPDiskTypesNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPDiskTypesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPDiskTypesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPDiskTypesNoCredentialsV2 Lists disk types from GCP
*/
func (a *Client) ListGCPDiskTypesNoCredentialsV2(params *ListGCPDiskTypesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPDiskTypesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPDiskTypesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPDiskTypesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/disktypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPDiskTypesNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPDiskTypesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPDiskTypesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPNetworks Lists networks from GCP
*/
func (a *Client) ListGCPNetworks(params *ListGCPNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPNetworks",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/gcp/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPNetworksNoCredentials Lists available GCP networks
*/
func (a *Client) ListGCPNetworksNoCredentials(params *ListGCPNetworksNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPNetworksNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPNetworksNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPNetworksNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPNetworksNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPNetworksNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPNetworksNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPNetworksNoCredentialsV2 Lists available GCP networks
*/
func (a *Client) ListGCPNetworksNoCredentialsV2(params *ListGCPNetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPNetworksNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPNetworksNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPNetworksNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPNetworksNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPNetworksNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPNetworksNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPSizes Lists machine sizes from GCP
*/
func (a *Client) ListGCPSizes(params *ListGCPSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPSizes",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/gcp/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPSizesNoCredentials Lists machine sizes from GCP
*/
func (a *Client) ListGCPSizesNoCredentials(params *ListGCPSizesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSizesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPSizesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPSizesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPSizesNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPSizesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPSizesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPSizesNoCredentialsV2 Lists machine sizes from GCP
*/
func (a *Client) ListGCPSizesNoCredentialsV2(params *ListGCPSizesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSizesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPSizesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPSizesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPSizesNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPSizesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPSizesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPSubnetworks Lists subnetworks from GCP
*/
func (a *Client) ListGCPSubnetworks(params *ListGCPSubnetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSubnetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPSubnetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPSubnetworks",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/gcp/{dc}/subnetworks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPSubnetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPSubnetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPSubnetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPSubnetworksNoCredentials Lists available GCP subnetworks
*/
func (a *Client) ListGCPSubnetworksNoCredentials(params *ListGCPSubnetworksNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSubnetworksNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPSubnetworksNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPSubnetworksNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/subnetworks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPSubnetworksNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPSubnetworksNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPSubnetworksNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPSubnetworksNoCredentialsV2 Lists available GCP subnetworks
*/
func (a *Client) ListGCPSubnetworksNoCredentialsV2(params *ListGCPSubnetworksNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPSubnetworksNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPSubnetworksNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPSubnetworksNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/subnetworks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPSubnetworksNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPSubnetworksNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPSubnetworksNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPZones Lists available GCP zones
*/
func (a *Client) ListGCPZones(params *ListGCPZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPZones",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/gcp/{dc}/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPZonesNoCredentials Lists available GCP zones
*/
func (a *Client) ListGCPZonesNoCredentials(params *ListGCPZonesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPZonesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPZonesNoCredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPZonesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPZonesNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPZonesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPZonesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGCPZonesNoCredentialsV2 Lists available GCP zones
*/
func (a *Client) ListGCPZonesNoCredentialsV2(params *ListGCPZonesNoCredentialsV2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGCPZonesNoCredentialsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGCPZonesNoCredentialsV2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGCPZonesNoCredentialsV2",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/clusters/{cluster_id}/providers/gcp/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGCPZonesNoCredentialsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGCPZonesNoCredentialsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGCPZonesNoCredentialsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectGCPNetworks Lists available GCP networks
*/
func (a *Client) ListProjectGCPNetworks(params *ListProjectGCPNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGCPNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGCPNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGCPNetworks",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gcp/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGCPNetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGCPNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGCPNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectGCPSubnetworks Lists available GCP subnetworks
*/
func (a *Client) ListProjectGCPSubnetworks(params *ListProjectGCPSubnetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGCPSubnetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGCPSubnetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGCPSubnetworks",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gcp/subnetworks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGCPSubnetworksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGCPSubnetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGCPSubnetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
