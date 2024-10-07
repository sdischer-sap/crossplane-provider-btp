/*
SAP XSUAA REST API

Provides access to RoleTemplates, Roles, RoleCollection etc. using the XSUAA REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type CustomidpAPI interface {

	/*
	AreMoreTrustsAllowed Return whether multiple custom platform IdP trusts are allowed

	Return whether multiple custom platform IdP trusts are allowed

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param globalAccountId
	@return CustomidpAPIAreMoreTrustsAllowedRequest
	*/
	AreMoreTrustsAllowed(ctx context.Context, globalAccountId string) CustomidpAPIAreMoreTrustsAllowedRequest

	// AreMoreTrustsAllowedExecute executes the request
	//  @return CustomIdpCockpitResponse
	AreMoreTrustsAllowedExecute(r CustomidpAPIAreMoreTrustsAllowedRequest) (*CustomIdpCockpitResponse, *http.Response, error)

	/*
	GetCustomIdPFeatureFlagForNeoUser Returns the feature flag value for a neo user for a IAS tenant

	Returns the feature flag value for a neo user for a IAS tenant

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest
	*/
	GetCustomIdPFeatureFlagForNeoUser(ctx context.Context) CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest

	// GetCustomIdPFeatureFlagForNeoUserExecute executes the request
	//  @return XSPlatformIdentityProvider
	GetCustomIdPFeatureFlagForNeoUserExecute(r CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest) (*XSPlatformIdentityProvider, *http.Response, error)
}

// CustomidpAPIService CustomidpAPI service
type CustomidpAPIService service

type CustomidpAPIAreMoreTrustsAllowedRequest struct {
	ctx context.Context
	ApiService CustomidpAPI
	globalAccountId string
}

func (r CustomidpAPIAreMoreTrustsAllowedRequest) Execute() (*CustomIdpCockpitResponse, *http.Response, error) {
	return r.ApiService.AreMoreTrustsAllowedExecute(r)
}

/*
AreMoreTrustsAllowed Return whether multiple custom platform IdP trusts are allowed

Return whether multiple custom platform IdP trusts are allowed

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param globalAccountId
 @return CustomidpAPIAreMoreTrustsAllowedRequest
*/
func (a *CustomidpAPIService) AreMoreTrustsAllowed(ctx context.Context, globalAccountId string) CustomidpAPIAreMoreTrustsAllowedRequest {
	return CustomidpAPIAreMoreTrustsAllowedRequest{
		ApiService: a,
		ctx: ctx,
		globalAccountId: globalAccountId,
	}
}

// Execute executes the request
//  @return CustomIdpCockpitResponse
func (a *CustomidpAPIService) AreMoreTrustsAllowedExecute(r CustomidpAPIAreMoreTrustsAllowedRequest) (*CustomIdpCockpitResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomIdpCockpitResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomidpAPIService.AreMoreTrustsAllowed")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sap/rest/globalaccount/{globalAccountId}/platform-identity-providers/status"
	localVarPath = strings.Replace(localVarPath, "{"+"globalAccountId"+"}", url.PathEscape(parameterValueToString(r.globalAccountId, "globalAccountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest struct {
	ctx context.Context
	ApiService CustomidpAPI
}

func (r CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest) Execute() (*XSPlatformIdentityProvider, *http.Response, error) {
	return r.ApiService.GetCustomIdPFeatureFlagForNeoUserExecute(r)
}

/*
GetCustomIdPFeatureFlagForNeoUser Returns the feature flag value for a neo user for a IAS tenant

Returns the feature flag value for a neo user for a IAS tenant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest
*/
func (a *CustomidpAPIService) GetCustomIdPFeatureFlagForNeoUser(ctx context.Context) CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest {
	return CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return XSPlatformIdentityProvider
func (a *CustomidpAPIService) GetCustomIdPFeatureFlagForNeoUserExecute(r CustomidpAPIGetCustomIdPFeatureFlagForNeoUserRequest) (*XSPlatformIdentityProvider, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *XSPlatformIdentityProvider
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomidpAPIService.GetCustomIdPFeatureFlagForNeoUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sap/rest/neo/platformidentityproviders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
